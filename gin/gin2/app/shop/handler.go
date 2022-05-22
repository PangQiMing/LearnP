package shop

import "github.com/gin-gonic/gin"

func helloShop(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello shop!",
	})
}
