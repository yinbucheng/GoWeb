package dao

import (
	"bucheng.first/entity"
	"github.com/jinzhu/gorm"
)

type UserDao interface {
	Save(user *entity.User, db *gorm.DB) int
	Update(user *entity.User, db *gorm.DB) int
	Delete(user *entity.User, db *gorm.DB) int
	FindOne(id int, db *gorm.DB) entity.User
	ListAll(db *gorm.DB) []entity.User
}
