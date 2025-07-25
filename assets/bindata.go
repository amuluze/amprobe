// Code generated for package assets by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/resources/configs/config.toml
// assets/resources/configs/model.conf
package assets

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

// Mode return file modify time
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

var _assetsResourcesConfigsConfigToml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x94\x5f\x53\x1a\x57\x18\xc6\xef\xcf\xa7\x38\xb3\xde\xc4\x99\x4e\x84\x24\x26\xde\x78\xa1\xc1\xd6\x4c\xa1\x71\x82\xbd\x62\x98\xcc\x22\x47\xd8\x11\x77\xe9\xee\x81\xe8\x74\x32\x43\xa6\x1a\x20\x2d\x85\x56\x10\x8b\xc6\x98\xa9\x56\x4c\x0b\x4b\xda\x8e\x5a\x5d\xe1\xcb\x70\xf6\xc0\x15\x5f\xa1\x73\xce\x2e\x7f\x96\xd4\x2e\x5c\x30\xec\xfb\xfc\xf6\xec\xf3\x3e\xef\x1b\xf8\x5c\x0a\x21\x35\x08\x16\x15\x0d\xc3\x59\x28\xb8\xee\xf2\x8f\x00\xf9\x35\x01\xa3\x18\xc7\xe9\xfe\x4f\xa4\xf0\x07\x39\x68\x90\xb7\x29\xb0\xa4\xa8\xac\x70\xc6\xe5\x72\xc1\xe1\x35\x5a\x48\x7f\xd7\x49\xfe\x57\xe0\x8f\x26\x70\x58\x79\x21\x2f\x4b\xeb\x48\x49\x30\xcd\x7d\x57\xbf\xb0\x6d\xec\x75\xf7\xb7\xc9\xf6\x5f\xdd\x72\x8d\xd6\xb2\xa4\xb9\xdd\x39\xdf\x36\xcb\xe7\x66\xf9\xbc\x5b\x6a\xdd\x21\xb9\x52\xfb\x26\x47\x4f\x7f\x9e\x04\x7e\x94\x44\xea\x22\x12\xc3\x48\x65\xc7\x8b\xab\x4a\x08\x09\x60\x2e\x1e\xff\x4a\x5c\x47\x23\xff\x2c\xa9\x68\x55\x51\xd7\xe0\x2c\x5c\x15\x63\x1a\x02\x20\xf0\x85\xa2\xae\x07\x81\x07\x85\x12\x11\x38\x0b\xb1\x9a\x40\x10\x3a\x4e\x6c\xee\xe9\xa4\xf0\x1b\x31\x52\xa4\xa0\x77\x1a\xdf\x75\xf4\x92\x59\x7d\x4f\x8c\x3c\xf0\xcc\x2f\x6f\xc6\x39\x5c\xfb\x26\x26\x61\x24\x8c\x68\x4a\x0d\x33\x57\x27\x57\x3b\xf4\xe3\x35\x39\xfc\xfe\x0e\xdd\xaf\x93\x6c\xce\x2c\xea\xe6\x0f\xaf\x68\x65\x6b\xec\x76\xcf\xa8\xc4\x15\x0d\x47\x54\xa4\x7d\x06\x2d\xd6\x24\xf0\x89\x1b\x5e\x69\x15\x61\x89\x9f\xff\xd1\xbd\x81\x8f\x13\xb0\x53\x6f\xd2\x9b\x7a\xa7\x75\x68\xfe\x78\x42\xf2\x7a\xfb\xfa\xa4\x9b\xce\xd1\x62\x95\x91\x0f\x52\xdd\x52\x8b\xf9\x53\xfe\xdb\xf6\xa7\x67\x54\xb8\x45\x3e\x71\xe3\x69\x1c\xc9\x8f\x15\x59\xd6\xe0\x2c\x74\x4f\x8f\x03\x87\xa7\xe2\x1c\x72\x7c\x6a\x66\x77\x88\x91\xb2\x1e\x64\x96\x1a\x0c\xf1\x24\x1c\x43\x7d\xc4\x80\x30\x40\xd0\xb3\xab\x6e\xf9\x4f\x5b\xf0\xf1\xa8\x7d\x59\x1b\xb0\x86\x94\x65\x31\x14\x43\xac\x11\xd2\x06\x37\xef\xb9\xf0\xa9\x6f\x9d\xf7\x55\x52\xc8\x91\x6c\x8e\x1a\x29\xb0\x20\x33\xc5\x5c\x02\x2b\x3e\x29\xa2\x8a\x18\xf5\x1b\x35\x68\x4e\x41\xa7\xc5\x6a\x27\xfd\x81\xbc\xa9\x9a\x7b\x47\xa4\xb1\x35\x0a\xa2\xd7\x3b\xe6\xe1\x16\x00\x01\xcf\xfc\x30\xc0\xc2\x58\x97\x6d\x37\x1d\xe1\xbd\xa5\xc8\x0e\xee\xd7\x9a\x95\xb5\xf1\x22\x5a\xac\x9a\x99\x0b\x52\xc8\x81\x25\x51\xd3\x5e\x28\x6a\xd8\x59\x35\x01\x89\xfe\x9a\x1e\xbd\x02\x9e\xf9\x7e\x38\xa7\x10\x5e\x99\x12\xd7\x79\x46\xa7\xac\xa4\x8e\x9a\x01\xfc\x7e\xaf\x4f\x09\xa3\xb1\xa7\x4d\x40\xbf\xdf\x6b\x67\x11\x04\xbc\x4a\x24\x08\x9e\x26\x70\x9c\xcf\x90\x13\x19\x53\x22\x9a\xc5\xbd\x1b\x53\x22\x9c\x5d\x3e\x21\xad\x32\xa9\xed\x99\xc5\x66\xe7\x42\x27\xcd\x2d\xe0\x45\x49\x14\x63\xd2\x30\x9b\x05\xe7\x5b\x59\xcd\xe1\x1a\x5a\xcb\xd2\xab\x53\xf0\x4c\xc1\x22\x96\x14\x99\x25\x69\xbc\xd4\xaa\xd6\xf3\xe4\xf8\x8c\x64\xd2\x24\xf3\xba\x7d\x99\x6a\x5f\x7e\xb0\xf4\xe6\x6e\xba\x7d\x7d\xce\xa2\x34\x17\xe1\xc9\xfe\x0f\x35\xd7\xb3\xd4\x54\xda\xad\xb7\xb4\xf4\x0b\x7c\x04\xc9\xf1\x19\x8b\x12\x47\x00\x10\x98\x4b\xe0\x68\xd0\x0e\xc6\xa7\x63\xeb\xcc\x05\xf0\x4b\x11\x59\x92\x23\x3e\x84\xa3\x0a\x6f\xc6\xa2\x7f\xda\x7d\x8f\xd9\x40\x6b\x4d\x52\xc8\x99\xbb\xff\x10\x23\x7f\xc7\x1a\xcf\x9e\x51\xe1\xb7\xa7\x16\xfd\xf7\x67\x1e\x4c\xf1\xdf\x93\x7d\xc4\x97\x68\x93\xe9\x6d\x5f\x05\x38\x40\xac\xa1\x4d\xb0\xb0\x11\x97\x54\x14\x76\x4e\xab\x9d\x9b\xb4\x79\xf0\xce\x9a\xca\x9e\x91\x19\xec\xad\x9e\x91\x05\xcf\xd0\xaa\x8a\xb4\xe8\x50\x3c\xf3\xf0\x81\x8b\x6f\x40\x92\xb9\x30\x77\x1b\xff\xa7\x1d\x0e\x91\x98\xc0\xd1\xe7\x23\x0b\x68\x0d\x6d\x42\x7b\x7a\x40\xe0\xb1\xa8\x85\x24\xf9\x36\xbb\x9c\x66\xad\xf0\xda\xc1\x3e\xe4\x5b\x72\xa4\x94\x67\xc3\xce\x1c\x1b\x48\xaf\x22\x86\x9d\x40\x27\x8e\xd4\x2b\xe6\xc1\x3b\x6b\x32\xc9\x9b\xa3\xce\xcd\x0d\xad\xed\xd2\xd2\xc9\x40\xfc\x44\xc6\x48\x95\x45\x16\xbc\x87\x2e\xf6\xd2\xb7\x08\x2c\x03\xd8\xb7\x52\x1c\xb7\x01\x04\x96\x45\x6d\x2d\x08\x38\x2b\xc9\x59\xee\xe9\x61\xc2\xdc\xc0\x23\x69\x6c\xe5\x7f\x0b\x3d\x28\x29\xad\x20\xb6\xbb\x02\x42\x32\x2c\xba\x85\xe0\x4b\xb0\x80\xa3\x48\x95\x11\xe6\x05\x6c\x22\xad\xdb\x08\x47\x5d\x42\xf0\xe5\xbf\x01\x00\x00\xff\xff\x89\xe7\xd9\x0c\xf9\x06\x00\x00"

