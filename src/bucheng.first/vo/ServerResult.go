package vo

import "github.com/gin-gonic/gin"

/**
快速成功
*/
func SuccessServerResult(data interface{}, c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "操作成功",
		"data":    data,
	})
}

/**
快速失败
*/
func FailServerResult(message string, c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    500,
		"message": message,
		"data":    nil,
	})
}
