package model

import (
	"time"
	"github.com/globalsign/mgo/bson"
)
//USER
type UserInfo struct {
	User []User `json:"users"`
}

type User struct {
	ID         bson.ObjectId	  `json:"user_id" bson:"_id,omitempty"`
	TName      string             `json:"TName" bson:"TName"`
	UserName   string             `json:"UserName" bson:"UserName"`
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
	SID				string             	`json:"SID" bson:"SID"`
	SName      		string             	`json:"SName" bson:"SName"`
	SDescription 	string				`json:"SDescription" bson:"Description"`
	STeacher	  	string          	`json:"STeacher" bson:"Teacher"`
	Sstudent		string      		`json:"Sstudent" bson:"Sstudent"`

	
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


//QRCODE
// type Qrcode struct {
// 	Time string `json:"Time" bson:"Time"`
// 	User string `json:"User" bson:"User"`
// 	Pass string  `json:"Pass" bson:"Pass"`
// }
