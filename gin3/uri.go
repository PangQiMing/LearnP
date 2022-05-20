package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login2 struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:user/:password", func(c *gin.Context) {
		var login Login2
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run()
}
