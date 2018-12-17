package impl

import (
	"bucheng.first/dao"
	"bucheng.first/dao/impl"
	"bucheng.first/entity"
	"bucheng.first/utils"
	"github.com/jinzhu/gorm"
)

type ClassRoomServiceImpl struct {
	dao dao.ClassRoomDao
	db  *gorm.DB
}

func NewClassRoomService() *ClassRoomServiceImpl {
	classRoomService := new(ClassRoomServiceImpl)
	classRoomService.dao = new(impl.ClassRoomDaoImpl)
	classRoomService.db = utils.InitDb("root", "introcks1234", "go_test")
	return classRoomService
}

func (service ClassRoomServiceImpl) Save(classRoom *entity.ClassRoom) int {
	service.dao.Save(classRoom, service.db)
	return 1
}
func (service ClassRoomServiceImpl) Update(classRoom *entity.ClassRoom) int {
	service.dao.Update(classRoom.ID, classRoom, service.db)
	return 1
}
func (service ClassRoomServiceImpl) Delete(classRoom *entity.ClassRoom) int {
	service.dao.Delete(classRoom.ID, classRoom, service.db)
	return 1
}
func (service ClassRoomServiceImpl) FindOne(id int) entity.ClassRoom {
	var classRoom entity.ClassRoom
	service.dao.FindOne(id, service.db, classRoom)
	return classRoom
}
