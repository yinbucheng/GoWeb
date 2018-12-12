package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type DbUtils struct {
}

func InitDb(userName string, password string, dbName string) *gorm.DB {
	if db == nil {
		var err error
		var db2 *gorm.DB
		db2, err = gorm.Open("mysql", userName+":"+password+"@(127.0.0.1:3306)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic(err)
		}
		fmt.Println("---->NEW DB<----")
		db = db2
	}
	return db
}
