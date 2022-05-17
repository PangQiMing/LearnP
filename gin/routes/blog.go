package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloBlog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello blog!",
	})
}

func LoadBlog(e *gin.Engine) {
	e.GET("/blog", helloBlog)
}
