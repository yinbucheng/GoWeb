package dao

import (
	"bucheng.first/entity"
	"github.com/jinzhu/gorm"
)

type RoomDao interface {
	Save(room *entity.Room, db *gorm.DB) int
	Update(room *entity.Room, db *gorm.DB) int
	Delete(id int, db *gorm.DB) int
	ListAll(db *gorm.DB) []entity.Room
	FindOne(id int, db *gorm.DB) entity.Room
}
