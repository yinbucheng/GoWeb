package impl

import (
	"bucheng.first/base/impl"
	"bucheng.first/dao"
	impl2 "bucheng.first/dao/impl"
	"bucheng.first/entity"
	"bucheng.first/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"os/user"
)

type UserServiceImpl struct {
	impl.BaseServiceImpl
	roomDao dao.RoomDao
}

func NewServiceInstance() *UserServiceImpl {
	instance := new(UserServiceImpl)
	instance.Dao = impl2.UserDaoImpl{}
	instance.Db = utils.InitDb("root", "introcks1234", "go_test")
	instance.roomDao = impl2.RoomDaoImpl{}
	fmt.Println("----->invoke UserServiceImpl attribute init<----")
	return instance
}

func (service UserServiceImpl) FindAll() []user.User {
	daoImpl := service.Dao.(impl2.UserDaoImpl)
	return daoImpl.FindAll(service.Db)
}

func (service UserServiceImpl) AffairTest() {
	service.ExecuteOnAffair(func(params ...interface{}) {
		len := len(params)
		var db *gorm.DB
		db = params[len-1].(*gorm.DB)
		user := entity.User{
			Name:   "cuicui",
			Age:    20,
			Gender: "nv",
		}
		service.Dao.Save(&user, db)
		j := 0
		i := 1 / j
		fmt.Println(i)
		room := entity.Room{
			RoomName: "hehe",
			Address:  "zhonggu",
		}
		service.roomDao.Save(&room, db)
	}, nil)
}
