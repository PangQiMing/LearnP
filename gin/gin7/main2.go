package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
)

type Person1 struct {
	Age     int    `form:"age" binding:"required,gt=10"`
	Name    string `form:"name" binding:"NotNullAndAdmin"`
	Address string `form:"address" binding:"required"`
}

func nameNotNullAndAdmin(v *validator.Validate, topStruct reflect.Value,
	currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if value, ok := field.Interface().(string); ok {
		return value != "" && !("51mh" == value)
	}
	return true
}

func main() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NotNullAndAdmin", nameNotNullAndAdmin)
	}
	r.GET("/51mh", func(c *gin.Context) {
		var person Person1
		if e := c.ShouldBind(&person); e == nil {
			c.String(http.StatusOK, "%v", person)
		} else {
			c.String(http.StatusOK, "person bind err :%v", e.Error())
		}
	})
	if err := r.Run(); err != nil {
		return
	}
}
