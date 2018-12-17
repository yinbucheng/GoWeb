package impl

import (
	"bucheng.first/base/impl"
	impl2 "bucheng.first/dao/impl"
	"bucheng.first/utils"
	"fmt"
	"os/user"
)

type UserServiceImpl struct {
	impl.BaseServiceImpl
}

func NewServiceInstance() *UserServiceImpl {
	instance := new(UserServiceImpl)
	instance.Dao = impl2.UserDaoImpl{}
	instance.Db = utils.InitDb("root", "introcks1234", "go_test")
	fmt.Println("----->invoke UserServiceImpl attribute init<----")
	return instance
}

func (service UserServiceImpl) FindAll() []user.User {
	daoImpl := service.Dao.(impl2.UserDaoImpl)
	return daoImpl.FindAll(service.Db)
}
