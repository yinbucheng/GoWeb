package service

import "bucheng.first/entity"

type IUserService interface {
	Save(user *entity.User) int
	Update(user *entity.User) int
	Delete(user *entity.User) int
	FindOne(id int) entity.User
	ListAll() []entity.User
	SaveUserAndRoom() int
}
