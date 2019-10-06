package repository

import (
	"Go-backend-mongoDB/model"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
	// "fmt"
	
)
//SubjectRepository is ...
type SubjectRepository interface {
	GetAllSubject() ([]model.Subject, error)
	AddSubject(subject model.Subject) error
	EditDescription(tsid string, subject model.Subject) error
	DeleteSubject(tsid string) error
	GetSubject(tsid string)  ([]model.Subject, error)
	JoinClass(tspassword int, sid string)    error 
	GetOneSubject(id string)  ([]model.Subject, error)
}
//SubjectRepositoryMongo is ...
type SubjectRepositoryMongo struct {
	ConnectionDB *mgo.Session
}
//GetAllSubject is ...
func (SubjectMongo SubjectRepositoryMongo) GetAllSubject() ([]model.Subject, error) {
	var subjects []model.Subject
	err := SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Find(nil).All(&subjects)
	return subjects, err
}
//AddSubject is ...
func (SubjectMongo SubjectRepositoryMongo) AddSubject(subject model.Subject) error {
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Insert(subject)
}
//EditDescription is ...
func (SubjectMongo SubjectRepositoryMongo) EditDescription(tsid string, subject model.Subject) error {
	name := bson.M{"TSID" : tsid ,}
	newDescription := bson.M{"$set": bson.M{"TSDescription": subject.TSDescription, }}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Update(name, newDescription)
}
//DeleteSubject is ...
func (SubjectMongo SubjectRepositoryMongo) 	DeleteSubject(tsid string) error{
	name := bson.M{"TSID" : tsid ,}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Remove(name)
}
//GetSubject is ...
func (SubjectMongo SubjectRepositoryMongo) GetSubject(id string)  ([]model.Subject, error){
	var subjects []model.Subject
	name := bson.M{"TSTeacher" : id ,}
	err:= SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Find(name).All(&subjects)
	return subjects, err
}
//GetOneSubject is ...
func (SubjectMongo SubjectRepositoryMongo) GetOneSubject(id string)  ([]model.Subject, error){
	var subjects []model.Subject
	objectID := bson.ObjectIdHex(id)
	err:= SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).FindId(objectID).All(&subjects)
	return subjects, err
}
//JoinClass is ...
func (SubjectMongo SubjectRepositoryMongo) JoinClass(tspassword int, sid string)   error {
	name := bson.M{"TSpassword" : tspassword,}
	newData := bson.M{"$push": bson.M{"TstudentInfo":bson.M{"StudentID": sid}}}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Update(name, newData)
}
