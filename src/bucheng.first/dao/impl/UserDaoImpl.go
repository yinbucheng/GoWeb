package impl

import (
	"bucheng.first/entity"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserDaoImpl struct {
}

func (userDao UserDaoImpl) Save(user *entity.User, db *gorm.DB) int {
	err := db.Create(user).Error
	fmt.Println(err)
	if err != nil {
		return -1
	}
	return 0
}

func (userDao UserDaoImpl) Delete(user *entity.User, db *gorm.DB) int {
	err := db.Where("id=?", user.ID).Delete(user).Error
	if err != nil {
		return -1
	}
	return 0
}

func (userDao UserDaoImpl) Update(user *entity.User, db *gorm.DB) int {
	err := db.Model(user).Update("name", user.Name).Error
	if err != nil {
		return -1
	}
	return 0
}

func (userDao UserDaoImpl) FindOne(id int, db *gorm.DB) entity.User {
	var user entity.User
	db.Where("id=?", id).First(&user)
	return user
}

func (userDao UserDaoImpl) ListAll(db *gorm.DB) []entity.User {
	var users []entity.User
	db.Find(&users)
	return users
}
