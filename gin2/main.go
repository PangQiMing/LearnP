package main

import (
	"LearnP/gin2/app/blog"
	"LearnP/gin2/app/shop"
	"LearnP/gin2/routes"
	"fmt"
)

func main() {
	routes.Include(shop.Router, blog.Router)
	r := routes.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err :%v\n", err)
	}
}
