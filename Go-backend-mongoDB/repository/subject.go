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
	JoinClass(tspassword int, join model.JoinSubject)   error
	GetOneSubject(id string)  ([]model.Subject, error)
	GETStudentinClass(id int)  ([]model.Subject, error)
	UserJoin(user model.JoinUser) error
	GETTime(id int)  ([]model.Subject, error)
	EditTimeLimit(id int, subject model.Subject ,time model.TimeLimitandtimeOut) error
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
func (SubjectMongo SubjectRepositoryMongo) 	DeleteSubject(id string) error{
	objectID := bson.ObjectIdHex(id)
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).RemoveId(objectID)
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
func (SubjectMongo SubjectRepositoryMongo) JoinClass(tspassword int, join model.JoinSubject)   error {
	name := bson.M{"TSpassword" : tspassword,}
	newData := bson.M{"$push": bson.M{"TstudentInfo":bson.M{"StudentID": join.StudentID,"SfirstName" : join.SfirstName,"SlastName" : join.SlastName,}}}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Update(name, newData)
}
//GETStudentinClass is ...
func (SubjectMongo SubjectRepositoryMongo) GETStudentinClass(id int)  ([]model.Subject, error){
	var subjects []model.Subject
	name := bson.M{"TSpassword" : id ,}
	err:= SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Find(name).All(&subjects)
	return subjects, err
}
//UserJoin is ...
func (SubjectMongo SubjectRepositoryMongo) UserJoin(user model.JoinUser) error{
	name := bson.M{"TSpassword" : user.Pass ,}
	newPassword := bson.M{"$push": bson.M{"TSTeacher": user.User, }}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Update(name, newPassword)
}
//GETTime is ...
func (SubjectMongo SubjectRepositoryMongo) GETTime(id int)  ([]model.Subject, error){
	var subjects []model.Subject
	name := bson.M{"TSpassword" : id ,}
	err:= SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Find(name).All(&subjects)
	return subjects, err
}
//EditTimeLimit is ...
func (SubjectMongo SubjectRepositoryMongo) EditTimeLimit(id int, subject model.Subject ,time model.TimeLimitandtimeOut) error {
	name := bson.M{"TSpassword" : id ,}
	newDescription := bson.M{"$set": bson.M{"TSlimit": time.TSlimit, "TStimeout": time.TStimeout}}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Update(name, newDescription)
}
