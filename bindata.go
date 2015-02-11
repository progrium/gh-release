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

var _bash_gh_release_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\xdf\x8f\xe2\x36\x10\x7e\x86\xbf\x62\x64\xa5\x7b\x4b\x55\x13\xf5\xc7\x13\xec\x6e\x6f\xbb\x87\x7a\xa7\x56\xb7\xd5\xc1\xb5\x27\x51\x84\x4c\xe2\x10\xf7\x12\x3b\x8a\x0d\x2a\xe2\xf8\xdf\x3b\x76\xec\x10\x58\x1e\xa8\xee\x89\xd8\x33\xf3\xcd\x37\x33\xdf\x98\x7e\xcd\x59\xaa\x64\xb1\x83\x9a\x67\x4b\x2e\xd3\x4a\x09\x69\xee\x49\x6e\x4c\xa5\x47\x71\xcc\x2a\x31\x5c\x0b\x93\x6f\x56\xc3\x44\x95\x71\xcd\x2b\xa5\xe3\x6f\x74\x8c\x77\x78\xc8\x74\x6c\xd8\xda\x5e\x90\x2e\x50\xc1\x99\xe6\xff\x03\xcc\x47\x5c\x02\xf9\x47\x2b\x79\xff\x6a\x4f\x30\xcd\x52\xb2\x92\x93\x11\x90\x2d\xa6\xfb\x0e\x48\x38\x36\x27\xc3\xea\x35\x37\x4b\xc4\x2d\x85\x11\x3a\xf7\x96\xc3\xab\x7e\xdf\x63\xd1\x04\xd1\x0d\xbf\x1d\xc0\xbe\xdf\x4b\x79\x52\xb0\x9a\x83\xe5\x60\x81\xee\x49\xf4\x3d\x81\x2d\xaf\xb5\xc0\x84\x24\xfa\x81\xc0\xaa\x66\x32\xc9\xf1\x7b\xff\xe3\x88\x96\x4c\x1b\x5e\x1f\x08\x78\xe7\x9f\x48\xbf\x57\xa8\x84\x15\x81\x29\xde\xdd\x56\x35\x96\x9b\x01\x89\xba\xec\x09\x9e\x3d\xae\xfd\x74\xac\xf1\xb7\x41\x27\x83\x73\x9c\xe5\xa6\x2e\x2e\x62\x85\x76\x12\x77\xd7\xb0\x76\xe1\x3c\xc9\x15\x90\x27\x5b\x9c\x90\xeb\x80\x03\xdb\x90\x15\xb2\x5a\x95\xbe\x1a\xf0\x79\x61\x38\x1c\x62\xe8\xa6\x2a\x14\x4b\x43\xca\x04\x7f\x81\x6a\xa0\xe9\x31\x2d\xe9\x30\x40\xf3\xcf\x2c\x49\xb8\xd6\x4b\xa3\x3e\x73\x79\x1f\xfd\xfa\x6e\xf6\xf6\xe3\x2f\xcb\xc7\xa7\xa7\xc9\x74\xba\x9c\x3d\xff\x36\x79\x4f\xe0\x0b\x34\xb0\x14\xfd\x2d\xbf\x4c\xd5\xc0\xb4\xe6\x06\x84\x84\xe8\xb6\xc0\x04\x8f\x81\xe5\x60\x0c\xa9\xea\xf7\x7c\x0b\x7c\x6f\x6f\x57\x68\xb1\xdf\x10\xb9\x38\x0b\xe2\xab\xfc\xe8\x90\x6d\x99\xae\x91\xbe\x8c\x5e\xc3\xfc\x13\xfc\xf1\x3c\x9d\x01\x7d\x8b\xdd\x50\xd2\x70\x69\xe8\x6c\x57\xf1\x11\x26\x2d\x45\xc9\x0d\x7e\x37\x61\x03\x02\x94\xa6\xcc\x30\xba\x12\x92\xd5\x3b\x20\xaf\x3d\x9f\xb8\xc9\x48\xe0\x6f\x44\xed\x91\xa8\xd3\x20\x17\x79\x73\x45\xfd\x0f\x10\xa7\x7c\x1b\xcb\x4d\x51\xa0\xce\x94\xe4\xfd\xc3\x51\x83\x29\xd7\xa6\x56\xbb\x6b\x45\xf8\xd5\xe2\x08\x0e\x22\xed\x4e\xf8\x64\xa6\x76\x62\x81\x9e\x48\xa9\x55\x0b\xc5\x75\x73\xfd\x25\xad\x8a\x3a\x52\x7b\x83\xce\x5d\xa9\x35\x43\x68\xd5\xf3\x09\xde\x4c\x7e\x9f\xcc\x26\xa7\x59\xe2\xe8\x48\xe5\x0a\x15\xbd\xc8\x85\x8c\x9a\x3c\xf6\x25\x78\xd9\x86\xec\x72\x0b\xce\x0b\xb8\x44\xd2\x03\x5e\x43\x0a\x07\xb9\xd1\x6c\xed\xdf\x10\x2f\x49\x7b\x31\x82\x75\x4e\xc3\xe6\xcd\xe9\x76\x01\x7a\xb3\xb2\x6f\x11\x93\x69\xa8\x25\x04\x4c\x5b\x8b\x1e\xb5\x75\x02\x34\xaf\x13\xdc\x05\xee\x0f\x70\xe7\x99\x3f\xc0\xbc\xd9\xda\x05\xcc\xad\x65\xd1\x89\xf2\x7a\xba\x18\xd6\x05\xcf\x79\xf2\x59\x6f\x4a\x0d\x77\xac\x58\xab\x1a\x1f\xe1\x32\xd8\xbb\xf2\x6c\xfd\x4e\x05\x8a\x31\x4e\x9b\x01\xf0\x2f\x04\x70\x4b\x88\x86\x16\x1b\x32\x51\x70\xdd\x0c\xe9\x9a\x9d\x4f\x98\x01\x72\xbe\x77\x5f\x8e\x70\xc4\xa2\xdb\x75\x3a\x3a\xed\x9d\xd7\x61\xe8\x2c\xc7\xe5\x2a\x99\x90\x0d\x61\x9b\x91\x72\x05\x95\xa8\x78\xc6\x44\x31\x86\xf9\x1c\x81\x66\x1f\x1e\x9f\x26\x04\x16\x0b\xb8\xb9\x01\xe7\xf3\x2f\x6a\xc1\xce\xca\x6d\x9c\x90\x96\x8e\xeb\xff\xa0\xd7\xd3\xb9\xc8\xcc\x18\x4e\xff\x35\xd0\xf1\x35\x19\x8f\xd1\xcf\x77\x7c\x70\xee\x17\x26\xd1\x3a\xb6\xbd\x7c\xe1\x7a\x9c\x46\xeb\x4c\xb7\x98\xd9\x77\x37\xfa\x73\xf2\x61\xfa\xee\xf9\x7d\x63\xf9\xd6\x1a\x9c\xea\xec\x91\x6b\x96\x60\xc9\xff\x05\x00\x00\xff\xff\xf9\xa6\xd5\xeb\xb7\x07\x00\x00")

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

	info := bindata_file_info{name: "bash/gh-release.bash", size: 1975, mode: os.FileMode(493), modTime: time.Unix(1423621139, 0)}
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