func assetsResourcesConfigsConfigTomlBytes() ([]byte, error) {
	return bindataRead(
		_assetsResourcesConfigsConfigToml,
		"assets/resources/configs/config.toml",
	)
}

func assetsResourcesConfigsConfigToml() (*asset, error) {
	bytes, err := assetsResourcesConfigsConfigTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/resources/configs/config.toml", size: 1785, mode: os.FileMode(420), modTime: time.Unix(1746693478, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsResourcesConfigsModelConf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x41\xca\x84\x30\x0c\x85\xf7\x3d\xc5\x5b\x89\x82\x78\x83\x9e\x44\x44\x6a\xff\x54\x2b\xd5\xf4\x8f\x11\x99\xdb\x0f\x75\x16\x83\x30\xab\x84\xc7\xf7\x25\xaf\x17\xfa\x3f\xe9\xd0\xf1\x8f\x42\xdc\xa3\x46\xde\x07\x23\xb0\x38\xce\xa9\x05\x4f\x6b\x0b\xe7\xd5\x98\x3e\x73\x8a\xfe\xf5\xc0\xf2\x0f\x4c\x38\xd1\x03\x9a\x61\x31\xb6\x18\xbf\x27\x28\x04\xf2\x3a\x18\x2a\x3a\x6f\x54\x5f\x0b\x09\xa1\xce\x1d\x05\x85\xb5\x70\x29\xf1\xd5\x34\xc6\xf4\x9b\x53\xbf\x90\x1c\x83\xd9\x60\x31\xd7\xd2\xdd\xff\x72\x19\x0d\xaa\x0a\xd2\xf1\xb4\x16\x27\xdf\xcb\x9d\x38\xaf\x9f\xa4\x34\x7a\x07\x00\x00\xff\xff\x69\x9d\x0d\x22\xe0\x00\x00\x00"

func assetsResourcesConfigsModelConfBytes() ([]byte, error) {
	return bindataRead(
		_assetsResourcesConfigsModelConf,
		"assets/resources/configs/model.conf",
	)
}

func assetsResourcesConfigsModelConf() (*asset, error) {
	bytes, err := assetsResourcesConfigsModelConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/resources/configs/model.conf", size: 224, mode: os.FileMode(420), modTime: time.Unix(1746673019, 0)}
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
	"assets/resources/configs/config.toml": assetsResourcesConfigsConfigToml,
	"assets/resources/configs/model.conf":  assetsResourcesConfigsModelConf,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
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
	"assets": &bintree{nil, map[string]*bintree{
		"resources": &bintree{nil, map[string]*bintree{
			"configs": &bintree{nil, map[string]*bintree{
				"config.toml": &bintree{assetsResourcesConfigsConfigToml, map[string]*bintree{}},
				"model.conf":  &bintree{assetsResourcesConfigsModelConf, map[string]*bintree{}},
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
