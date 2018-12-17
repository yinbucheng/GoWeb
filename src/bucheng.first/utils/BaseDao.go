package utils

import . "github.com/jinzhu/gorm"

type BaseDao interface {
	Save(bean interface{}, db *DB)
	Delete(id int, bean interface{}, db *DB)
	Update(id int, bean interface{}, db *DB)
	FindOne(id int, db *DB, bean interface{})
}
