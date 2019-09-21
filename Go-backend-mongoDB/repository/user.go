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
	newDescription := bson.M{"$set": bson.M{"SDescription": subject.SDescription, }}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).UpdateId(objectID, newDescription)
}
func (SubjectMongo SubjectRepositoryMongo) DeleteSubjectByID(subjectID string) error {
	objectID := bson.ObjectIdHex(subjectID)
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).RemoveId(objectID)
}
func (SubjectMongo SubjectRepositoryMongo) Editaddstudent(ID string, subject model.Subject) error {
	objectID := bson.ObjectIdHex(ID)
	newStudent := bson.M{"$set": bson.M{"Sstudent": subject.Sstudent, }}
	return SubjectMongo.ConnectionDB.DB(DBName).C(collectionSubject).UpdateId(objectID, newStudent)
}

//Attendance
type AttendanceRepository interface {
	GetAllAttendance() ([]model.Attendance, error)
	CreateAttendance(attendance model.Attendance) error
	FindbyAName(Aname string)([]model.Attendance, error)
	
}

type AttendanceRepositoryMongo struct {
	ConnectionDB *mgo.Session
}
func (AttendanceMongo AttendanceRepositoryMongo) GetAllAttendance() ([]model.Attendance, error) {
	var attendance []model.Attendance
	err := AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(nil).All(&attendance)
	return attendance, err
}
func (AttendanceMongo AttendanceRepositoryMongo) CreateAttendance(attendance model.Attendance) error {
	return AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Insert(attendance)
}
func (AttendanceMongo AttendanceRepositoryMongo) FindbyAName(Aname string) ([]model.Attendance, error) {
	var attendance []model.Attendance
	objectID := bson.ObjectIdHex(Aname)
	err := AttendanceMongo.ConnectionDB.DB(DBName).C(collectionAttendance).Find(bson.M{"AName": objectID, }).All(&attendance)
	return attendance, err
}

//Student
type StudentRepository interface {
	GetAllStudent() ([]model.Student, error)
	
}

type StudentRepositoryMongo struct {
	ConnectionDB *mgo.Session
}
func (StudentMongo StudentRepositoryMongo) GetAllStudent() ([]model.Student, error) {
	var student []model.Student
	err := StudentMongo.ConnectionDB.DB(DBName).C(collectionStudent).Find(nil).All(&student)
	return student, err
}