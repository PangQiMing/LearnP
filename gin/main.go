package main

import (
	"LearnP/gin/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//r := routes.SetupRouter()
	//if err := r.Run(); err != nil {
	//	fmt.Printf("startup service failed, err:%v\n", err)
	//})

	r := gin.Default()
	routes.LoadShop(r)
	routes.LoadBlog(r)
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err: %v", err)
	}
}
