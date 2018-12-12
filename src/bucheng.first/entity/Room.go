package entity

import "github.com/jinzhu/gorm"

type Room struct {
	ID       int    `gorm:"primary_key" json:"id"`
	RoomName string `gorm:"type:varchar(20);not null" json:"roomName"`
	Address  string `gorm:"type:varchar(20);not null" json:"address"`
}

func CreateTable(db *gorm.DB, value interface{}) {
	if !db.HasTable(value) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(value).Error; err != nil {
			panic(err)
		}
	}
}
