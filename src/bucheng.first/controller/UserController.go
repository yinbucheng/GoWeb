package controller

import (
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

func NewUserController(engine *gin.Engine) *UserController {
	userController := new(UserController)
	userController.userService = impl.NewServiceInstance()
	userController.SaveUser(engine)
	userController.FindOne(engine)
	userController.DeleteUser(engine)
	userController.ListAll(engine)
	userController.UpdateUser(engine)
	userController.SaveUserAndRoom(engine)
	return userController
}

func (userController UserController) SaveUser(engine *gin.Engine) {
	engine.POST("/user/save", userController.saveUser)
}

func (userController UserController) saveUser(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)
	result := userController.userService.Save(&user)
	if result >= 0 {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "完成操作",
		})
	} else {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "操作失败",
		})
	}
}

func (userController UserController) DeleteUser(engine *gin.Engine) {
	engine.POST("/user/delete", userController.deleteUser)
}

func (userController UserController) deleteUser(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)
	result := userController.userService.Delete(&user)
	if result >= 0 {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "完成操作",
		})
	} else {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "操作失败",
		})
	}
}

func (userController UserController) ListAll(engine *gin.Engine) {
	engine.POST("/user/listAll", userController.listAll)
}

func (userController UserController) listAll(c *gin.Context) {
	users := userController.userService.ListAll()
	c.JSON(200, gin.H{
		"code":    200,
		"message": "操作成功",
		"data":    users,
	})
}

func (userController UserController) FindOne(engine *gin.Engine) {
	engine.GET("/user/findOne", userController.findOne)
}

func (userController UserController) findOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("Error:", err)
	}
	user := userController.userService.FindOne(id)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "操作成功",
		"data":    user,
	})
}

func (userController UserController) UpdateUser(engine *gin.Engine) {
	engine.POST("/user/update", userController.updateUser)
}

func (userController UserController) updateUser(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)
	result := userController.userService.Update(&user)
	if result >= 0 {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "完成操作",
		})
	} else {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "操作失败",
		})
	}
}

func (userController UserController) SaveUserAndRoom(engine *gin.Engine) {
	engine.POST("/user/saveUserAndRoom", userController.saveUserAndRoom)
}

func (userController UserController) saveUserAndRoom(c *gin.Context) {
	result := userController.userService.SaveUserAndRoom()
	if result != -1 {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "添加学生和Room成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "添加学生和Room失败",
		})
	}
}
