package controller

import (
	"bucheng.first/entity"
	"bucheng.first/service"
	"bucheng.first/service/impl"
	"bucheng.first/vo"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ClassController struct {
	service service.IClassRoomService
}

func InitClassController(engine *gin.Engine) {
	controller := new(ClassController)
	controller.service = impl.NewClassRoomService()
	engine.POST("/classRoom/save", controller.save)
}

func (controller ClassController) save(c *gin.Context) {
	var room entity.ClassRoom
	c.BindJSON(&room)
	fmt.Println(room)
	controller.service.Save(&room)
	vo.SuccessServerResult(nil, c)
}
