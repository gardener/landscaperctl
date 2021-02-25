package installations

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gardener/component-cli/ociclient"
	ociopts "github.com/gardener/component-cli/ociclient/options"
	"github.com/gardener/component-cli/pkg/commands/constants"
	cdv2 "github.com/gardener/component-spec/bindings-go/apis/v2"
	"github.com/gardener/component-spec/bindings-go/ctf"
	cdoci "github.com/gardener/component-spec/bindings-go/oci"
	lsv1alpha1 "github.com/gardener/landscaper/apis/core/v1alpha1"
	"github.com/gardener/landscaper/pkg/kubernetes"
	lsjsonschema "github.com/gardener/landscaper/pkg/landscaper/jsonschema"
	registrycomponents "github.com/gardener/landscaper/pkg/landscaper/registry/components"
	"github.com/gardener/landscaper/pkg/landscaper/registry/components/cdutils"
	"github.com/gardener/landscaper/pkg/utils"
	"github.com/go-logr/logr"
	"github.com/mandelsoft/vfs/pkg/memoryfs"
	"github.com/mandelsoft/vfs/pkg/osfs"
	"github.com/mandelsoft/vfs/pkg/vfs"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/xeipuuv/gojsonschema"
	yamlv3 "gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"sigs.k8s.io/yaml"

	"github.com/gardener/landscapercli/pkg/logger"
	"github.com/gardener/landscapercli/pkg/util"
)

type createOpts struct {
	// baseURL is the oci registry where the component is stored.
	baseURL string
	// componentName is the unique name of the component in the registry.
	componentName string
	// version is the component version in the oci registry.
	version string
	// OciOptions contains all exposed options to configure the oci client.
	OciOptions ociopts.Options

	//outputPath is the path to write the installation.yaml to
	outputPath string

	// name of the blueprint resource in the component descriptor (optional if only one blueprint resource is specified in the component descriptor)
	blueprintResourceName string
	name                  string
	renderSchemaInfo      bool
}

func NewCreateCommand(ctx context.Context) *cobra.Command {
	opts := &createOpts{}
	cmd := &cobra.Command{
		Use:     "create [baseURL] [componentName] [componentVersion]",
		Args:    cobra.ExactArgs(3),
		Aliases: []string{"c"},
		Example: "landscaper-cli installations create my-registry:5000 github.com/my-component v0.1.0",
		Short:   "create an installation template for a component which is stored in an OCI registry",
		Run: func(cmd *cobra.Command, args []string) {
			if err := opts.Complete(args); err != nil {
				cmd.PrintErr(err.Error())
				os.Exit(1)
			}

			if err := opts.run(ctx, cmd, logger.Log, osfs.New()); err != nil {
				cmd.PrintErr(err.Error())
				os.Exit(1)
			}
		},
	}

	cmd.SetOut(os.Stdout)

	opts.AddFlags(cmd.Flags())

	return cmd
}

