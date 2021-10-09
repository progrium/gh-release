package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/progrium/go-basher"
)

var Version string

func assert(err error) {
	if err != nil {
		println("!!", err.Error())
		os.Exit(10)
	}
}

func fatal(msg string) {
	println("!!", msg)
	os.Exit(11)
}

func UploadUrl(args []string) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	assert(err)
	var release map[string]interface{}
	assert(json.Unmarshal(bytes, &release))
	url, ok := release["upload_url"].(string)
	if !ok {
		println("!! could not find upload_url")
		if os.Getenv("DEBUG") != "" {
			fmt.Printf("!! response: %s\n", string(bytes))
		}
		os.Exit(12)
	}
	i := strings.Index(url, "{")
	if i > -1 {
		url = url[:i]
	}
	i = strings.Index(url, "?")
	if i > -1 {
		url = url[:i]
	}
	fmt.Println(url + "?name")
}

func ReleaseIdFromTagname(args []string) {
	tagname := args[0]
	bytes, err := ioutil.ReadAll(os.Stdin)
	assert(err)
	var releases []map[string]interface{}
	assert(json.Unmarshal(bytes, &releases))
	for _, release := range releases {
		if release["tag_name"].(string) == tagname {
			fmt.Println(strconv.Itoa(int(release["id"].(float64))))
			return
		}
	}
	println(fmt.Sprintf("!! no tag_name matching %s found in releases", tagname))
	os.Exit(13)
}

func MimeType(args []string) {
	filename := args[0]
	ext := filepath.Ext(filename)
	mime.AddExtensionType(".gz", "application/gzip")
	mime.AddExtensionType(".tgz", "application/gzip")
	mime.AddExtensionType(".tar", "application/tar")
	mime.AddExtensionType(".zip", "application/zip")
	mimetype := mime.TypeByExtension(ext)
	if mimetype != "" {
		fmt.Println(mimetype)
	} else {
		fmt.Println("application/octet-stream")
	}
}

func Checksum(args []string) {
	if len(args) < 1 {
		fatal("No algorithm specified")
	}
	var h hash.Hash
	switch args[0] {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	default:
		fatal("Algorithm '" + args[0] + "' is unsupported")
	}
	io.Copy(h, os.Stdin)
	fmt.Printf("%x\n", h.Sum(nil))
}

func main() {
	os.Setenv("VERSION", Version)
	basher.Application(map[string]func([]string){
		"upload-url":              UploadUrl,
		"release-id-from-tagname": ReleaseIdFromTagname,
		"mimetype":                MimeType,
		"checksum":                Checksum,
	}, []string{
		"bash/gh-release.bash",
	}, Asset, true)
}
