package repository

import (
	"Go-backend-mongoDB/model"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
)
//StudentRepository is ...
type StudentRepository interface {
	GetAllStudent() ([]model.Student, error)
	CreateStudent(student model.Student) error
	LoginStudent(username string) ([]model.Student, error) 
	GetStudent(username string)  ([]model.Student, error)
	Getbyid(id string)  ([]model.Student, error)
	GetsubjectbySid(sid string)  ([]model.Subject, error)
	GetStudentAttendance(password string)  ([]model.Attendance, error)
	GETStudenClasses(id int)  ([]model.Subject, error)
	DeleteStudent(id string) error
}
//StudentRepositoryMongo is ...
type StudentRepositoryMongo struct {
	ConnectionDB *mgo.Session
}
//GetAllStudent is ...
func (StudentMongo StudentRepositoryMongo) GetAllStudent() ([]model.Student, error) {
	var student []model.Student
	err := StudentMongo.ConnectionDB.DB(DBName).C(collectionStudent).Find(nil).All(&student)
	return student, err
}
//CreateStudent is ...
func (StudentMongo StudentRepositoryMongo) CreateStudent(student model.Student) error {
	return StudentMongo.ConnectionDB.DB(DBName).C(collectionStudent).Insert(student)
}
//LoginStudent is ...
func (StudentMongo StudentRepositoryMongo) LoginStudent(username string) ([]model.Student, error) {
	var users []model.Student
	checker := bson.M{"Susername" : username ,} 
	err := StudentMongo.ConnectionDB.DB(DBName).C(collectionStudent).Find(checker).All(&users)
	return users, err
	
}
//GetStudent is ...
func (StudentMongo StudentRepositoryMongo) GetStudent(id string)  ([]model.Student, error){
	var users []model.Student
	name := bson.M{"SID" : id ,}
	err:= StudentMongo.ConnectionDB.DB(DBName).C(collectionStudent).Find(name).All(&users)
	return users, err
}
//Getbyid is ...
func (StudentMongo StudentRepositoryMongo) Getbyid(id string)  ([]model.Student, error){
	var users []model.Student
	objectID := bson.ObjectIdHex(id)
	err:= StudentMongo.ConnectionDB.DB(DBName).C(collectionStudent).FindId(objectID).All(&users)
	return users, err
}
//GetsubjectbySid is ...
func (StudentMongo StudentRepositoryMongo) GetsubjectbySid(sid string)  ([]model.Subject, error){
	var users []model.Subject
	name := bson.M{"TstudentInfo.StudentID" : sid,}
	err:= StudentMongo.ConnectionDB.DB(DBName).C(collectionSubject).Find(name).All(&users)
	return users, err
}
//GetStudentAttendance is ...
func (StudentMongo StudentRepositoryMongo) GetStudentAttendance(password string)  ([]model.Attendance, error) {
	var attendances []model.Attendance
	name := bson.M{"ASpassword" : password ,}
	err:= StudentMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(name).All(&attendances)
	return attendances, err
}
//GETStudenClasses is ...
func (StudentMongo StudentRepositoryMongo) GETStudenClasses(id int)  ([]model.Subject, error){
	var subjects []model.Subject
	name := bson.M{"TSpassword" : id ,}
	err:= StudentMongo.ConnectionDB.DB(DBName).C(collectionSubject).Find(name).All(&subjects)
	return subjects, err
}
//DeleteStudent is ...
func (StudentMongo StudentRepositoryMongo) DeleteStudent(id string) error{
	objectID := bson.ObjectIdHex(id)
	return StudentMongo.ConnectionDB.DB(DBName).C(collectionStudent).RemoveId(objectID)
}
