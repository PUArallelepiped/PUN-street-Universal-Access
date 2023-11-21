package delivery

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Sorce struct {
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

func NewImgHandler(e *gin.Engine) {
	api_endpoint := "https://freeimage.host/api/1/upload"
	// path := "api/v1/test2"

	e.POST(api_endpoint, upload)
}

func upload(c *gin.Context) {
	encoded_string := read_string()

	data := Sorce{
		Key:    "6d207e02198a847aa98d0a2a901485a5",
		Action: "upload",
		Source: encoded_string,
		Format: "json",
	}

	c.JSON(200, data)

	var test ResponseInfo
	if err := c.BindJSON(&test); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(test)
}

func read_string() string {
	path := "/home/wj461/code/database_system/PUN-street-Universal-Access/backend/uploadimg/delivery/cat.txt"
	content, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
