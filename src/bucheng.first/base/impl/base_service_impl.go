package impl

import (
	"bucheng.first/base"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type BaseServiceImpl struct {
	Dao base.BaseDao
	Db  *gorm.DB
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

//这里返回很有意思它通过参数加类型 然后在defer中赋值err就可以了
func (service BaseServiceImpl) ExecuteOnAffair(method func(params ...interface{}), param ...interface{}) (err error) {
	tx := service.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			logrus.Error(r)
			err = fmt.Errorf("%v", r)
			tx.Rollback()
		} else {
			tx.Commit()
			err = nil
		}
	}()
	//这里需要执行事务的方法最后的参数为*gorm.DB类型
	method(param, tx)
	return nil
}
