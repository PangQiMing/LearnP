package routes

import "github.com/gin-gonic/gin"

import (
	"net/http"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.QiMing.com!",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	return r
}

//func main() {
//	r := gin.Default()
//	r.GET("/hello", helloHandler)
//	if err := r.Run(); err != nil {
//		fmt.Println("startup service failed, err:%v\n", err)
//	}
//}
