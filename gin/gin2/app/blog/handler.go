package blog

import "github.com/gin-gonic/gin"

func helloBlog(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello blog!",
	})
}
