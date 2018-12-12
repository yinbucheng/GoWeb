package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type WebController struct {

}



func (controller WebController)HelloWord(engine *gin.Engine)  {
	engine.GET("/first/HelloWord",controller.helloWord)
}


func (controller WebController)helloWord(c *gin.Context){
	name:=c.Query("name")
	age:=c.Query("age")
	address:=c.DefaultQuery("address","")
	fmt.Println("name:",name,"age:"+age,"address:",address)
	c.JSON(200,gin.H{
		"status" :200,
		"error": "success",
		"data": nil,
	})
}