func (o *createOpts) run(ctx context.Context, cmd *cobra.Command, log logr.Logger, fs vfs.FileSystem) error {
	repoCtx := cdv2.RepositoryContext{
		Type:    cdv2.OCIRegistryType,
		BaseURL: o.baseURL,
	}
	ociRef, err := cdoci.OCIRef(repoCtx, o.componentName, o.version)
	if err != nil {
		return fmt.Errorf("invalid component reference: %w", err)
	}

	ociClient, _, err := o.OciOptions.Build(log, fs)
	if err != nil {
		return fmt.Errorf("unable to build oci client: %s", err.Error())
	}

	cdresolver := cdoci.NewResolver().WithOCIClient(ociClient).WithRepositoryContext(repoCtx)
	cd, blobResolver, err := cdresolver.Resolve(ctx, o.componentName, o.version)
	if err != nil {
		return fmt.Errorf("unable to to fetch component descriptor %s: %w", ociRef, err)
	}

	blueprintRes, err := o.getBlueprintResource(cd)
	if err != nil {
		return err
	}

	data, err := resolveBlueprint(ctx, *blueprintRes, ociClient, blobResolver)
	if err != nil {
		return fmt.Errorf("cannot resolve blueprint: %w", err)
	}

	memFS := memoryfs.New()
	if err := utils.ExtractTarGzip(data, memFS, "/"); err != nil {
		return fmt.Errorf("cannot extract blueprint blob: %w", err)
	}

	blueprintData, err := vfs.ReadFile(memFS, lsv1alpha1.BlueprintFileName)
	if err != nil {
		return fmt.Errorf("cannot read %s: %w", lsv1alpha1.BlueprintFileName, err)
	}

	blueprint := &lsv1alpha1.Blueprint{}
	if _, _, err := serializer.NewCodecFactory(kubernetes.LandscaperScheme).UniversalDecoder().Decode(blueprintData, nil, blueprint); err != nil {
		return fmt.Errorf("cannot decode blueprint: %w", err)
	}

	installation := buildInstallation(o.name, cd, *blueprintRes, blueprint)

	var marshaledYaml []byte
	if o.renderSchemaInfo {
		ociRegistry, err := registrycomponents.NewOCIRegistryWithOCIClient(ociClient)
		if err != nil {
			return fmt.Errorf("cannot build oci registry: %w", err)
		}

		schemaRefResolver := &jsonschemaRefResolver{
			loaderConfig: &lsjsonschema.LoaderConfig{
				LocalTypes:                 blueprint.LocalTypes,
				BlueprintFs:                memFS,
				ComponentDescriptor:        cd,
				ComponentResolver:          ociRegistry,
				ComponentReferenceResolver: cdutils.ComponentReferenceResolverFromResolver(ociRegistry, repoCtx),
			},
		}

		commentedYaml, err := annotateInstallationWithSchemaComments(installation, blueprint, schemaRefResolver)
		if err != nil {
			return fmt.Errorf("cannot add JSON schema comment: %w", err)
		}

		marshaledYaml, err = util.MarshalYaml(commentedYaml)
		if err != nil {
			return fmt.Errorf("cannot marshal installation yaml: %w", err)
		}
	} else {
		marshaledYaml, err = yaml.Marshal(installation)
		if err != nil {
			return fmt.Errorf("cannot marshal installation yaml: %w", err)
		}
	}

	if o.outputPath == "" {
		cmd.Println(string(marshaledYaml))
	} else {
		f, err := os.Create(o.outputPath)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", o.outputPath, err)
		}
		_, err = f.Write(marshaledYaml)
		if err != nil {
			return fmt.Errorf("error writing file %s: %w", o.outputPath, err)
		}
		cmd.Printf("Wrote installation to %s", o.outputPath)
	}

	return nil
}

func (o *createOpts) Complete(args []string) error {
	o.baseURL = args[0]
	o.componentName = args[1]
	o.version = args[2]

	cliHomeDir, err := constants.CliHomeDir()
	if err != nil {
		return err
	}
	o.OciOptions.CacheDir = filepath.Join(cliHomeDir, "components")
	if err := os.MkdirAll(o.OciOptions.CacheDir, os.ModePerm); err != nil {
		return fmt.Errorf("unable to create cache directory %s: %w", o.OciOptions.CacheDir, err)
	}

	if len(o.baseURL) == 0 {
		return errors.New("the base url must be defined")
	}
	if len(o.componentName) == 0 {
		return errors.New("a component name must be defined")
	}
	if len(o.version) == 0 {
		return errors.New("a component's Version must be defined")
	}
	return nil
}

func annotateInstallationWithSchemaComments(installation *lsv1alpha1.Installation, blueprint *lsv1alpha1.Blueprint, schemaRefResolver *jsonschemaRefResolver) (*yamlv3.Node, error) {
	out, err := yaml.Marshal(installation)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal installation yaml: %w", err)
	}

	commentedInstallationYaml := &yamlv3.Node{}
	err = yamlv3.Unmarshal(out, commentedInstallationYaml)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal installation yaml: %w", err)
	}

	err = addImportSchemaComments(commentedInstallationYaml, blueprint, schemaRefResolver)
	if err != nil {
		return nil, fmt.Errorf("cannot add schema comments for imports: %w", err)
	}

	err = addExportSchemaComments(commentedInstallationYaml, blueprint, schemaRefResolver)
	if err != nil {
		return nil, fmt.Errorf("cannot add schema comments for exports: %w", err)
	}

	return commentedInstallationYaml, nil
}

