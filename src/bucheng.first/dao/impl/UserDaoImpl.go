package impl

import (
	"bucheng.first/base/impl"
	"github.com/jinzhu/gorm"
	"os/user"
)

type UserDaoImpl struct {
	impl.BaseDaoImpl
}

func (dao UserDaoImpl) FindAll(db *gorm.DB) []user.User {
	var users []user.User
	err := db.Find(&users).Error
	if err == nil {
		return users
	} else {
		return nil
	}
}
