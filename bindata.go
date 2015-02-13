package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _bash_gh_release_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\x51\x6f\xdb\x36\x10\x7e\xb6\x7f\xc5\x81\xf3\xd2\x78\x18\x2d\xac\xdb\x93\x9d\x64\xcd\x52\x63\x0d\x56\x34\x45\xec\x6c\x05\x3c\x43\xa0\x25\xca\xe2\x2a\x91\x82\x48\x1b\x0b\x52\xff\xf7\x1e\x29\x52\x56\x1c\x3f\x18\xc8\x93\x45\xde\xdd\x77\xdf\xdd\x7d\x47\xf7\x6b\xce\x52\x25\x8b\x47\xa8\x79\x16\x73\x99\x56\x4a\x48\x73\x49\x06\x4f\x7f\xde\xce\x3f\x3c\xfc\x11\x5f\x7f\xbe\x8d\x1f\xee\x3f\x8e\x69\x6e\x4c\xa5\xc7\x51\xc4\x2a\x31\x5a\x0b\x93\x6f\x56\xa3\x44\x95\xbb\xa8\xe6\x95\xd2\xd1\x8f\x3a\xc2\x4b\x3c\x64\x3a\x32\x6c\x6d\x2f\x48\x17\xbb\xe0\x4c\xf3\xd7\xe1\x7b\x90\x63\xb8\xff\x69\x25\x2f\xdf\x3c\x11\xcc\x1c\x4b\x56\x72\x32\x06\xb2\x45\x06\x3f\x03\x09\xc7\xe6\x64\x58\xbd\xe6\x26\x46\xe0\x52\x18\xa1\x73\x6f\xd9\xbd\xe9\xf7\x3d\x16\x4d\x10\xdd\xf0\xf3\x21\x3c\xf5\x7b\x29\x4f\x0a\x56\x73\xb0\x1c\x2c\x10\xf2\xfe\x85\xc0\x96\xd7\x5a\x60\x42\x2c\xe2\xed\x0f\xdb\x1d\x81\x55\xcd\x64\x92\xdb\xf3\xaf\x63\x5a\x32\x6d\x78\x8d\xb7\x3e\xe0\x37\xd2\xef\x15\x2a\x61\x45\x60\x8b\x77\xe7\x55\x8d\x5d\xc8\x80\x0c\xba\x15\x10\x3c\x7b\x6c\xfb\xe9\x98\xe3\x6f\x83\x4e\x86\x87\x38\xf1\xa6\x2e\x8e\x62\x85\x2e\x13\x77\xd7\x30\x77\xe1\x3c\xc9\x15\x90\x1b\x5b\xa0\x90\xeb\x80\x03\xdb\x90\x15\xb2\x5a\x95\xbe\x1a\xf0\x79\x61\x34\x1a\x61\xe8\xa6\x2a\x14\x4b\x43\xca\x04\x7f\x81\x6a\xa0\xe9\x3e\x2d\xe9\x30\x40\xf3\xef\x2c\x49\xb8\xd6\xb1\x51\x5f\xb9\xbc\x1c\x84\x69\xdf\xdc\x4c\x67\xb3\x78\x7e\xf7\xd7\xf4\x13\x81\x6f\xd0\xc0\x52\xf4\xb7\xfc\x32\x55\x03\xd3\x9a\x1b\x10\x12\x06\xe7\x05\x26\xb8\x0e\x2c\x87\x13\x48\x55\xbf\xe7\x5b\xe0\x7b\x7b\xbe\x42\x8b\xfd\x86\x81\x8b\xb3\x20\xbe\xca\x07\x87\x6c\xcb\x74\x8d\xf4\x65\xf4\x1a\xe6\x5f\xe0\xf3\xdd\x6c\x0e\xf4\x03\x76\x43\x49\xc3\xa5\xa1\xf3\xc7\x8a\x8f\x31\x69\x29\x4a\x6e\xf0\xbb\x09\x1b\x12\xa0\x34\x65\x86\xd1\x95\x90\xac\x7e\x04\xf2\xce\xf3\x89\x9a\x8c\x04\xfe\x45\xd4\x1e\x19\x74\x1a\xe4\x22\xcf\x4e\xa8\xff\x0a\xa2\x94\x6f\x23\xb9\x29\x0a\xd4\x9a\x92\xbc\xbf\xdb\xeb\x30\xe5\xda\xd4\xea\xf1\x34\x21\xbe\x7d\xbd\x38\x82\x83\x48\xbb\x13\x7e\x36\x53\x3b\xb1\x40\x4f\xa4\xd4\xaa\x85\xe2\xca\xb9\xfe\x92\x56\x45\x1d\xa9\xbd\x47\xe7\xae\xd4\x9a\x21\xb4\xea\xf9\x02\xef\xa7\x1f\xa7\xf3\xe9\xf3\x2c\xd1\x60\x4f\xe5\x04\x15\xbd\xc8\x85\x8c\x9a\x3c\xf6\x35\x78\xd9\x86\xec\x78\x0b\x0e\x0b\x38\x46\xd2\x03\x9e\x42\x0a\x07\xb9\xd1\x6c\xed\xdf\x11\x2f\x49\x7b\x31\x86\x75\x4e\xc3\xe6\x2d\xe8\x76\x09\x7a\xb3\xb2\xef\x11\x93\xa9\xaf\x25\xf8\xcf\x5a\x83\x1e\xb7\x65\x02\x34\x0f\x14\x5c\x04\xea\x57\x70\xe1\x89\x5f\xc1\xa2\x59\xda\x25\x2c\xac\x65\xd9\x89\xf2\x72\x3a\x1a\xd6\x05\xcf\x79\xf2\x55\x6f\x4a\x0d\x17\xac\x58\xab\x1a\x1f\xe2\x32\xd8\xbb\xea\x6c\xfd\x9e\xeb\x13\x63\x9c\x34\x03\xe0\x3f\x08\xe0\x76\x10\x0d\x2d\x36\x64\xa2\xe0\xba\x99\xd1\x29\x2b\x9f\x30\x03\xe4\x70\xed\xbe\xed\xe1\x88\x45\xb7\xdb\xb4\x77\x7a\x72\x5e\xbb\x91\xb3\xec\x77\xab\x64\x42\x36\x84\x6d\x46\xca\x15\x54\xa2\xe2\x19\x13\xc5\x04\x16\x0b\x04\x9a\xdf\x5f\xdf\x4c\x09\x2c\x97\x70\x76\x06\xce\xe7\x7f\x94\x82\x1d\x95\x5b\x38\x21\x2d\x1d\xd7\xff\x61\xaf\xa7\x73\x91\x99\x09\x3c\xff\xe3\x40\xc7\x77\x64\x32\x41\x3f\xdf\xf1\xe1\xa1\x5f\x98\x44\xeb\xd8\xf6\xf2\x85\xeb\x7e\x1a\xad\x33\xdd\x62\x66\xdf\xdd\xc1\xdf\xd3\xfb\xd9\xed\xdd\xa7\xc6\xf2\x93\x35\x38\xd1\xd9\x23\xd7\x2c\xc1\x92\xbf\x07\x00\x00\xff\xff\xbf\x6e\xb7\x30\xe0\x07\x00\x00")

func bash_gh_release_bash_bytes() ([]byte, error) {
	return bindata_read(
		_bash_gh_release_bash,
		"bash/gh-release.bash",
	)
}

func bash_gh_release_bash() (*asset, error) {
	bytes, err := bash_gh_release_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "bash/gh-release.bash", size: 2016, mode: os.FileMode(493), modTime: time.Unix(1423787091, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"bash/gh-release.bash": bash_gh_release_bash,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"bash": &_bintree_t{nil, map[string]*_bintree_t{
		"gh-release.bash": &_bintree_t{bash_gh_release_bash, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