func resolveBlueprint(ctx context.Context, blueprintRes cdv2.Resource, ociClient ociclient.Client, blobResolver ctf.BlobResolver) (*bytes.Buffer, error) {
	var data bytes.Buffer
	if blueprintRes.Access.GetType() == cdv2.OCIRegistryType {
		ref, ok := blueprintRes.Access.Object["imageReference"].(string)
		if !ok {
			return nil, fmt.Errorf("cannot parse imageReference to string")
		}

		manifest, err := ociClient.GetManifest(ctx, ref)
		if err != nil {
			return nil, fmt.Errorf("cannot get manifest: %w", err)
		}

		err = ociClient.Fetch(ctx, ref, manifest.Layers[0], &data)
		if err != nil {
			return nil, fmt.Errorf("cannot get manifest layer: %w", err)
		}
	} else {
		_, err := blobResolver.Resolve(ctx, blueprintRes, &data)
		if err != nil {
			return nil, fmt.Errorf("unable to to resolve blob of blueprint resource: %w", err)
		}
	}

	return &data, nil
}

func (o *createOpts) getBlueprintResource(cd *cdv2.ComponentDescriptor) (*cdv2.Resource, error) {
	blueprintResources := map[string]cdv2.Resource{}
	for _, resource := range cd.ComponentSpec.Resources {
		if resource.IdentityObjectMeta.Type == lsv1alpha1.BlueprintResourceType || resource.IdentityObjectMeta.Type == lsv1alpha1.OldBlueprintType {
			blueprintResources[resource.Name] = resource
		}
	}

	var blueprintRes cdv2.Resource
	numberOfBlueprints := len(blueprintResources)
	if numberOfBlueprints == 0 {
		return nil, fmt.Errorf("no blueprint resources defined in the component descriptor")
	} else if numberOfBlueprints == 1 && o.blueprintResourceName == "" {
		// access the only blueprint in the map. the flag blueprint-resource-name is ignored in this case.
		for _, entry := range blueprintResources {
			blueprintRes = entry
		}
	} else {
		if o.blueprintResourceName == "" {
			return nil, fmt.Errorf("the blueprint resource name must be defined since multiple blueprint resources exist in the component descriptor")
		}
		ok := false
		blueprintRes, ok = blueprintResources[o.blueprintResourceName]
		if !ok {
			return nil, fmt.Errorf("blueprint %s is not defined as a resource in the component descriptor", o.blueprintResourceName)
		}
	}

	return &blueprintRes, nil
}

func addExportSchemaComments(commentedInstallationYaml *yamlv3.Node, blueprint *lsv1alpha1.Blueprint, schemaRefResolver *jsonschemaRefResolver) error {
	_, exportsDataValueNode := util.FindNodeByPath(commentedInstallationYaml, "spec.exports.data")
	if exportsDataValueNode != nil {

		for _, dataImportNode := range exportsDataValueNode.Content {
			n1, n2 := util.FindNodeByPath(dataImportNode, "name")
			exportName := n2.Value

			var expdef lsv1alpha1.ExportDefinition
			for _, bpexp := range blueprint.Exports {
				if bpexp.Name == exportName {
					expdef = bpexp
					break
				}
			}

			schemaLoader := gojsonschema.NewBytesLoader(expdef.Schema.RawMessage)
			schemas, err := schemaRefResolver.resolveRefs("", schemaLoader)
			if err != nil {
				return fmt.Errorf("cannot load jsonschema for export definition %s: %w", expdef.Name, err)
			}

			schemasStr, err := schemas.toString()
			if err != nil {
				return fmt.Errorf("cannot convert jsonschema to string: %w", err)
			}
			n1.HeadComment = schemasStr
		}
	}

	_, exportTargetsValueNode := util.FindNodeByPath(commentedInstallationYaml, "spec.exports.targets")
	if exportTargetsValueNode != nil {
		for _, targetExportNode := range exportTargetsValueNode.Content {
			n1, n2 := util.FindNodeByPath(targetExportNode, "name")
			targetName := n2.Value

			var expdef lsv1alpha1.ExportDefinition
			for _, bpexp := range blueprint.Exports {
				if bpexp.Name == targetName {
					expdef = bpexp
					break
				}
			}
			n1.HeadComment = "Target type: " + expdef.TargetType
		}
	}

	return nil
}

