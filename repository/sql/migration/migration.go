// Code generated by go-bindata.
// sources:
// repository/sql/migration/resources/001_initial_db.sql
// DO NOT EDIT!

package migration

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

var _resources001_initial_dbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x5f\x6f\xdb\x36\x10\x7f\xd7\xa7\xb8\xb7\xd8\x58\x3a\xd4\x41\x03\x6c\x28\xfa\xa0\x24\x5c\x2b\xd4\x51\x3a\x45\x1e\xd0\x27\x96\x92\x68\x87\xb0\x44\x79\x14\x65\xcc\xdf\x7e\xa0\xfe\x92\x94\x94\xc8\x01\x1a\x54\x4f\xf2\xe9\xfe\xdf\xef\x8e\x47\xbf\x7b\x07\xbf\x65\x6c\x27\x88\xa4\xb0\x39\x38\xb7\x01\x72\x43\x04\xa1\x7b\xb3\x46\x50\x16\x54\x14\xce\xc2\x81\xea\x0d\xb3\x04\x00\xc0\xf3\x43\x70\x37\xe1\x03\xf6\xfc\xdb\x00\xdd\x23\x3f\x74\x14\xf9\x5b\xe0\xdd\xbb\xc1\x77\xf8\x8a\xbe\x5f\x3a\x00\x9c\x64\x14\xea\xe7\x1f\x37\xb8\xfd\xe2\x06\x8b\xd5\x9f\xab\x25\x8c\x3f\xfe\x43\x08\xfe\x66\xbd\x56\x92\x34\x23\x2c\x9d\x2f\xd9\x48\xc5\x82\x12\x49\x13\x4c\x24\xdc\xb9\x21\x0a\xbd\x7b\x04\x77\xe8\x2f\x77\xb3\x0e\xe1\x76\x13\x04\xc8\x0f\xb1\x22\x3e\x86\xee\xfd\x37\xc3\x5e\x79\x48\x06\x92\xd3\x4f\x2b\xc5\x0a\x1c\xa5\x79\xbc\xa7\x09\xdc\x78\x61\x67\x2a\xba\x78\x7f\xf1\x42\x7c\xac\xc0\x24\xc9\x18\x07\x38\x57\xf2\xf6\xc1\x7f\x0c\x03\x57\x15\xa0\xaa\x0c\x56\x49\xc6\x25\xe3\x09\xfd\xcf\x01\xd8\xf8\xde\xdf\x1b\x04\x0b\x45\x5d\x3a\x4b\x07\x00\xf9\x9f\x3d\x1f\xc1\x27\xf0\x38\xcf\xef\x6e\x94\x8a\x2f\x6e\xf0\x88\x42\xf8\x04\xa5\xdc\xfe\x91\x45\x1f\x3e\x3a\x4e\x9d\x3a\x90\x24\x4a\x69\x5d\x68\x52\xca\x27\xad\xee\xea\xa7\x2a\x3e\xe3\x12\x48\x29\x73\xcc\x78\x2c\x68\x46\xb9\xac\x2a\x7f\x10\x2c\x23\xe2\x04\x7b\x7a\xba\x34\xa1\x02\x95\x48\xf7\xf0\x5c\x02\x2f\xd3\x54\x71\x65\x54\x3e\xe5\x0d\x13\x1c\x89\x88\x9f\x88\x58\x5c\x5d\x5f\x2f\x0d\xae\x23\x49\xcb\x16\x46\x20\x19\x3f\x45\x69\x1e\xd5\xba\x1a\x8e\x38\xe7\x85\x14\x44\xd9\xd9\xee\x71\xef\x7d\xf5\xea\x00\x6c\x73\x41\xd9\x8e\x2b\xe7\x60\xd1\xb8\xb6\x04\x41\xb7\x54\x50\x1e\xd3\xa2\xce\x64\xff\xa9\x8a\x28\xe7\x0d\x2a\x20\x26\x45\x4c\x12\xda\x52\x13\x9a\x52\x8d\xaa\xb8\x29\xdf\x31\x4e\xf5\x1c\xab\x50\x0a\x2a\xc7\x72\x5c\x95\x6a\xdc\x51\x65\xb3\x23\xf6\xfe\xd8\xf5\x89\x89\xa4\xbb\x5c\x30\x5a\xd7\xa7\xf9\x79\x9a\x5f\x1e\xad\x31\xcd\xbc\xdb\x15\x62\x05\x3e\xb2\x82\x29\xa3\x10\x31\x09\x09\xdd\x92\x32\x95\x10\x5d\xac\x2e\x6c\x46\x12\x4b\x76\x54\x6a\x4d\xc6\xf7\x26\xe3\x81\x08\xca\x65\x8d\x0d\x03\x18\x13\x05\xed\x63\xc5\xc3\x57\x15\x32\xde\xee\xed\x12\x77\x36\x8c\x22\xf7\xe2\xb0\xd0\xe4\xdf\xa4\xda\x33\xa3\xc8\xb9\xe1\x65\x1f\x87\x8d\x80\x83\xc8\x93\x32\x96\x75\xfd\x9b\x1f\xf3\xca\x3f\x16\xc1\x47\x67\x5c\x3d\x3e\x52\x51\xb0\x9c\x9b\x66\x1a\xe2\x7c\xb4\x69\xfe\xb5\xcf\xa0\xf2\xc3\x47\x07\x8d\x8e\xd7\x29\xd8\x3e\xaf\xe1\x73\xe8\xf9\xd6\xe7\x4a\x7c\xf5\x61\x4a\x5a\xc7\xe3\x41\xb0\xd8\x72\x20\xa1\x31\xcb\x48\xba\x58\x5d\x5f\xc2\xd5\xb8\x0e\xdd\x3e\x49\x92\xfa\x78\xd1\x34\x10\x49\x25\xcb\x68\xd7\x2b\xc3\x23\x4a\xd7\x50\x03\xd1\xd0\xd1\x69\x78\x39\x02\xbd\x91\x9b\x67\xd0\xcf\xcf\x47\xf0\x6f\x49\xb8\x64\xf2\x34\x91\x83\x89\x3c\xda\xd2\xb8\xe4\xac\x09\x60\x4e\x05\x87\xf3\xc0\x86\x26\x6e\x5b\x01\xf7\x30\x1b\x1b\x09\xdd\x47\x63\x26\xb4\xc2\x06\xc3\xc4\x40\x78\x65\xeb\x77\xfe\x3d\x7e\xdd\xe0\xf6\x98\xce\xf9\x20\x10\xc3\x85\x89\x7e\xef\x27\x46\x46\x0e\x76\xef\x57\x4d\xa5\x57\xcc\x3e\x17\x8c\x31\xdc\xb7\xa8\x6e\xf8\x52\x17\x5a\x4e\xa4\x5e\x77\x02\xb3\x68\xbb\xc7\x57\x3f\x21\xdd\x30\x36\x7f\x67\x7b\xb4\xb2\x3d\xd2\xe3\xfa\x55\x4e\x85\xe7\xdd\xd7\x30\xa2\x33\x98\x6e\xda\x40\x29\x64\x1e\xef\x2b\x64\x54\x6f\xfd\x39\x7b\xd6\x98\x9e\x31\x9e\x0d\x34\x69\xeb\xde\x2c\xd1\xe1\x48\x39\xd7\x22\x49\xba\x81\xaa\x26\x60\x21\x49\x76\x98\x39\x48\x35\xfc\xd4\x39\x7a\xab\x09\x32\x6a\x7c\x1c\xad\xe7\x2f\xa9\x0a\x68\x2a\x42\x0b\xa1\x1d\xf9\x95\x10\xb5\x9c\xcc\x79\x4d\x19\xdb\x4d\x75\x01\x63\xe8\xb5\xab\x4d\x23\xf9\xcc\x98\x93\x82\xf0\x42\xed\x91\xed\xce\xa1\x11\x5e\x79\xf9\x38\x0f\x91\xd6\xaa\x72\x8e\xa8\x71\x4d\x99\xb5\x1f\x18\xa8\x94\x64\xa7\x7f\x9a\x71\x38\x4e\xed\x16\xe7\xf6\x83\x76\xf7\x35\xc5\x5f\x8e\xd9\xbc\x78\xe9\xd5\xc3\xf5\xdf\x06\x3f\xe9\xee\xd5\xa2\xda\xf2\xc1\x70\xe0\xcd\xf6\x02\x45\x7d\xf0\xe1\x0e\xad\x51\x88\x40\x5d\xac\xd5\x5d\xfd\xb5\x0d\xf7\xc3\x08\xa2\x75\xfe\xf7\x6a\x79\x50\x01\xfc\xa8\x7b\x49\xe7\x9a\x68\xa9\x5a\x5f\x93\xd7\x31\xa9\xbe\x85\xff\x0f\x00\x00\xff\xff\x21\x61\x61\xa5\x00\x12\x00\x00")

func resources001_initial_dbSqlBytes() ([]byte, error) {
	return bindataRead(
		_resources001_initial_dbSql,
		"resources/001_initial_db.sql",
	)
}

func resources001_initial_dbSql() (*asset, error) {
	bytes, err := resources001_initial_dbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/001_initial_db.sql", size: 4608, mode: os.FileMode(420), modTime: time.Unix(1519914826, 0)}
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
	"resources/001_initial_db.sql": resources001_initial_dbSql,
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
	"resources": &bintree{nil, map[string]*bintree{
		"001_initial_db.sql": &bintree{resources001_initial_dbSql, map[string]*bintree{}},
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

