package model

import (
	"time"
	"github.com/globalsign/mgo/bson"
)
//USERTeacher
type UserInfo struct {
	User []User `json:"users"`
}

type User struct {
	ID         bson.ObjectId	  `json:"user_id" bson:"_id,omitempty"`
	FirstName  string         		`json:"tFirstName" bson:"tFirstName"`
	LastName   string          		`json:"tLastName" bson:"tLastName"`
	UserName   string             `json:"serName" bson:"UserName"`
	TId        string             `json:"TId" bson:"TId"`
	TEmail     string             `json:"TEmail" bson:"TEmail"`
	TWorkPlace string             `json:"TWorkPlace" bson:"TWorkPlace"`
	TPassword  string             `json:"TPassword" bson:"TPassword"`
	TPicture   string             `json:"TPicture" bson:"TPicture"`

	
}


//Subject
type SubjectInfo struct {
	Subject []Subject `json:"subjects"`
}

type Subject struct {
	SubjectID       bson.ObjectId 		`json:"subject_id" bson:"_id,omitempty"`
	SID				string             	`json:"TSID" bson:"TSID"`
	SName      		string             	`json:"TSName" bson:"TSName"`
	SDescription 	string				`json:"TSDescription" bson:"TSDescription"`
	STeacher	  	string          	`json:"TSTeacher" bson:"TSTeacher"`
	Spassword		string				`json:"TSpassword" bson:"TSpassword"`
	TstudentInfo	[]TstudentInfo     `json:"TstudentInfo" bson:"TstudentInfo"`

	
}
type TstudentInfo struct {
	StudentID		string             	`json:"StudentID" bson:"StudentID"`
	SName      		string             	`json:"SName" bson:"SName"`
}



//Attendance
type AttendanceInfo struct {
	Attendance []Attendance `json:"attendances"`
}

type Attendance struct {
	AttendanceID       bson.ObjectId 	`json:"attendance_id" bson:"_id,omitempty"`
	AID				string             	`json:"AID" bson:"AID"`
	AName      		string             	`json:"AName" bson:"AName"`
	ATime			time.Time			`json:"ATime" bson:"ATime"`
	Astudent		string				`json:"Astudent" bson:"Astudent"`
	
}
//Student
type StudentInfo struct {
	Student []Student `json:"students"`
}

type Student struct {
	ObjectID       bson.ObjectId 	`json:"student_id" bson:"_id,omitempty"`
	StudentID		string          `json:"SID" bson:"SID"`
	SName      		string          `json:"SName" bson:"SName"`
	Susername		string          `json:"Susername" bson:"Susername"`
	Spassword		string          `json:"Spassword" bson:"Spassword"`
	Sphone			string          `json:"Sphone" bson:"Sphone"`
	Semail			string  		`json:"Semail" bson:"Semail"`
	
}


//QRCODE
// type Qrcode struct {
// 	Time string `json:"Time" bson:"Time"`
// 	User string `json:"User" bson:"User"`
// 	Pass string  `json:"Pass" bson:"Pass"`
// }
