package main

import (
	controller2 "bucheng.first/controller"
	"bucheng.first/entity"
	"bucheng.first/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	var db *gorm.DB
	db = utils.InitDb("root", "introcks1234", "go_test")
	entity.CreateTable(db, entity.User{})
	entity.CreateTable(db, entity.Room{})
	entity.CreateTable(db, entity.ClassRoom{})
	fmt.Println("================初始化表完成====================")
	engine := gin.Default()
	controller2.InitWebController(engine)
	controller2.NewUserController(engine)
	controller2.InitClassController(engine)
	engine.Run(":9999")
}
