package shop

import "github.com/gin-gonic/gin"

func Router(engine *gin.Engine) {
	engine.GET("/shop", helloShop)
}
