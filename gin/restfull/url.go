package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//URL参数可以通过DefaultQuery()或Query()方法获取
//DefaultQuery()若参数不村则，返回默认值
//Query()若不存在，返回空串 API ? name=zs
func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Hello")
		c.String(http.StatusOK, fmt.Sprintf("is %s", name))
	})
	r.Run()
}
