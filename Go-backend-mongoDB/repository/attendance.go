package repository

import (
	"Go-backend-mongoDB/model"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
)
//Attendance
type AttendanceRepository interface {
	GetAllAttendance() ([]model.Attendance, error)
	CreateAttendance(attendance model.Attendance) error
	FindbyAName(Aname string, attendance model.AttendanceInfo)error
	
}

type AttendanceRepositoryMongo struct {
	ConnectionDB *mgo.Session
}
func (AttendanceMongo AttendanceRepositoryMongo) GetAllAttendance() ([]model.Attendance, error) {
	var attendance []model.Attendance
	err := AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(nil).All(&attendance)
	return attendance, err
}
func (AttendanceMongo AttendanceRepositoryMongo) CreateAttendance(attendance model.Attendance) error {
	return AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Insert(attendance)
}
func (AttendanceMongo AttendanceRepositoryMongo) FindbyAName(Aname string, attendance model.AttendanceInfo)  error {
	objectID := bson.ObjectIdHex(Aname)
	return AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(bson.M{"AName": objectID, }).All(&attendance)
}
