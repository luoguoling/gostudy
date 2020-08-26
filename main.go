package main

import (
	"ginClass/mongolibs"
	"github.com/gin-gonic/gin"
)


func main() {
	mongolibs.GetList()
	mongolibs.OpMongo()
	r := gin.Default()
	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.Query("user")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			"success": "true",
			"id":id,
			"user":user,
			"pwd":pwd,
		})
	})
	r.POST("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user","aaaa")
		pwd := c.PostForm("pwd")
		c.JSON(200,gin.H{
			"success":true,
			"user":user,
			"pwd":pwd,
		})

	})
	r.PUT("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user","aaaa")
		pwd := c.PostForm("pwd")
		c.JSON(200,gin.H{
			"success":true,
			"user":user,
			"pwd":pwd,
		})

	})
	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080

}