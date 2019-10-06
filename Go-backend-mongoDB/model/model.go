package model

import (

	"github.com/globalsign/mgo/bson"
)
//UserInfo is ...
type UserInfo struct {
	User []User `json:"users"`
}
//User is ...
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
//Login is ...
type Login struct {
	UserName   string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
}

//SubjectInfo is ...
type SubjectInfo struct {
	Subject []Subject `json:"subjects"`
}



//Subject is ...
type Subject struct {
	SubjectID       bson.ObjectId 		`json:"subject_id" bson:"_id,omitempty"`
	TSID			string             	`json:"TSID" bson:"TSID"`
	TSName      	string             	`json:"TSName" bson:"TSName"`
	TSDescription 	string				`json:"TSDescription" bson:"TSDescription"`
	TSTeacher	  	string          	`json:"TSTeacher" bson:"TSTeacher"`
	TSpassword		int				`json:"TSpassword" bson:"TSpassword"`
	TstudentInfo	[]TstudentInfos     `json:"TstudentInfo" bson:"TstudentInfo"`

}
//TstudentInfos is ...
type TstudentInfos struct {
	StudentID		string             	`json:"StudentID" bson:"StudentID"`
}
//JoinSubject is ...
type JoinSubject struct {
	Jpassword		string				`json:"Jpassword" bson:"Jpassword"`
	JSID			string         		 `json:"JSID" bson:"JSID"`
	ID				int        		 	`json:"id" bson:"id"`
}




//AttendanceInfo is ...
type AttendanceInfo struct {
	Attendance []Attendance `json:"attendances"`
}




//StudentInfo is ...
type StudentInfo struct {
	Student []Student `json:"students"`
}
//Student is ...
type Student struct {
	ObjectID       bson.ObjectId 	`json:"student_id" bson:"_id,omitempty"`
	StudentID		string          `json:"SID" bson:"SID"`
	SfirstName     	string          `json:"SfirstName" bson:"SfirstName"`
	SlastName      	string          `json:"SlastName" bson:"SlastName"`
	Susername		string          `json:"Susername" bson:"Susername"`
	Spassword		string          `json:"Spassword" bson:"Spassword"`
	Sphone			string          `json:"Sphone" bson:"Sphone"`
	Semail			string  		`json:"Semail" bson:"Semail"`
	
}

//Qrcode is ...
type Qrcode struct {
	Time 		string `json:"time"`
	User 		string `json:"user"`
	Pass 		int    `json:"passOfCouse"`
	Clientid  	string `json:"clientId"`
}
//CreateQr is ...
type CreateQr struct {
	ObjectID       bson.ObjectId 	`json:"qrcode_id" bson:"_id,omitempty"`
	Qrcode			string			`json:"qrcode" bson:"qrcode"`
	Time 			string 			`json:"time" bson:"time"`
	Pass 			int    			`json:"passOfCouse" bson:"passOfCouse"`
	TimeAuthen 		int 			`json:"time_authen" bson:"time_authen"`
}
//Attendance is ...
type Attendance struct {
	AttendanceID       bson.ObjectId 	`json:"attendance_id" bson:"_id,omitempty"`
	Date			string             	`json:"ADate" bson:"ADate"`
	ASpassword		string				`json:"ASpassword" bson:"ASpassword"`
	AQRcode			string				`json:"AQRcode" bson:"AQRcode"`
	PicQRcode		string				`json:"PicQRcode" bson:"PicQRcode"`
	ATimeAuthen		int 				`json:"ATimeAuthen" bson:"ATimeAuthen"`
	Clientid		string				`json:"AClientid" bson:"AClientid"`
	Astudent		[]CheckStudent		`json:"Astudent" bson:"Astudent"`
}
//CheckStudent is ...	
type CheckStudent struct{
	StudentID		string				`json:"StudentID" bson:"StudentID"`
	ImageSelfie		string				`json:"ImageSelfie" bson:"ImageSelfie"`
	AQRcode			string				`json:"AQRcode" bson:"AQRcode"`
	SfirstName     	string         		`json:"SfirstName" bson:"SfirstName"`
	SlastName      	string          	`json:"SlastName" bson:"SlastName"`
}
//CheckQRcode is ...	
type CheckQRcode struct{
	AQRcode			string				`json:"AQRcode" bson:"AQRcode"`
}
//SubjectOnly is ...	
type SubjectOnly struct{
	SubjectName      	string             	`json:"SubjectName" bson:"SubjectName"`
}
//NewSubjectName is ...	
type NewSubjectName struct{
	TSNames      	string             	`json:"TSName" bson:"TSName"`
}
//TimeLimit is ...	
type TimeLimit struct{
	Date			string             	`json:"ADate" bson:"ADate"`
}