package repository

import (
	"Go-backend-mongoDB/model"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
	// "fmt"
	
)
type SubjectRepository interface {
	GetAllSubject() ([]model.Subject, error)
	AddSubject(subject model.Subject) error
	EditDescription(ID string, subject model.Subject) error
	DeleteSubjectByID(subjectID string) error
	Editaddstudent(ID string, subject model.Subject) error
	GetSubject(subjectsInfo model.SubjectInfo, ID string)  ([]model.Subject, error)
	JoinClass(ID string ,Sid string)   error 
}

type SubjectRepositoryMongo struct {
	ConnectionDB *mgo.Session
}

func (SubjectMongo SubjectRepositoryMongo) GetAllSubject() ([]model.Subject, error) {
	var subjects []model.Subject
	err := SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Find(nil).All(&subjects)
	return subjects, err
}

func (SubjectMongo SubjectRepositoryMongo) AddSubject(subject model.Subject) error {
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Insert(subject)
}

func (SubjectMongo SubjectRepositoryMongo) EditDescription(ID string, subject model.Subject) error {
	objectID := bson.ObjectIdHex(ID)
	newDescription := bson.M{"$set": bson.M{"TSDescription": subject.TSDescription, }}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).UpdateId(objectID, newDescription)
}
func (SubjectMongo SubjectRepositoryMongo) DeleteSubjectByID(subjectID string) error {
	objectID := bson.ObjectIdHex(subjectID)
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).RemoveId(objectID)
}
func (SubjectMongo SubjectRepositoryMongo) Editaddstudent(ID string, subject model.Subject) error {
	objectID := bson.ObjectIdHex(ID)
	newDescription := bson.M{"$set": bson.M{"TSDescription": subject.TSDescription, }}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).UpdateId(objectID, newDescription)

}

func (SubjectMongo SubjectRepositoryMongo) GetSubject(subjectsInfo model.SubjectInfo, ID string)  ([]model.Subject, error){
	var subjects []model.Subject
	objectID := bson.ObjectIdHex(ID)
	err:= SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).FindId(objectID).All(&subjects)
	// err := SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Find(nil).All(&subjects)
	return subjects, err
}



func (SubjectMongo SubjectRepositoryMongo) JoinClass(ID string ,Sid string)   error {

	objectID := bson.ObjectIdHex(ID)
	newData := bson.M{"$push": bson.M{"TstudentInfo":bson.M{"StudentID": Sid}}}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).UpdateId(objectID, newData)
}