func addImportSchemaComments(commentedInstallationYaml *yamlv3.Node, blueprint *lsv1alpha1.Blueprint, schemaRefResolver *jsonschemaRefResolver) error {
	_, importDataValueNode := util.FindNodeByPath(commentedInstallationYaml, "spec.imports.data")
	if importDataValueNode != nil {
		for _, dataImportNode := range importDataValueNode.Content {
			n1, n2 := util.FindNodeByPath(dataImportNode, "name")
			importName := n2.Value

			var impdef lsv1alpha1.ImportDefinition
			for _, bpimp := range blueprint.Imports {
				if bpimp.Name == importName {
					impdef = bpimp
					break
				}
			}

			schemaLoader := gojsonschema.NewBytesLoader(impdef.Schema.RawMessage)
			schemas, err := schemaRefResolver.resolveRefs("", schemaLoader)
			if err != nil {
				return fmt.Errorf("cannot load jsonschema for import definition %s: %w", impdef.Name, err)
			}

			schemasStr, err := schemas.toString()
			if err != nil {
				return fmt.Errorf("cannot convert jsonschema to string: %w", err)
			}
			n1.HeadComment = schemasStr
		}
	}

	_, targetsValueNode := util.FindNodeByPath(commentedInstallationYaml, "spec.imports.targets")
	if targetsValueNode != nil {
		for _, targetImportNode := range targetsValueNode.Content {
			n1, n2 := util.FindNodeByPath(targetImportNode, "name")
			targetName := n2.Value

			var impdef lsv1alpha1.ImportDefinition
			for _, bpimp := range blueprint.Imports {
				if bpimp.Name == targetName {
					impdef = bpimp
					break
				}
			}
			n1.HeadComment = "Target type: " + impdef.TargetType
		}
	}

	return nil
}

func (o *createOpts) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.name, "name", "my-installation", "name of the installation")
	fs.BoolVar(&o.renderSchemaInfo, "render-schema-info", true, "render schema information of the component's imports and exports as comments into the installation")
	fs.StringVar(&o.blueprintResourceName, "blueprint-resource-name", "", "name of the blueprint resource in the component descriptor (optional if only one blueprint resource is specified in the component descriptor)")
	fs.StringVarP(&o.outputPath, "output-file", "o", "", "file path for the resulting installation yaml")
	o.OciOptions.AddFlags(fs)
}

