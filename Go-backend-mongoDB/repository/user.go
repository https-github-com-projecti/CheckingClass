package repository

import (
	"Go-backend-mongoDB/model"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
)
//UserRepository is ...
type UserRepository interface {
	GetAllUser() ([]model.User, error)
	AddUser(user model.User) error
	DeleteUser(username string) error
	EditPassword(username string, password string) error
	GetUser(username string)  ([]model.User, error)
	CheckLogin(username string,password string) ([]model.User, error)
	Getbyid(id string)  ([]model.User, error)
	

}
//UserRepositoryMongo is ...
type UserRepositoryMongo struct {
	ConnectionDB *mgo.Session
}

const (
	//DBName is ...
	DBName     = "ProjectCheckName"
	collectionUser = "User"
	collectionSubject = "Subject"
	collectionAttendance = "Attendance"
	collectionStudent = "Student"
)
//GetAllUser is ...
func (UserMongo UserRepositoryMongo) GetAllUser() ([]model.User, error) {
	var users []model.User
	err := UserMongo.ConnectionDB.DB(DBName).C(collectionUser).Find(nil).All(&users)
	return users, err
}
//AddUser is ...
func (UserMongo UserRepositoryMongo) AddUser(user model.User) error {
	return UserMongo.ConnectionDB.DB(DBName).C(collectionUser).Insert(user)
}
//EditPassword is ...
func (UserMongo UserRepositoryMongo) EditPassword(username string, password string) error {
	name := bson.M{"UserName" : username ,}
	newPassword := bson.M{"$set": bson.M{"TPassword": password, }}
	return UserMongo.ConnectionDB.DB(DBName).C(collectionUser).Update(name, newPassword)
}
//DeleteUser is ...
func (UserMongo UserRepositoryMongo) DeleteUser(id string) error {
	objectID := bson.ObjectIdHex(id)
	return UserMongo.ConnectionDB.DB(DBName).C(collectionUser).RemoveId(objectID)
}
//GetUser is ...
func (UserMongo UserRepositoryMongo) GetUser(username string)  ([]model.User, error){
	var users []model.User
	name := bson.M{"UserName" : username ,}
	err:= UserMongo.ConnectionDB.DB(DBName).C(collectionUser).Find(name).All(&users)
	return users, err
}
//CheckLogin is ...
func (UserMongo UserRepositoryMongo) CheckLogin(username string,password string) ([]model.User, error) {
	var users []model.User
	checker := bson.M{"UserName" : username ,"TPassword" : password,} 
	err := UserMongo.ConnectionDB.DB(DBName).C(collectionUser).Find(checker).All(&users)
	return users, err
	
}
//Getbyid is ...
func (UserMongo UserRepositoryMongo) Getbyid(id string)  ([]model.User, error){
	var users []model.User
	objectID := bson.ObjectIdHex(id)
	err:= UserMongo.ConnectionDB.DB(DBName).C(collectionUser).FindId(objectID).All(&users)
	return users, err
}




