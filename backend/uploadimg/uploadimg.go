package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type PostInfo struct {
	Key    string `json:"key"`
	Action string `json:"action"`
	Source string `json:"source"`
	Format string `json:"format"`
}

type ResponseInfo struct {
	StatusCode int       `json:"status_code"`
	Image      ImageInfo `json:"image"`
}

type ImageInfo struct {
	Url       string `json:"url"`
	UrlViewer string `json:"url_viewer"`
}

func main() {
	test2()
}

func read_string() string {
	path := "/home/wj461/code/database_system/PUN-street-Universal-Access/backend/uploadimg/delivery/cat.txt"
	content, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func test2() {
	httpposturl := "https://freeimage.host/api/1/upload"
	// httpposturl = "http://localhost:8888/"
	s := PostInfo{
		Key:    "6d207e02198a847aa98d0a2a901485a5",
		Action: "upload",
		Source: read_string(),
		Format: "json"}
	v := make(map[string][]string)
	v["key"] = []string{s.Key}
	v["action"] = []string{s.Action}
	v["source"] = []string{s.Source}
	v["format"] = []string{s.Format}

	qs := url.Values(v)
	response, err := http.PostForm(httpposturl, qs)
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
	}
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}
