package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}

		if headers.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			return
		}

		if headers.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			return
		}

		c.SaveUploadedFile(headers, headers.Filename)
		c.String(http.StatusOK, headers.Filename)
	})
	r.Run()
}
