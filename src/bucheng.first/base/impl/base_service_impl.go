package impl

import (
	"bucheng.first/base"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type BaseServiceImpl struct {
	Dao base.BaseDao
	Db  *gorm.DB
}

func (service BaseServiceImpl) InitAttriute(dao base.BaseDao, db *gorm.DB) {
	service.Dao = dao
	service.Db = db
}

func (service BaseServiceImpl) Save(bean interface{}) int {
	err := service.Dao.Save(bean, service.Db)
	if err != nil {
		logrus.Error(err)
		return -1
	}
	return 1
}
func (service BaseServiceImpl) Delete(id int, bean interface{}) int {
	err := service.Dao.Delete(id, bean, service.Db)
	if err != nil {
		logrus.Error(err)
		return -1
	}
	return 1
}
func (service BaseServiceImpl) Update(id int, bean interface{}) int {
	err := service.Dao.Update(id, bean, service.Db)
	if err != nil {
		logrus.Error(err)
		return -1
	}
	return 1
}
func (service BaseServiceImpl) FindOne(id int, bean interface{}) {
	err := service.Dao.FindOne(id, bean, service.Db)
	if err != nil {
		logrus.Error(err)
	}
}
func (service BaseServiceImpl) ListAll(bean interface{}) {
	err := service.Dao.ListAll(bean, service.Db)
	if err != nil {
		logrus.Error(err)
	}
}
