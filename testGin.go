package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type PostParams struct {
	Name string `json:"name" binding:"required"`
	Age int `json:"age" binding:"required,mustBig"`
	Sex bool `json:"sex"  binding:"required"`
}
func mustBig(f1 validator.FieldLevel) bool{
	if  (f1.Field().Interface().(int) <= 18){
		return false
	}
	return true
}
func main(){
	r := gin.Default()
	if v,ok := binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("mustBig",mustBig)
	}
	r.POST("/testBind", func(c  *gin.Context) {
		var p PostParams
		err := c.ShouldBindJSON(&p)
		fmt.Println(err)
		if err != nil{
			fmt.Println(err.Error())
			c.JSON(500,gin.H{
				"msg":"error",
				"data":gin.H{"error":"error"},
			})
		}else{
			c.JSON(200,gin.H{
				"msg":"success",
				"data":p.Age,
			})
		}


	})
	r.Run(":8001")
}
