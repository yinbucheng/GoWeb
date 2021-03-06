package main

import (
	"bucheng.first/controller"
	"bucheng.first/entity"
	"bucheng.first/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	var db *gorm.DB
	db = utils.InitDb("root", "123456", "go_test")
	entity.CreateTable(db, entity.User{})
	entity.CreateTable(db, entity.Room{})
	fmt.Println("================初始化表完成====================")
	engine := gin.Default()
	controller.InitController(engine)
	engine.Run(":9999")
}
