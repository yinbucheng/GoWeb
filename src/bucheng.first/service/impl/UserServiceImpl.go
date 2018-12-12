package impl

import (
	"bucheng.first/dao"
	"bucheng.first/dao/impl"
	"bucheng.first/entity"
	"bucheng.first/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserServiceImpl struct {
	userDao dao.UserDao
	roomDao dao.RoomDao
	db      *gorm.DB
}

func NewServiceInstance() *UserServiceImpl {
	instance := new(UserServiceImpl)
	instance.userDao = new(impl.UserDaoImpl)
	instance.roomDao = new(impl.RoomDaoImpl)
	instance.db = utils.InitDb("root", "introcks1234", "go_test")
	return instance
}

func (userService UserServiceImpl) Save(user *entity.User) int {
	return userService.userDao.Save(user, userService.db)
}

func (userService UserServiceImpl) Delete(user *entity.User) int {
	return userService.userDao.Delete(user, userService.db)
}

func (userService UserServiceImpl) Update(user *entity.User) int {
	return userService.userDao.Update(user, userService.db)
}

func (userService UserServiceImpl) FindOne(id int) entity.User {
	return userService.userDao.FindOne(id, userService.db)
}

func (userService UserServiceImpl) ListAll() []entity.User {
	return userService.userDao.ListAll(userService.db)
}

func (userService UserServiceImpl) SaveUserAndRoom() int {
	var tx *gorm.DB
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			if tx != nil {
				tx.Rollback()
			}
		}
	}()
	//开启是否
	tx = userService.db.Begin()
	user := entity.User{
		Name:   "yinchong10",
		Age:    30,
		Gender: "nan",
	}
	result := userService.userDao.Save(&user, tx)
	if result == -1 {
		tx.Rollback()
		return -1
	}
	room := entity.Room{
		RoomName: "home2",
		Address:  "wuhai",
	}
	userService.roomDao.Save(&room, tx)
	if result == -1 {
		tx.Rollback()
		return -1
	}
	tx.Commit()
	return 0
}
