package repository

import (
	"Go-backend-mongoDB/model"
	//"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
)
//Student
type StudentRepository interface {
	GetAllStudent() ([]model.Student, error)
	CreateStudent(student model.Student) error
}

type StudentRepositoryMongo struct {
	ConnectionDB *mgo.Session
}
func (StudentMongo StudentRepositoryMongo) GetAllStudent() ([]model.Student, error) {
	var student []model.Student
	err := StudentMongo.ConnectionDB.DB(DBName).C(collectionStudent).Find(nil).All(&student)
	return student, err
}
func (StudentMongo StudentRepositoryMongo) CreateStudent(student model.Student) error {
	return StudentMongo.ConnectionDB.DB(DBName).C(collectionStudent).Insert(student)
}