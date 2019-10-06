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
	name := bson.M{"TstudentInfo" :bson.M{"StudentID" : sid,}}
	err:= StudentMongo.ConnectionDB.DB(DBName).C(collectionSubject).Find(name).All(&users)
	return users, err
}