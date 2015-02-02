package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"os"
	"strconv"
	"strings"

	"github.com/progrium/go-basher"
)

var Version string

func assert(err error) {
	if err != nil {
		println("!!", err.Error())
		os.Exit(2)
	}
}

func UploadUrl(args []string) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	assert(err)
	var release map[string]interface{}
	assert(json.Unmarshal(bytes, &release))
	url, ok := release["upload_url"].(string)
	if !ok {
		os.Exit(2)
	}
	url = strings.Replace(url, "{", "", 1)
	url = strings.Replace(url, "}", "", 1)
	fmt.Println(url)
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
	os.Exit(2)
}

func MimeType(args []string) {
	filename := args[0]
	ext := filename[strings.LastIndex(filename, "."):]
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

func main() {
	os.Setenv("VERSION", Version)
	basher.Application(map[string]func([]string){
		"upload-url":              UploadUrl,
		"release-id-from-tagname": ReleaseIdFromTagname,
		"mimetype":                MimeType,
	}, []string{
		"bash/gh-release.bash",
	}, Asset, true)
}
