// Code generated by go-bindata.
// sources:
// sql/insert_account.sql
// sql/select_account_credentials.sql
// DO NOT EDIT!

package store

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

var _sqlInsert_accountSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xcb\xc1\x0a\xc3\x30\x08\x80\xe1\x7b\x9e\xc2\x63\x0b\x7d\x83\x1d\xf6\x28\x45\x12\x0b\x82\xd5\xa2\xa6\x7b\xfd\x8d\x25\x97\xc1\xc4\xc3\xff\x21\xb2\x06\x79\x02\x6b\x1a\x60\xad\xd6\x35\x61\x29\xf0\x19\x37\xa1\xed\x5b\x37\x39\x1f\x4c\x6d\x88\x4e\x64\x19\x79\x61\xc4\xcb\x7c\x1e\x0e\xf6\xc8\x5d\xf1\x9c\x6f\x82\x3f\xac\x4e\x98\xd4\x76\xd3\xe1\x7e\xb5\xe9\xb2\xc2\x8d\xd2\x29\x60\x79\x6e\xf0\x67\xd7\x47\x79\x07\x00\x00\xff\xff\xe6\x08\x46\x65\xa8\x00\x00\x00")

func sqlInsert_accountSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlInsert_accountSql,
		"sql/insert_account.sql",
	)
}

func sqlInsert_accountSql() (*asset, error) {
	bytes, err := sqlInsert_accountSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/insert_account.sql", size: 168, mode: os.FileMode(420), modTime: time.Unix(1465501808, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlSelect_account_credentialsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x04\xc0\xc1\x0d\x85\x20\x0c\x80\xe1\x7b\xa7\xf8\x07\xe0\xf2\x06\x78\x71\x16\x02\x35\x36\x29\xd6\x14\x0c\xeb\xfb\x4d\x75\x6d\x0b\xeb\x85\x0c\xd7\xc2\x53\xe7\xdc\x91\x5d\xe0\xcc\x18\xd4\xd6\xe2\xbd\x97\xb0\x2f\x4d\x45\x47\x35\xe7\xcf\x21\xb8\x0d\x5b\xfc\xe4\x0b\x00\x00\xff\xff\x18\xf5\x5a\xe8\x43\x00\x00\x00")

func sqlSelect_account_credentialsSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlSelect_account_credentialsSql,
		"sql/select_account_credentials.sql",
	)
}

func sqlSelect_account_credentialsSql() (*asset, error) {
	bytes, err := sqlSelect_account_credentialsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/select_account_credentials.sql", size: 67, mode: os.FileMode(420), modTime: time.Unix(1465500676, 0)}
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
	"sql/insert_account.sql": sqlInsert_accountSql,
	"sql/select_account_credentials.sql": sqlSelect_account_credentialsSql,
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
	"sql": &bintree{nil, map[string]*bintree{
		"insert_account.sql": &bintree{sqlInsert_accountSql, map[string]*bintree{}},
		"select_account_credentials.sql": &bintree{sqlSelect_account_credentialsSql, map[string]*bintree{}},
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

