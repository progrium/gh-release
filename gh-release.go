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

func assert(err error) {
	if err != nil {
		println("!!", err.Error())
		os.Exit(2)
	}
}

func UploadUrl(args []string) int {
	bytes, err := ioutil.ReadAll(os.Stdin)
	assert(err)
	var release map[string]interface{}
	assert(json.Unmarshal(bytes, &release))
	url, ok := release["upload_url"].(string)
	if !ok {
		return 2
	}
	url = strings.Replace(url, "{", "", 1)
	url = strings.Replace(url, "}", "", 1)
	fmt.Println(url)
	return 0
}

func ReleaseIdFromTagname(args []string) int {
	tagname := args[0]
	bytes, err := ioutil.ReadAll(os.Stdin)
	assert(err)
	var releases []map[string]interface{}
	assert(json.Unmarshal(bytes, &releases))
	for _, release := range releases {
		if release["tag_name"].(string) == tagname {
			fmt.Println(strconv.Itoa(int(release["id"].(float64))))
			return 0
		}
	}
	return 2
}

func MimeType(args []string) int {
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
	return 0
}

func main() {
	basher.Application(map[string]func([]string) int{
		"upload-url":              UploadUrl,
		"release-id-from-tagname": ReleaseIdFromTagname,
		"mimetype":                MimeType,
	}, []string{
		"bash/gh-release.bash",
	}, Asset, true)
}
