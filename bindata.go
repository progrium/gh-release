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

var _bash_gh_release_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x54\x5d\x4f\xdb\x3c\x14\xbe\x6e\x7f\xc5\x91\x95\x17\xda\x57\x73\xa3\x7d\x5c\xb5\x94\xc1\x4a\x35\xd0\x26\x98\x68\x99\x90\xba\xaa\x72\x13\xa7\xf1\x96\xd8\x91\xed\x56\xab\x80\xff\xbe\xe3\xc4\x81\x00\xbd\xe8\xb4\xab\xc4\xf6\xf1\xf3\x71\xfc\xd8\x6d\xcd\x59\xac\x64\xb6\x05\xcd\x93\x05\x97\x71\xa1\x84\xb4\x43\x92\x5a\x5b\x98\x7e\x18\xb2\x42\xf4\x56\xc2\xa6\xeb\x65\x2f\x52\x79\xa8\x79\xa1\x4c\xf8\x9f\x09\x71\x0e\x07\x89\x09\x2d\x5b\xb9\x09\xd2\x04\xca\x38\x33\xfc\x2f\xc0\xfc\x8e\x5d\x20\x3f\x8d\x92\xc3\xc3\x3b\x82\x34\x0b\xc9\x72\x4e\xfa\x40\x36\x48\xf7\x06\x48\x3d\xac\x46\x96\xe9\x15\xb7\x0b\xc4\xcd\x85\x15\x26\xf5\x2b\x0f\x87\xed\xb6\xc7\xa2\x11\xa2\x5b\xde\xe9\xc2\x5d\xbb\x15\xf3\x28\x63\x9a\x83\xd3\xe0\x80\x86\x24\x78\x4b\x60\xc3\xb5\x11\x48\x48\x82\x77\x04\x96\x9a\xc9\x28\xc5\xff\xbb\xf7\x7d\x9a\x33\x63\xb9\x7e\x20\xe0\x8b\x3f\x90\x76\x2b\x53\x11\xcb\x6a\xa5\x38\xd7\x29\x34\xda\x4d\x80\x04\x4d\xf5\x04\xc7\x1e\xd7\xfd\x96\xaa\xf1\x5b\xa1\x93\xee\x4b\x9c\xc5\x5a\x67\x3b\xb1\xea\x76\x92\x72\xae\x52\x5d\x6e\xe7\x51\xaa\x80\x8c\x9c\x39\x21\x57\x35\x0e\x6c\x6a\x56\x48\xb4\xca\xbd\x1b\xf0\xbc\xd0\xeb\xf5\x70\xeb\xba\xc8\x14\x8b\x6b\xca\x08\xbf\x40\x0d\xd0\xf8\x89\x96\x34\x14\xe0\xf2\x47\x16\x45\xdc\x98\x85\x55\xbf\xb8\x1c\x06\x9f\x2f\xa6\xe7\x37\x9f\x16\xa7\xa3\xd1\x78\x32\x59\x4c\xaf\xbe\x8c\x2f\x09\xdc\x43\x05\x4b\xb1\xde\xe9\x4b\x94\x06\x66\x0c\xb7\x20\x24\x04\x9d\x0c\x09\x4e\x6b\x95\xdd\x01\xc4\xaa\xdd\xf2\x2d\xf0\xbd\xed\x2c\x71\xc5\xfd\x43\x50\xee\x73\x20\xde\xe5\x4d\x89\xec\x6c\x96\x8d\xf4\x36\x5a\x95\xf2\x5b\xf8\x76\x35\x99\x02\x3d\xc7\x6e\x28\x69\xb9\xb4\x74\xba\x2d\x78\x1f\x49\x73\x91\x73\x8b\xff\xd5\xb6\x2e\x01\x4a\x63\x66\x19\x5d\x0a\xc9\xf4\x16\xc8\x89\xd7\x13\x56\x8c\x04\x7e\x20\x6a\x8b\x04\x8d\x06\x95\x3b\x0f\xf6\xf0\x7f\x0c\x61\xcc\x37\xa1\x5c\x67\x19\xe6\x4c\x49\xde\x7e\x78\xca\x60\xcc\x8d\xd5\x6a\xbb\x6f\x08\xff\x39\x1c\x75\x81\x88\x9b\x27\xfc\xec\x4c\xdd\x89\xd5\xf2\x44\x4c\x5d\x5a\x28\x5e\xb7\xb2\xbf\xe4\x31\x45\x8d\xa8\x9d\x61\x71\x33\x6a\xd5\x21\x3c\xa6\xe7\x16\xce\xc6\x5f\xc7\xd3\xf1\x73\x96\x30\x78\x92\xb2\x47\x8a\x5e\x71\xa1\xa2\x8a\xc7\xbd\x04\xaf\xdb\x90\xec\x6e\xc1\x4b\x03\xbb\x44\x7a\xc0\x7d\x44\xe1\x41\xae\x0d\x5b\xf9\x37\xc4\x47\xd2\x4d\xf4\x61\x95\xd2\xfa\xe6\xcd\xe8\x66\x0e\xd5\x63\x73\xef\xcf\x1b\x8e\x6a\x49\xc7\x70\xe4\x05\x1d\xc3\xac\xba\x8c\x73\x98\xb9\x95\xb9\x37\xed\x58\x72\x26\x64\x45\xe2\xee\x0d\xe5\x0a\x0a\x51\xf0\x84\x89\x6c\x00\xb3\x19\x8a\x9e\x5e\x9f\x8e\xc6\x04\xe6\x73\x38\x38\x80\xb2\xe6\x37\xba\x73\xec\x65\x86\x84\x74\xb7\xa2\x94\xd0\x6d\xb5\x4c\x2a\x12\x3b\x80\xe7\xef\x20\x16\x9e\x90\xc1\x00\xeb\xbc\xc6\xee\xcb\xba\x5a\xfb\x63\x21\xdd\x20\x98\xb7\x1d\x7c\x1f\x5f\x4f\x2e\xae\x2e\xab\x95\xff\xdd\x42\xd9\x1a\x37\xe4\x86\x45\xe8\xe2\x4f\x00\x00\x00\xff\xff\x61\xb7\x2d\xa9\x5c\x06\x00\x00")

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

	info := bindata_file_info{name: "bash/gh-release.bash", size: 1628, mode: os.FileMode(493), modTime: time.Unix(1417368897, 0)}
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

