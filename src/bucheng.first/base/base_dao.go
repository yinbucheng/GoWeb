package base

import "github.com/jinzhu/gorm"

type BaseDao interface {
	//保存记录
	Save(bean interface{}, db *gorm.DB) error
	//通过id删除记录
	Delete(id int, bean interface{}, db *gorm.DB) error
	//通过id更新记录
	Update(id int, bean interface{}, db *gorm.DB) error
	//通过id查询记录
	FindOne(id int, bean interface{}, db *gorm.DB) error
	//查询所有记录
	ListAll(bean interface{}, db *gorm.DB) error
}
