// Code generated by go-bindata.
// sources:
// config.default.toml
// DO NOT EDIT!

package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func bindataRead(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len
	return b, nil
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configDefaultToml = "\x70\x6f\x72\x74\x20\x3d\x20\x33\x30\x30\x30\x0a\x0a\x6c\x6f\x67\x5f\x66\x6f\x72\x6d\x61\x74\x20\x3d\x20\x22\x24\x7b\x74\x69\x6d\x65\x5f\x72\x66\x63\x33\x33\x33\x39\x7d\x20\x24\x7b\x72\x65\x6d\x6f\x74\x65\x5f\x69\x70\x7d\x20\x24\x7b\x73\x74\x61\x74\x75\x73\x7d\x20\x24\x7b\x6d\x65\x74\x68\x6f\x64\x7d\x5c\x74\x24\x7b\x75\x72\x69\x7d\x5c\x74\x5c\x74\x24\x7b\x6c\x61\x74\x65\x6e\x63\x79\x5f\x68\x75\x6d\x61\x6e\x7d\x5c\x6e\x22\x0a\x0a\x23\x20\x4e\x6f\x74\x20\x69\x6d\x70\x6c\x65\x6d\x65\x6e\x74\x65\x64\x0a\x5b\x68\x74\x74\x70\x73\x5d\x0a\x70\x6f\x72\x74\x20\x3d\x20\x34\x34\x33\x0a\x23\x20\x52\x65\x64\x69\x72\x65\x63\x74\x20\x68\x74\x74\x70\x20\x74\x6f\x20\x68\x74\x74\x70\x73\x0a\x72\x65\x64\x69\x72\x65\x63\x74\x20\x3d\x20\x74\x72\x75\x65\x0a\x0a\x5b\x63\x72\x61\x77\x6c\x5d\x0a\x23\x20\x48\x6f\x77\x20\x6f\x66\x74\x65\x6e\x20\x66\x65\x65\x64\x73\x20\x73\x68\x6f\x75\x6c\x64\x20\x62\x65\x20\x66\x65\x74\x63\x68\x65\x64\x0a\x69\x6e\x74\x65\x72\x76\x61\x6c\x20\x3d\x20\x22\x31\x35\x6d\x22\x0a\x23\x20\x4d\x61\x78\x69\x6d\x75\x6d\x20\x6e\x75\x6d\x62\x65\x72\x20\x6f\x66\x20\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x73\x20\x74\x6f\x20\x75\x73\x65\x20\x77\x68\x65\x6e\x20\x66\x65\x74\x63\x68\x69\x6e\x67\x20\x66\x65\x65\x64\x73\x0a\x6d\x61\x78\x5f\x63\x6f\x6e\x6e\x20\x3d\x20\x31\x32\x38\x0a"

func configDefaultTomlBytes() ([]byte, error) {
	return bindataRead(
		_configDefaultToml,
		"config.default.toml",
	)
}

func configDefaultToml() (*asset, error) {
	bytes, err := configDefaultTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.default.toml", size: 322, mode: os.FileMode(436), modTime: time.Unix(1493349997, 0)}
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
	"config.default.toml": configDefaultToml,
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
	"config.default.toml": &bintree{configDefaultToml, map[string]*bintree{}},
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

