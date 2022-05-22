package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()

	since := time.Since(start)
	fmt.Println("程序用时:", since)
}

func main() {
	r := gin.Default()
	r.Use(myTime)
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	if err := r.Run(); err != nil {
		return
	}
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
