package base

import "github.com/gin-gonic/gin"

func Fail(content string, c *gin.Context) {
	c.JSON(200, gin.H{
		"state": 500,
		"info":  content,
		"data":  nil,
	})
}

func Success(c *gin.Context) {
	c.JSON(200, gin.H{
		"state": 200,
		"info":  "operation success",
		"data":  nil,
	})
}

func SuccessWithData(data interface{}, c *gin.Context) {
	c.JSON(200, gin.H{
		"state": 200,
		"info":  "operation success",
		"data":  data,
	})
}
