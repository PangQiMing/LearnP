package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//Person
type Person struct {
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"'`
}

func main() {
	r := gin.Default()
	r.GET("/51mh", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, fmt.Sprintf("%#v", err))
			return
		}
		c.String(200, fmt.Sprintf("%#v", person))
	})
	if err := r.Run(); err != nil {
		return
	}
}