func buildInstallation(name string, cd *cdv2.ComponentDescriptor, blueprintRes cdv2.Resource, blueprint *lsv1alpha1.Blueprint) *lsv1alpha1.Installation {
	dataImports := []lsv1alpha1.DataImport{}
	targetImports := []lsv1alpha1.TargetImportExport{}
	for _, imp := range blueprint.Imports {
		if imp.TargetType != "" {
			targetImport := lsv1alpha1.TargetImportExport{
				Name: imp.Name,
			}
			targetImports = append(targetImports, targetImport)
		} else {
			dataImport := lsv1alpha1.DataImport{
				Name: imp.Name,
			}
			dataImports = append(dataImports, dataImport)
		}
	}

	dataExports := []lsv1alpha1.DataExport{}
	targetExports := []lsv1alpha1.TargetImportExport{}
	for _, exp := range blueprint.Exports {
		if exp.TargetType != "" {
			targetExport := lsv1alpha1.TargetImportExport{
				Name: exp.Name,
			}
			targetExports = append(targetExports, targetExport)
		} else {
			dataExport := lsv1alpha1.DataExport{
				Name: exp.Name,
			}
			dataExports = append(dataExports, dataExport)
		}
	}

	obj := &lsv1alpha1.Installation{
		TypeMeta: metav1.TypeMeta{
			APIVersion: lsv1alpha1.SchemeGroupVersion.String(),
			Kind:       "Installation",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: lsv1alpha1.InstallationSpec{
			ComponentDescriptor: &lsv1alpha1.ComponentDescriptorDefinition{
				Reference: &lsv1alpha1.ComponentDescriptorReference{
					RepositoryContext: &cd.RepositoryContexts[0],
					ComponentName:     cd.ObjectMeta.Name,
					Version:           cd.ObjectMeta.Version,
				},
			},
			Blueprint: lsv1alpha1.BlueprintDefinition{
				Reference: &lsv1alpha1.RemoteBlueprintReference{
					ResourceName: blueprintRes.Name,
				},
			},
			Imports: lsv1alpha1.InstallationImports{
				Data:    dataImports,
				Targets: targetImports,
			},
			Exports: lsv1alpha1.InstallationExports{
				Data:    dataExports,
				Targets: targetExports,
			},
		},
	}

	return obj
}

type jsonSchema struct {
	Ref    string                 `json:"ref"`
	Schema map[string]interface{} `json:"schema"`
}
type jsonSchemaList []jsonSchema

func (l jsonSchemaList) toString() (string, error) {
	if len(l) == 0 {
		return "", nil
	}

	buf := bytes.Buffer{}

	_, err := buf.WriteString("JSON Schema\n")
	if err != nil {
		return "", fmt.Errorf("cannot write to buffer: %w", err)
	}

	blueprintSchema, err := json.MarshalIndent(l[0].Schema, "", "  ")
	if err != nil {
		return "", fmt.Errorf(`cannot marshal blueprint jsonschema: %w`, err)
	}

	_, err = buf.Write(blueprintSchema)
	if err != nil {
		return "", fmt.Errorf("cannot write to buffer: %w", err)
	}

	if len(l) > 1 {
		_, err = buf.WriteString("\n \nReferenced JSON Schemas\n")
		if err != nil {
			return "", fmt.Errorf("cannot write to buffer: %w", err)
		}

		marshaledSchemaRefs, err := json.MarshalIndent(l[1:], "", "  ")
		if err != nil {
			return "", fmt.Errorf(`cannot marshal jsonschema refs: %w`, err)
		}

		_, err = buf.Write(marshaledSchemaRefs)
		if err != nil {
			return "", fmt.Errorf("cannot write to buffer: %w", err)
		}
	}

	return buf.String(), nil
}

type jsonschemaRefResolver struct {
	loaderConfig *lsjsonschema.LoaderConfig
}

func (l *jsonschemaRefResolver) resolveRefs(ref string, schemaLoader gojsonschema.JSONLoader) (jsonSchemaList, error) {
	//Wrap default loader if config is defined
	if l.loaderConfig != nil {
		schemaLoader = lsjsonschema.NewWrappedLoader(*l.loaderConfig, schemaLoader)
	}

	schema, err := schemaLoader.LoadJSON()
	if err != nil {
		return nil, fmt.Errorf("cannot load jsonschema: %w", err)
	}

	schemamap, ok := schema.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("cannot convert ")
	}

	allSchemas := jsonSchemaList{}
	allSchemas = append(allSchemas, jsonSchema{Ref: ref, Schema: schemamap})

	for key, value := range schemamap {
		if key == "$ref" {
			refStr, ok := value.(string)
			if !ok {
				return nil, fmt.Errorf("cannot parse value of $ref to string")
			}

			newLoader := schemaLoader.LoaderFactory().New(refStr)
			ref, err := newLoader.JsonReference()
			if err != nil {
				return nil, err
			}

			if !ref.IsCanonical() {
				return nil, fmt.Errorf("ref %s must be canonical", ref.String())
			}

			subSchemas, err := l.resolveRefs(ref.String(), newLoader)
			if err != nil {
				return nil, fmt.Errorf("cannot resolve ref %s: %w", ref.String(), err)
			}

			allSchemas = append(allSchemas, subSchemas...)
		}
	}

	return allSchemas, nil
}
