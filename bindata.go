// Code generated by go-bindata.
// sources:
// db/migrations/1_initial_schema.sql
// DO NOT EDIT!

package conveyor

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

var _dbMigrations1_initial_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x53\x4d\x8f\xd3\x30\x10\xbd\xe7\x57\x3c\xed\x25\xa9\xa0\x20\x24\x6e\x15\x87\x42\xbd\x10\xa9\x4a\xa1\x4d\xb4\x7b\x8b\xdc\x64\x9a\x58\x24\x76\x70\xec\x96\xf0\xeb\x71\xdc\xb4\xac\xca\xa1\x07\xc4\x6d\xe6\x79\x3e\xde\xbc\x19\xcf\xe7\x78\xd5\x8a\x4a\x73\x43\xc8\xba\xe0\xd3\x96\x2d\x53\x06\xf6\x9c\xb2\x64\x17\x6f\x12\xc4\x8f\x48\x36\xa9\x03\xe2\x5d\xba\x43\xdd\x1b\xa5\x69\x71\x2f\xec\xc1\x5a\x51\xce\x55\xdf\x77\x0f\x8b\xe0\x12\x9c\x2e\x3f\xae\x19\xf6\x56\x34\x65\x8f\x28\x00\x44\x89\x31\xce\x27\x26\xd9\x7a\x8d\x15\x7b\x5c\x66\xeb\xd4\xa3\x79\x45\x92\x46\x56\xf9\xf1\x7d\x34\x43\xa7\x45\xcb\xf5\x80\xef\x34\xbc\x76\xa9\x9a\x3a\xd5\x0b\xc7\x65\x80\xa1\x9f\x66\x84\xf6\x9a\xcb\xa2\xbe\xba\x7d\xcd\xff\xd8\x66\x9c\xee\xe2\x15\x9a\x9c\x5b\xe6\xdc\xc0\x88\x96\xdc\x63\xdb\xe1\x24\x4c\xad\xec\x19\xc1\x2f\x25\x09\x25\x1d\xb8\x6d\x0c\x22\xa9\x4e\x8e\x00\x7f\xf9\x16\x5a\x53\x84\xb3\xa9\xb4\xbe\x5f\xcc\xb7\x55\x6d\xd7\xd0\xfd\xd8\x60\x76\x2b\x99\xeb\x20\x0e\xbc\x30\xff\xac\x9a\xd7\x3e\xff\xab\x80\xa6\x03\x69\x92\x05\xf5\xd3\x76\x22\x51\xfa\xe1\x5c\x72\xf5\x1f\x74\xf3\x13\xba\xbb\x7b\x22\xb7\x25\x65\x9b\x12\x24\x7b\xab\x5d\xa7\xda\x45\x9f\x08\x4a\x36\x03\x6a\x7e\x24\xbc\x43\x47\xb2\x14\xb2\x7a\xeb\x99\x39\xe3\x4c\x11\x07\xa5\xc1\xe5\x80\x4a\x1c\x49\x8e\xcb\x7e\x73\xd1\x2c\x4b\xe2\x6f\x19\x43\x9c\xac\xd8\x33\xac\x14\x3f\x2c\xe5\xe7\x1c\x77\xa6\xd3\xf5\x65\xbb\x38\xf9\x8c\xbd\xd1\x44\x88\x5c\xf2\x0c\x4f\x5f\xd8\x96\x39\xdb\x5f\xca\x07\x84\x97\x76\x21\x36\x5b\x5c\xd1\x89\x4c\x38\x0d\x70\xfd\x38\x2b\x75\x92\xc1\x6a\xbb\xf9\x7a\xbb\xb1\xc5\x4b\xf4\xdc\x7c\x11\xfc\x0e\x00\x00\xff\xff\xdd\x3b\x35\x29\x73\x03\x00\x00")

func dbMigrations1_initial_schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations1_initial_schemaSql,
		"db/migrations/1_initial_schema.sql",
	)
}

func dbMigrations1_initial_schemaSql() (*asset, error) {
	bytes, err := dbMigrations1_initial_schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/1_initial_schema.sql", size: 883, mode: os.FileMode(420), modTime: time.Unix(1454021875, 0)}
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
	"db/migrations/1_initial_schema.sql": dbMigrations1_initial_schemaSql,
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
	"db": &bintree{nil, map[string]*bintree{
		"migrations": &bintree{nil, map[string]*bintree{
			"1_initial_schema.sql": &bintree{dbMigrations1_initial_schemaSql, map[string]*bintree{}},
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

