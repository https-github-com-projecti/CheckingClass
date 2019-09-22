package repository

import (
	"Go-backend-mongoDB/model"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
)

type UserRepository interface {
	GetAllUser() ([]model.User, error)
	AddUser(user model.User) error
	EditPassword(ID string, user model.User) error
	DeleteUserByID(UserID string) error
}

type UserRepositoryMongo struct {
	ConnectionDB *mgo.Session
}

const (
	DBName     = "ProjectCheckName"
	collectionUser = "User"
	collectionSubject = "Subject"
	collectionAttendance = "Attendance"
	collectionStudent = "Student"
)

func (UserMongo UserRepositoryMongo) GetAllUser() ([]model.User, error) {
	var users []model.User
	err := UserMongo.ConnectionDB.DB(DBName).C(collectionUser).Find(nil).All(&users)
	return users, err
}

func (UserMongo UserRepositoryMongo) AddUser(user model.User) error {
	return UserMongo.ConnectionDB.DB(DBName).C(collectionUser).Insert(user)
}

func (UserMongo UserRepositoryMongo) EditPassword(ID string, user model.User) error {
	objectID := bson.ObjectIdHex(ID)
	newPassword := bson.M{"$set": bson.M{"TPassword": user.TPassword, }}
	return UserMongo.ConnectionDB.DB(DBName).C(collectionUser).UpdateId(objectID, newPassword)
}

func (UserMongo UserRepositoryMongo) DeleteUserByID(ID string) error {
	objectID := bson.ObjectIdHex(ID)
	return UserMongo.ConnectionDB.DB(DBName).C(collectionUser).RemoveId(objectID)
}






