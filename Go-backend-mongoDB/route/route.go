package route

import (
	"Go-backend-mongoDB/api"
	"Go-backend-mongoDB/repository"
	"Go-backend-mongoDB/websocket"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)
//NewRouteUser is ...
func NewRouteUser(route *gin.Engine, connectionDB *mgo.Session) {
	//User
	userRepository := repository.UserRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	userAPI := api.UserAPI{
		UserRepository: &userRepository,
	}
	route.POST("user/login", userAPI.LoginHandler)
	route.GET("user/list", userAPI.UserListHandler)
	route.POST("user/Add", userAPI.AddUserHandeler)
	route.PUT("user/Edit", userAPI.EditPasswordHandler)
	route.DELETE("user/Delete/:id", userAPI.DeleteUserByIDHandler)
	route.GET("user/GETONE", userAPI.GETONEUserHandeler)
	route.GET("user/GETONE/:username", userAPI.GETONEnameHandeler)
	route.GET("user/getMyPic/:id", userAPI.GETPictureUserHandeler)

	//SUBject
	subjectRepository := repository.SubjectRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	subjectAPI := api.SubjectAPI{
		SubjectRepository: &subjectRepository,
	}
	route.GET("Subject/list", subjectAPI.SubjectListHandler)
	route.POST("Subject/Add", subjectAPI.AddSubjectHandeler)
	route.PUT("Subject/Edit", subjectAPI.EditDescriptionHandler)
	route.PUT("Subject/TimeEdit/:pass", subjectAPI.EditTimelimitHandler)
	route.DELETE("Subject/Delete/:id", subjectAPI.DeleteSubjectHandler)
	route.GET("Subject/GetOneSubject/:id", subjectAPI.GETONESubjectHandeler)
	route.GET("Subject/GetMySubject/:id", subjectAPI.GETMySubjectHandeler)
	route.GET("Subject/GetStudent/:pass", subjectAPI.GETStudentinClassHandeler)
	route.GET("Subject/time/:pass", subjectAPI.GETTimeHandeler)
	route.POST("Subject/join", subjectAPI.JoinClassHandeler)
	route.POST("Subject/Userjoin", subjectAPI.UserJoinHandler)

	//Attendance
	attendanceRepository := repository.AttendanceRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	attendanceAPI := api.AttendanceAPI{
		AttendanceRepository: &attendanceRepository,
	}
	route.GET("Attendance/alllist", attendanceAPI.AllAttendanceListHandler)
	route.POST("Attendance/new", attendanceAPI.CreateAttendanceHandeler)
	route.POST("Attendance/Create", attendanceAPI.CreateQrcodeAndAttendanceHandeler)
	route.GET("Attendance/GETONE", attendanceAPI.GETONEAttendanceHandeler)
	route.POST("Attendance/CheckQR", attendanceAPI.CheckQRcodeHandeler)
	route.GET("Attendance/getAttendance/:pass", attendanceAPI.GetAttendanceCheck)
	route.GET("Attendance/getQRcode/:pass", attendanceAPI.GETQRcoderHandeler)
	route.POST("Attendance/Checkname", attendanceAPI.CheckNameClassHandeler)
	route.POST("Attendance/time", attendanceAPI.TimelimitHandeler)
	route.GET("Attendance/selfie/:pass/:date/:timeauthens", attendanceAPI.GETSelfieHandeler)
	route.GET("Attendance/clientid/:clientid", attendanceAPI.GETClientHandeler)
	route.POST("Attendance/Checknamewithsocket", attendanceAPI.CheckwithSocketHandeler)
	route.POST("Attendance/timescore", attendanceAPI.TimeScoreHandeler)
	route.GET("Attendance/info/:pass/:date/:timeauthens", attendanceAPI.GETinfoStudentHandeler)
	route.DELETE("Attendance/Delete/:id", attendanceAPI.DeleteAttendanceHandler)
	// route.GET("Attendance/Checking/:cilentId", attendanceAPI.ShowCheckingHandeler)


	//Student
	studentRepository := repository.StudentRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	studentAPI := api.StudentAPI{
		StudentRepository: &studentRepository,
	}
	route.GET("Student/alllist", studentAPI.AllStudentListHandler)
	route.POST("Student/new", studentAPI.CreateStudentHandeler)
	route.GET("Student/subjectlist:ID", studentAPI.ShowsubjectlistStudentHandeler)
	route.POST("Student/login", studentAPI.StudentLoginHandler)
	route.GET("Student/info/:id", studentAPI.StudentinfoListHandler)
	route.GET("Student/subject/:sid", studentAPI.StudenSubjectListHandler)
	route.GET("Student/Attendance/:sid/:pass", studentAPI.GetAttendancebyPassandID)
	route.DELETE("Student/Delete/:id", studentAPI.DeleteStudentHandler)
	route.PUT("Student/Editpassword", studentAPI.EditstudentPasswordHandler)
	route.PUT("Student/Editemail", studentAPI.EditstudentEmailHandler)
	route.PUT("Student/Editphone", studentAPI.EditstudentPhoneHandler)
	
	
	//WS
	route.GET("Websocket/ws", websocket.Websockethandler)
	go websocket.Manager.Start()
	
	
}


