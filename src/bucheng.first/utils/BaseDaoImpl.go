package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type BaseDaoImpl struct {
}

func (baseDao BaseDaoImpl) Save(bean interface{}, db *gorm.DB) {
	err := db.Create(bean).Error
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
}

func (baseDao BaseDaoImpl) Delete(id int, bean interface{}, db *gorm.DB) {
	err := db.Where("id=?", id).Delete(bean).Error
	if err != nil {
		panic(err)
	}
}

func (baseDao BaseDaoImpl) Update(id int, bean interface{}, db *gorm.DB) {
	t := reflect.TypeOf(bean)
	v := reflect.ValueOf(bean)
	data := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Name == "id" {
			continue
		}
		val := v.Field(i).Interface()
		data[f.Name] = val
	}
	err := db.Model(bean).Update(data).Error
	if err != nil {
		panic(err)
	}
}

func (baseDao BaseDaoImpl) FindOne(id int, db *gorm.DB, bean interface{}) {
	db.Where("id=?", id).First(bean)
}
