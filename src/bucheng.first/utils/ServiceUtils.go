package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

/**
* 事务执行方法
 */
func ServiceMethod(method func(bean interface{}, db *gorm.DB), bean interface{}) {
	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			if tx != nil {
				tx.Rollback()
			}
		}
	}()
	method(bean, tx)
	tx.Commit()
}
