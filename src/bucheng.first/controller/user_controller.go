package controller

import (
	"bucheng.first/base"
	"bucheng.first/entity"
	"bucheng.first/service"
	"bucheng.first/service/impl"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
	userService service.IUserService
}

func InitController(engine *gin.Engine) {
	userController := new(UserController)
	userController.userService = impl.NewServiceInstance()
	engine.POST("/user/save", userController.saveUser)
	engine.POST("/user/delete", userController.deleteUser)
	engine.POST("/user/listAll", userController.listAll)
	engine.GET("/user/findOne", userController.findOne)
	engine.POST("/user/update", userController.updateUser)
	engine.POST("/user/findAll", userController.findAll)
	engine.POST("/user/affairTest", userController.affairTest)
}

func (userController UserController) saveUser(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)
	code := userController.userService.Save(&user)
	if code != -1 {
		base.Success(c)
	} else {
		base.Fail("添加失败", c)
	}
}

func (userController UserController) deleteUser(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)
	result := userController.userService.Delete(user.ID, &user)
	if result >= 0 {
		base.Success(c)
	} else {
		base.Fail("添加失败", c)
	}
}

func (userController UserController) listAll(c *gin.Context) {
	var users []entity.User
	userController.userService.ListAll(&users)
	base.SuccessWithData(users, c)
}

func (userController UserController) findOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("Error:", err)
	}
	var user entity.User
	userController.userService.FindOne(id, &user)
	base.SuccessWithData(user, c)
}

func (userController UserController) UpdateUser(engine *gin.Engine) {
	engine.POST("/user/update", userController.updateUser)
}

func (userController UserController) updateUser(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)
	result := userController.userService.Update(user.ID, &user)
	if result >= 0 {
		base.Success(c)
	} else {
		base.Fail("更新失败", c)
	}
}

func (controller UserController) findAll(c *gin.Context) {
	users := controller.userService.FindAll()
	base.SuccessWithData(users, c)
}

func (controller UserController) affairTest(c *gin.Context) {
	controller.userService.AffairTest()
	base.Success(c)
}
