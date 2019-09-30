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
	UserName   string             `json:"UserName" bson:"UserName"`
	TId        string             `json:"TId" bson:"TId"`
	TEmail     string             `json:"TEmail" bson:"TEmail"`
	TWorkPlace string             `json:"TWorkPlace" bson:"TWorkPlace"`
	TPassword  string             `json:"TPassword" bson:"TPassword"`
	TPicture   string             `json:"TPicture" bson:"TPicture"`

	
}
type Login struct {
	UserName   string             `json:"username" bson:"username"`
	TPassword  string             `json:"password" bson:"password"`
}


//Subject
type SubjectInfo struct {
	Subject []Subject `json:"subjects"`
}




type Subject struct {
	SubjectID       bson.ObjectId 		`json:"subject_id" bson:"_id,omitempty"`
	TSID			string             	`json:"TSID" bson:"TSID"`
	TSName      	string             	`json:"TSName" bson:"TSName"`
	TSDescription 	string				`json:"TSDescription" bson:"TSDescription"`
	TSTeacher	  	string          	`json:"TSTeacher" bson:"TSTeacher"`
	TSpassword		int				`json:"TSpassword" bson:"TSpassword"`
	TstudentInfo	[]TstudentInfos     `json:"TstudentInfo" bson:"TstudentInfo"`

}
type TstudentInfos struct {
	StudentID		string             	`json:"StudentID" bson:"StudentID"`
}
type JoinSubject struct {
	Jpassword		string				`json:"Jpassword" bson:"Jpassword"`
	JSID			string         		 `json:"JSID" bson:"JSID"`
	ID				int        		 	`json:"id" bson:"id"`
}




//Attendance
type AttendanceInfo struct {
	Attendance []Attendance `json:"attendances"`
}

type Attendance struct {
	AttendanceID       bson.ObjectId 	`json:"attendance_id" bson:"_id,omitempty"`
	AID				string             	`json:"AID" bson:"AID"`
	AName      		string             	`json:"AName" bson:"AName"`
	ASpassword		string				`json:"ASpassword" bson:"ASpassword"`
	ATime			time.Time			`json:"ATime" bson:"ATime"`
	Astudent		[]TstudentInfos		`json:"Astudent" bson:"Astudent"`
	
}


//Student
type StudentInfo struct {
	Student []Student `json:"students"`
}

type Student struct {
	ObjectID       bson.ObjectId 	`json:"student_id" bson:"_id,omitempty"`
	StudentID		string          `json:"SID" bson:"SID"`
	SfirstName      string          `json:"SfirstName" bson:"SfirstName"`
	SlastName      string          `json:"SlastName" bson:"SlastName"`
	Susername		string          `json:"Susername" bson:"Susername"`
	Spassword		string          `json:"Spassword" bson:"Spassword"`
	Sphone			string          `json:"Sphone" bson:"Sphone"`
	Semail			string  		`json:"Semail" bson:"Semail"`
	
}


// type Qrcode struct {
// 	Time 	string `json:"time"`
// 	User 	string `json:"user"`
// 	Pass 	int    `json:"passOfCouse"`
// }

type CreateQr struct {
	ObjectID       bson.ObjectId 	`json:"qrcode_id" bson:"_id,omitempty"`
	Qrcode			string			`json:"qrcode" bson:"qrcode"`
	Time 			string 			`json:"time" bson:"time"`
	Pass 			int    			`json:"passOfCouse" bson:"passOfCouse"`
	TimeAuthen 		int 			`json:"time_authen" bson:"time_authen"`
}
