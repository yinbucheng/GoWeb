package dao

import (
	"bucheng.first/base"
	"os/user"
)

type UserDao interface {
	base.BaseDao
	FindAll() []user.User
}
