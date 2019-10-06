package repository

import (
	"Go-backend-mongoDB/model"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
)
//AttendanceRepository is ...
type AttendanceRepository interface {
	GetAllAttendance() ([]model.Attendance, error)
	CreateAttendance(attendance model.Attendance) error
	GetAttendance(password string)  ([]model.Attendance, error)
	GetONEAllAttendance(password string) ([]model.Attendance, error)
	Checkcode(check string) ([]model.Attendance, error)
	CheckingName(check model.CheckStudent)  error
	GetSelfie(pass string,date string,timeauthens int)   ([]model.Attendance, error)
	GetInfoStudent(pass string,date string,timeauthens int)   ([]model.Attendance, error)
	CheckingNamewithSocket(check model.CheckStudent)  ([]model.Attendance, error)
	GetClient(clientid string)   ([]model.Attendance, error)
}
//AttendanceRepositoryMongo is ...
type AttendanceRepositoryMongo struct {
	ConnectionDB *mgo.Session
}
//GetAllAttendance is ...
func (AttendanceMongo AttendanceRepositoryMongo) GetAllAttendance() ([]model.Attendance, error) {
	var attendance []model.Attendance
	err := AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(nil).All(&attendance)
	return attendance, err
}
//CreateAttendance is ...
func (AttendanceMongo AttendanceRepositoryMongo) CreateAttendance(attendance model.Attendance) error {
	return AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Insert(attendance)
}
//GetAttendance is ...
func (AttendanceMongo AttendanceRepositoryMongo) GetAttendance(password string)  ([]model.Attendance, error) {
	var attendances []model.Attendance
	name := bson.M{"ASpassword" : password ,}
	err:= AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(name).All(&attendances)
	return attendances, err
}
//GetONEAllAttendance is ...
func (AttendanceMongo AttendanceRepositoryMongo) GetONEAllAttendance(password string) ([]model.Attendance, error) {
	var attendance []model.Attendance
	name := bson.M{"ASpassword" : password ,}
	err := AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(name).All(&attendance)
	return attendance, err
}
//Checkcode is ...
func (AttendanceMongo AttendanceRepositoryMongo) Checkcode(check string) ([]model.Attendance, error) {
	var attendance []model.Attendance
	name := bson.M{"AQRcode" : check ,}
	err := AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(name).All(&attendance)
	return attendance, err
}
//CheckingName is ...
func (AttendanceMongo AttendanceRepositoryMongo) CheckingName(check model.CheckStudent)  error {
	name := bson.M{"AQRcode" : check.AQRcode,}
	newData := bson.M{"$push": bson.M{"Astudent":bson.M{"StudentID": check.StudentID,"ImageSelfie" : check.ImageSelfie}}}
	return AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Update(name, newData)
}
//GetSelfie is ...
func (AttendanceMongo AttendanceRepositoryMongo) GetSelfie(pass string,date string,timeauthens int)   ([]model.Attendance, error) {
	var attendance []model.Attendance
	name := bson.M{"ASpassword" : pass ,"ADate" : date ,"ATimeAuthen" : timeauthens ,}
	err := AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(name).All(&attendance)
	return attendance, err
}
//GetInfoStudent is ...
func (AttendanceMongo AttendanceRepositoryMongo) GetInfoStudent(pass string,date string,timeauthens int)   ([]model.Attendance, error) {
	var attendance []model.Attendance
	name := bson.M{"ASpassword" : pass ,"ADate" : date ,"ATimeAuthen" : timeauthens ,}
	err := AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(name).All(&attendance)
	return attendance, err
}
//CheckingNamewithSocket is ...
func (AttendanceMongo AttendanceRepositoryMongo) CheckingNamewithSocket(check model.CheckStudent)  ([]model.Attendance, error) {
	var attendance []model.Attendance
	name := bson.M{"AQRcode" : check.AQRcode,}
	newData := bson.M{"$push": bson.M{"Astudent":bson.M{"StudentID": check.StudentID,"ImageSelfie" : check.ImageSelfie,"SfirstName" : check.SfirstName,"SlastName" : check.SlastName,}}}
	err := AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Update(name, newData)
	err = AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(name).All(&attendance)
	return attendance, err
}
//GetClient is ...
func (AttendanceMongo AttendanceRepositoryMongo) GetClient(clientid string)   ([]model.Attendance, error) {
	var attendance []model.Attendance
	name := bson.M{"AClientid" : clientid ,}
	err := AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(name).All(&attendance)
	return attendance, err
}