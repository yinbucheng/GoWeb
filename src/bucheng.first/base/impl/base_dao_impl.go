package impl

import (
	"github.com/jinzhu/gorm"
)

type BaseDaoImpl struct {
}

func (baseDao BaseDaoImpl) Save(bean interface{}, db *gorm.DB) error {
	err := db.Create(bean).Error
	return err
}
func (baseDao BaseDaoImpl) Delete(id int, bean interface{}, db *gorm.DB) error {
	err := db.Where("id=?", id).Delete(bean).Error
	return err
}
func (baseDao BaseDaoImpl) Update(id int, bean interface{}, db *gorm.DB) error {
	err := db.Model(bean).Where("id=?", id).Update(bean).Error
	return err
}
func (baseDao BaseDaoImpl) FindOne(id int, bean interface{}, db *gorm.DB) error {
	err := db.Where("id=?", id).First(bean).Error
	return err
}
func (baseDao BaseDaoImpl) ListAll(bean interface{}, db *gorm.DB) error {
	err := db.Find(bean).Error
	return err
}
