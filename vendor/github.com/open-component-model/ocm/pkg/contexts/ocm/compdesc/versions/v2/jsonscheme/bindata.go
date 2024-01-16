// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package jsonscheme generated by go-bindata.// sources:
// ../../../../../../../resources/component-descriptor-v2-schema.yaml
package jsonscheme

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ResourcesComponentDescriptorV2SchemaYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdd\x6f\x1b\xb9\x11\x7f\xd7\x5f\x31\x38\x07\xa0\x1c\x7b\x2d\x47\x45\x0a\x44\x2f\x86\x9b\x43\x81\xa0\xbd\xf3\x21\x49\xfb\x50\xc7\x0d\xa8\xdd\x91\xc4\x94\x4b\xaa\x24\xa5\x58\x97\xf3\xff\x7e\x20\xb9\xdc\xef\x5d\x7d\xd9\xb9\x0b\x10\x3f\x24\x22\x77\x38\x33\x1c\xfe\xe6\x83\xb3\xfb\x8c\x25\x13\x20\x0b\x63\x96\x7a\x32\x1a\xcd\xa9\x4a\x50\xa0\xba\x88\xb9\x5c\x25\x23\x1d\x2f\x30\xa5\x7a\x14\xcb\x74\x29\x05\x0a\x13\x25\xa8\x63\xc5\x96\x46\xaa\x68\x3d\x26\x83\x67\x9e\xa2\xc4\xe1\x93\x96\x22\xf2\xb3\x17\x52\xcd\x47\x89\xa2\x33\x33\x1a\x5f\x8e\x2f\xa3\x17\xe3\x8c\x21\x19\x04\x36\x4c\x8a\x09\x90\x9b\x25\x0a\x78\x1d\x64\xc0\x4f\x32\x41\x0e\xeb\x31\x14\xd4\x33\x26\x98\x25\xd6\x93\x01\x40\x8a\x86\xda\xff\x01\xcc\x66\x89\x13\x20\x72\xfa\x09\x63\x43\xdc\x54\x95\x73\xae\x38\x14\x8a\xbb\xf5\x09\x35\xd4\x2f\x50\xf8\xff\x15\x53\x98\x78\x8e\x00\x11\x10\x2f\xf7\xdf\xa8\x34\x93\xc2\x53\x2d\x95\x5c\xa2\x32\x0c\x75\xa0\xab\x10\x85\xc9\x5c\x25\x6d\x14\x13\x73\x32\x70\xea\xaa\x39\x76\xea\xdb\x64\x4c\xf9\x5c\x2a\x66\x16\x69\xc1\x74\x49\x8d\x41\x65\x37\xf4\xdf\x5b\x1a\xfd\x7a\x67\xff\xb9\x8c\x5e\x8d\x3e\x46\x77\x67\xcf\x48\x46\x16\x4b\x31\x63\xf3\x09\x7c\x81\x07\x37\x43\x93\xc4\xd9\x8c\xf2\x5f\x0a\x19\x30\xa3\x5c\xe3\x00\x80\xd3\x29\xf2\x4e\xad\x5a\x8c\x22\x68\x8a\xa4\x18\xae\x29\x5f\x61\xd7\x16\x2c\x6d\xa7\x49\xfc\xa4\x5b\x3f\x81\x2f\x0f\x61\x5c\x37\x64\x69\xcf\xeb\xdb\xcb\xe8\x55\x69\xa7\x9a\xcd\x05\x13\xf3\x86\x84\xa9\x94\x1c\xa9\x08\x64\x25\xc3\xdb\xbf\x67\x0a\x67\x13\x20\x27\xa3\x12\x9c\x46\x8e\xc6\x1d\x53\x0e\x95\x9f\x73\xe5\x5b\x14\x4f\xe9\xfd\x3f\x51\xcc\xcd\x62\x02\xe3\x97\x2f\x07\xad\x87\x13\xf9\xd3\xb9\x7b\x3e\xbc\xbd\xb8\xab\x4d\x9d\x3e\x0f\x73\x5f\xc6\xe7\x0f\xc3\x51\xe5\xf1\xc7\x96\x25\x1f\xed\x9a\x53\xbb\xf7\x01\x00\x4b\x50\x18\x66\x36\xd7\xc6\x28\x36\x5d\x19\xfc\x07\x6e\xbc\xaa\x29\x13\xb9\x5e\x6d\x5a\x59\xe1\xc3\xdb\xe8\xe3\x59\x50\x24\x4c\x9e\x5e\x79\xd6\x0a\x39\xbd\xc7\xe4\x1d\xa6\x6b\x54\x9e\xe7\x09\x18\xfa\x3f\x14\x30\x53\x32\x05\xed\x1e\x58\x97\x06\x2a\x12\xa0\xc9\xa7\x95\x36\x98\x80\x91\x40\x39\x97\x9f\x81\x0a\x90\x4b\x8f\x37\xe0\x48\x13\x26\xe6\x40\xd6\xe4\x1c\x52\xfa\x49\xaa\x48\x0a\xbe\x39\x77\x4b\xdd\xf8\x22\x65\x22\x9b\x0d\xb2\x16\x4c\x43\x8a\x54\x68\x30\x0b\x84\x99\xb4\x5c\x2d\x13\x6f\x7e\x0d\x54\xa1\x15\x65\x91\xc3\x92\xaa\xbe\x3a\x28\xfc\xe2\x62\x7c\xf1\x97\xf2\xef\x68\x26\xe5\xd9\x94\xaa\x6c\x6e\x5d\x26\x58\xb7\x51\xbc\xb8\x18\x87\x5f\x39\x59\x89\x3e\xff\x59\x59\x56\x36\xf6\xfa\xee\x6a\x78\xf9\xdb\xed\x8b\xe8\xd5\xdd\x87\xe4\xf9\xe9\xf0\x6a\xf2\xe1\xa2\x3c\x71\x7a\xd5\x3e\x15\x0d\x87\x57\x93\x62\xf2\xb7\x0f\x89\x3b\xa3\xeb\xe8\x3f\xd1\x9d\xc5\x7f\xf8\x1d\x58\xee\x48\x7c\x1a\x24\x9e\x0d\xcb\x0f\xce\x1c\x93\xca\x8c\xa3\xcc\x7c\xac\x19\xc5\x1a\xd0\xdb\x16\xd1\x36\xd6\x8f\xb4\x0d\x47\xad\x8e\xd7\x06\x65\x02\x0f\x1e\x8a\x4b\xa9\x99\x91\x6a\xf3\x5a\x0a\x83\xf7\x66\x9f\x30\x65\xa9\xba\xc2\x92\xe3\xd0\x13\xa9\x65\xcc\xde\xb6\xcb\xa6\x9c\xdf\xcc\x0a\x29\xad\x3b\x6a\xa8\x5d\x44\xcb\xba\x9e\x99\xae\x53\xaa\xf1\x5f\x8a\x93\x22\xe6\x35\x54\xb6\x7f\x19\x59\x79\xaa\x23\xa8\xfa\x34\x50\x8a\x63\x3f\xd1\xe5\xb2\x12\x29\x7b\x97\x02\xa0\x58\xa5\x13\xb8\x25\x2b\xc5\x7f\xa1\x66\x41\xce\x81\xe8\x05\x1d\xbf\xfc\x6b\x94\xb0\x39\x6a\x43\xee\x06\x35\x3e\xfb\x72\x76\x36\x9e\x33\x6d\xd4\xc6\x72\xbf\x79\xfd\x26\x1f\xde\xd9\x33\xa0\x71\x8c\x5a\xef\x98\xde\xad\x65\x1c\x15\xcc\xa4\xca\x96\xa2\x86\xa1\x1d\xe1\xbd\x41\x61\x53\x8a\x3e\xdd\x02\x96\x01\xc0\x9c\x99\xc5\x6a\x7a\xdd\x2f\xbb\x17\x6d\x6e\x68\x21\x50\x3a\x50\x37\x33\x3b\x08\x8d\x75\xb3\x79\x05\x73\xf3\x67\x82\xb6\x2c\xb7\x28\xed\xa7\x88\x65\x9a\x32\xd3\xe7\x13\x42\x0a\x3c\xc6\x2e\x47\xee\xfb\x67\x29\xd0\x03\x43\xcb\x95\x8a\xf1\xc7\xdc\xe1\xf6\x50\xc7\x96\x23\xf9\x20\x2b\x34\xf2\xb1\xe5\x90\x0f\x3c\x84\x0e\xaf\x6a\x3a\xaa\x8c\xd6\x60\x97\x2d\xc1\x7b\xa3\xe8\x9b\x8c\x60\x4b\xb5\xd2\xe0\x43\xba\xaa\xa7\x8e\x08\x55\xca\x99\x64\xf7\xe3\x70\xb5\xa2\x6e\x10\x51\xa5\xe8\xa6\xd8\x39\x33\x98\x56\xe2\x56\xab\x0e\x8e\x57\x58\x54\x76\x76\x37\x16\x9b\x9b\x59\x35\x48\xb6\x32\xf1\xeb\xc8\x76\xc2\xb2\x5f\xef\x40\x6e\x2f\x31\x81\x78\x00\xe0\x63\xde\xbb\x25\xc6\x7b\x80\x6d\x41\xf5\xe2\x3a\x94\xf0\x05\x04\xa5\x4a\x29\x67\x9a\x5a\x41\xcd\xc7\xae\x1a\xee\x80\x5d\x85\x61\xfd\x10\xfc\x41\x05\x80\xb6\x0a\xe9\x5d\xe2\xcb\xf0\x76\x8a\x81\xaf\xb4\xa9\x59\x29\xdc\xd3\x08\xb4\x67\x87\x76\x94\x62\xc2\xe8\xfb\xe0\x79\x3b\xdd\x81\xf6\x54\xde\x4f\xe5\x72\x0a\xaa\x6a\x06\x79\xbf\x40\x4f\xe4\xd3\x88\x9c\xb9\xe2\x33\xdf\x36\x94\xae\x39\xbd\xf6\x39\x34\x1a\xe5\x0c\xf6\x88\x39\x95\x1d\x7a\x8c\x6e\x71\xfc\x02\xc8\xe5\x2b\x54\x49\xf1\xce\x95\x15\x00\xf8\x84\x80\xb6\xee\xff\xf1\x10\xd7\xc8\x77\x7e\xc0\x3e\x1b\x11\xae\x85\xe6\x51\x42\xe9\xde\x06\xcd\x6d\x92\x77\x30\xbc\x71\xf6\x49\x98\x9d\x19\x6a\x9b\xa5\x5a\xb5\xab\x94\x80\x8f\x90\x22\x0e\xc4\x98\xc2\x2c\x67\x97\xcd\x01\x47\xe6\x8f\x3a\xfc\x9c\xfd\xb5\x8a\xdf\x86\x5a\x67\x6b\xd1\x48\x6d\x5d\x84\x0a\x45\x8c\xee\xf6\x0a\xc3\xa2\xbd\xc5\x65\x4c\xf9\x69\x56\x6b\x74\x15\x30\x01\x3a\xef\x90\x63\x6c\xa4\x3a\x14\x69\x4f\x90\x56\xcb\x7d\x8c\xb7\x61\x97\x87\xda\x25\xe7\xb4\x6b\x4f\xa8\x15\x77\x11\x90\x75\x7f\x27\xad\xa5\xf3\xb2\x1f\xb4\xfb\x0a\x33\x38\x01\x1a\x9b\x15\xe5\x7c\x33\x29\x24\x45\x2e\xda\x7f\x1e\x81\x5e\x62\xcc\x28\xb7\x58\x35\x8a\xc5\x4e\xc8\xb7\x5b\xcb\x3d\x59\xa1\x56\x8f\x00\x52\x60\xbd\x50\xcb\x64\x89\x15\xe7\x3b\x54\x5a\xb5\x00\x1a\x42\x45\x91\xaa\xf7\xbc\xfb\x05\x06\x7a\xe7\xfe\x65\x86\x49\x38\x71\xeb\x9d\xe3\x17\x5c\xce\xb3\x76\xd4\x4a\x1b\x48\xa9\x89\x17\x25\x67\xd0\x8d\x2b\x44\xf3\x1a\xc8\x5d\x09\x56\x9a\x2a\x57\xac\xdf\x6f\x16\xf9\xae\x7c\xe0\x7e\x24\xc4\x7a\x66\x45\xf6\xf1\x87\xb0\xf3\x55\xd3\x41\x80\x9c\x03\xc1\x7b\x83\x4a\x50\x9e\xdf\xb6\xbf\xdd\xfb\x8f\x8c\xd9\xdf\xb8\xdc\xfd\x02\xe4\x6c\xf0\x77\xc6\x51\x6f\xb4\xc1\x74\xff\xb5\x37\x6d\x02\x9f\x3a\x7a\xc8\x98\xbd\x49\xe9\xfc\xa8\x3e\x85\x1b\x32\xcb\x25\xcf\x9b\x8f\xd2\xc0\x28\xf7\xbb\x02\x9e\xaa\x62\xb6\x74\x24\x0b\x73\x1e\xb1\x31\x4e\x37\xc1\x2f\x8f\xdb\x0f\x90\x4c\x25\x02\x45\x2f\x6a\xd6\x75\xbb\xba\xb6\x1b\xa8\x96\x15\xf6\x7a\x95\x52\xc1\x66\xa8\x4d\xfd\x5e\x55\x13\x7a\xe0\xe5\xcd\x5b\xc6\x07\x70\xef\x28\x5e\x03\x0d\x46\x6e\x91\x58\x07\x6a\x53\x9c\xa7\x08\xa2\x0c\x55\x73\x34\x98\x40\x2c\x85\xc9\x0b\xa5\x4e\xf6\x9a\xfd\xda\xbb\x17\xfb\x1c\x98\x80\xe9\xc6\xa0\x0e\x32\xa6\xd6\xd8\x75\xbe\x62\x95\x4e\xed\x81\x0e\x00\x3a\x5d\xf6\x08\xb8\xcc\x18\xc7\x22\x5f\x1e\x8b\x98\x16\x0d\x0b\xf4\x04\x51\x5d\x76\x09\xcf\xcb\xe6\x00\xb3\xa0\x06\x98\x76\x7b\xb7\xe6\x67\xc2\x3d\xfb\xc1\x3e\xd4\x3f\x40\xc2\x94\x2b\xcc\x37\x9d\xe7\x11\xec\x76\xf3\x48\xfe\xf5\x04\x06\xbb\xa9\xfb\x59\x3f\x38\xab\xc0\x74\xfe\x0e\x9f\x99\x59\x64\xa6\x89\x57\x4a\xa1\x30\xd0\xf6\x42\xbc\xcf\x4a\x21\xb4\xbe\xcd\x2a\xa3\x7d\x6c\xd4\x51\x71\x75\x1a\xf1\x7b\x8d\xb4\x3d\x97\xb8\xc3\xf8\xfa\x85\x49\x57\x71\x51\x4a\xbb\x5f\x27\xd9\x0f\x00\x8a\x26\xed\x11\x0e\xbb\x0a\x6f\x69\x8e\x4c\xef\x56\x99\xfc\x38\x56\x3d\x6f\x64\x06\x00\x73\x14\xa8\x58\xfc\x07\xbe\x4d\xc9\x34\xf0\x2f\x54\xb2\xc1\x77\xcf\xfe\x13\x78\x76\x71\x30\x7e\xfe\x8f\x75\xec\x0a\x50\xbf\x56\x11\x9f\x67\xa6\x9d\xdb\x55\x7b\xf7\xa7\x9a\x38\x6d\xbc\xb3\xd7\xa5\x87\x4b\x25\xd7\x2c\x29\x4e\x34\x02\x52\x69\x32\x54\x7b\x5e\x79\x3d\xaf\x2b\xfc\x2b\x2b\xfe\x14\xdd\xdc\x58\xa1\xbb\x18\xbf\x67\x4d\xb7\xbb\x0d\x08\x3d\xcf\x8e\xb1\x78\xdf\x3f\x93\x2a\xa5\x66\x02\x09\x35\x18\x19\x96\xf7\xab\x9b\x26\xdc\x1b\xb6\x8d\x6b\x6f\xdf\x7d\xb6\xf1\x85\x06\x81\x93\x50\xde\xf0\xcd\x39\x7c\x46\x90\x82\x6f\xb2\xaf\x92\xdc\x2d\x40\x8a\xa0\x6c\x38\xd2\x2d\x8e\xf9\x64\xee\x97\xa1\xe1\x91\xfa\x1d\xb5\x37\xe2\xf9\x01\x37\x21\xf9\x38\x02\x9b\x8c\xeb\xad\xfe\xa7\x3c\xfb\x72\x8f\x90\xec\x08\x96\x4a\xed\xba\xd3\xa2\x5a\x56\x74\xa1\xa9\xdd\xa4\xf0\xe5\x61\x30\x18\xd4\xe2\x54\x39\x08\x45\x40\x52\xf4\x9f\x99\x96\x03\x05\x19\x54\xc3\x40\xf1\x39\x6b\xc7\x17\x8a\x9e\x45\x2d\x3e\xf6\x1f\x10\x29\xbf\x9b\xac\xd6\x1a\xa5\x03\xa9\x1c\x46\xff\xeb\x3f\x52\x7b\xf3\x77\x04\xcf\xf6\x97\x65\xe4\xf7\x00\x00\x00\xff\xff\xa3\xb7\xcb\x40\x89\x2c\x00\x00")

func ResourcesComponentDescriptorV2SchemaYamlBytes() ([]byte, error) {
	return bindataRead(
		_ResourcesComponentDescriptorV2SchemaYaml,
		"../../../../../../../resources/component-descriptor-v2-schema.yaml",
	)
}

func ResourcesComponentDescriptorV2SchemaYaml() (*asset, error) {
	bytes, err := ResourcesComponentDescriptorV2SchemaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../../../../../resources/component-descriptor-v2-schema.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"../../../../../../../resources/component-descriptor-v2-schema.yaml": ResourcesComponentDescriptorV2SchemaYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"..": {nil, map[string]*bintree{
		"..": {nil, map[string]*bintree{
			"..": {nil, map[string]*bintree{
				"..": {nil, map[string]*bintree{
					"..": {nil, map[string]*bintree{
						"..": {nil, map[string]*bintree{
							"..": {nil, map[string]*bintree{
								"resources": {nil, map[string]*bintree{
									"component-descriptor-v2-schema.yaml": {ResourcesComponentDescriptorV2SchemaYaml, map[string]*bintree{}},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
