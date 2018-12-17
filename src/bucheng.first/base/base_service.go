package base

type BaseService interface {
	Save(bean interface{}) int
	Delete(id int, bean interface{}) int
	Update(id int, bean interface{}) int
	FindOne(id int, bean interface{})
	ListAll(bean interface{})
	//开启事务执行此方法
	ExecuteOnAffair(method func(params ...interface{}), param ...interface{})
}
