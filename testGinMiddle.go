package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
//中间件写法
func middle()gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("我在调用方法前")
		c.Next()
		fmt.Println("我在掉用方法后")
	}
}
func main()  {
	r := gin.Default()
	v1 := r.Group("v1").Use(middle())
	v1.GET("test", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"success":true,
		})

	})
	r.Run(":8002")
}
