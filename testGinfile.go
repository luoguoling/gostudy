package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.POST("testfile", func(c *gin.Context) {
		file,_ := c.FormFile("file")
		c.SaveUploadedFile(file,"./"+file.Filename)
		c.Writer.Header().Add("Content-Disposition",fmt.Sprintf("attachment:filename=%s",file.Filename))
		c.File("./"+file.Filename)
		c.JSON(200,gin.H{
			"msg":file,
		})

	})
	r.Run(":8002")
}
