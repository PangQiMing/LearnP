package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/longAsync", func(c *gin.Context) {
		copyContext := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Printf("异步执行:" + copyContext.Request.URL.Path)
		}()
	})

	r.GET("/logSync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行:" + c.Request.URL.Path)
	})
	if err := r.Run(); err != nil {
		return
	}
}
