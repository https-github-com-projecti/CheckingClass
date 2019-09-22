package repository

import (
	"Go-backend-mongoDB/model"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
)
type SubjectRepository interface {
	GetAllSubject() ([]model.Subject, error)
	AddSubject(subject model.Subject) error
	EditDescription(ID string, subject model.Subject) error
	DeleteSubjectByID(subjectID string) error
	Editaddstudent(ID string, subject model.Subject) error
}

type SubjectRepositoryMongo struct {
	ConnectionDB *mgo.Session
}

func (SubjectMongo SubjectRepositoryMongo) GetAllSubject() ([]model.Subject, error) {
	var suibjects []model.Subject
	err := SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Find(nil).All(&suibjects)
	return suibjects, err
}

func (SubjectMongo SubjectRepositoryMongo) AddSubject(subject model.Subject) error {
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).Insert(subject)
}

func (SubjectMongo SubjectRepositoryMongo) EditDescription(ID string, subject model.Subject) error {
	objectID := bson.ObjectIdHex(ID)
	newDescription := bson.M{"$set": bson.M{"TSDescription": subject.SDescription, }}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).UpdateId(objectID, newDescription)
}
func (SubjectMongo SubjectRepositoryMongo) DeleteSubjectByID(subjectID string) error {
	objectID := bson.ObjectIdHex(subjectID)
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).RemoveId(objectID)
}
func (SubjectMongo SubjectRepositoryMongo) Editaddstudent(ID string, subject model.Subject) error {
	objectID := bson.ObjectIdHex(ID)
	newStudent := bson.M{"$set": bson.M{"TSstudents": subject.TstudentInfo, }}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).UpdateId(objectID, newStudent)
}