package impl

import (
	"bucheng.first/entity"
	"github.com/jinzhu/gorm"
)

type RoomDaoImpl struct {
}

func (roomDao RoomDaoImpl) Save(room *entity.Room, db *gorm.DB) int {
	err := db.Create(room).Error
	if err == nil {
		return -1
	}
	return 0
}

func (roomDao RoomDaoImpl) Update(room *entity.Room, db *gorm.DB) int {
	err := db.Model(room).Update("roomName", room.RoomName).Error
	if err != nil {
		return -1
	}
	return 0
}

func (roomDao RoomDaoImpl) Delete(id int, db *gorm.DB) int {
	err := db.Where("id=?", id).Delete(new(entity.Room)).Error
	if err != nil {
		return -1
	}
	return 0
}

func (roomDao RoomDaoImpl) ListAll(db *gorm.DB) []entity.Room {
	var rooms []entity.Room
	db.Find(&rooms)
	return rooms
}

func (roomDao RoomDaoImpl) FindOne(id int, db *gorm.DB) entity.Room {
	var room entity.Room
	db.Where("id=?", id).First(&room)
	return room
}
