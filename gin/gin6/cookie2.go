package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := c.Cookie("abc"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
			c.Abort()
			return
		}
		c.Next()
		return
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		c.String(200, "Login success!")
	})
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
	if err := r.Run(); err != nil {
		return
	}
}
