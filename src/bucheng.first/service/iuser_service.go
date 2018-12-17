package service

import (
	"bucheng.first/base"
	"os/user"
)

type IUserService interface {
	base.BaseService
	FindAll() []user.User
	AffairTest() (err error)
}
