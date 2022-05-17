package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func shopHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello shop!",
	})
}

func LoadShop(e *gin.Engine) {
	e.GET("/shop", shopHandler)
}
