package service

import "bucheng.first/entity"

type IClassRoomService interface {
	Save(user *entity.ClassRoom) int
	Update(user *entity.ClassRoom) int
	Delete(user *entity.ClassRoom) int
	FindOne(id int) entity.ClassRoom
}
