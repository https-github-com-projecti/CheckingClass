package route

import (
	"Go-backend-mongoDB/api"
	"Go-backend-mongoDB/repository"

	"github.com/globalsign/mgo"

	"github.com/gin-gonic/gin"
)

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
	route.PUT("TESTuser/Edit", userAPI.EditPasswordHandler)
	route.DELETE("user/Delete", userAPI.DeleteUserByIDHandler)
	route.GET("user/GETONE",userAPI.GETONEUserHandeler)
	route.GET("user/GETONE/:username",userAPI.GETONEnameHandeler)
	route.GET("user/getMyPic/:id",userAPI.GETPictureUserHandeler)
	
	

	
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
	route.DELETE("Subject/Delete", subjectAPI.DeleteSubjectHandler)
	route.GET("Subject/GetOneSubject/:id",subjectAPI.GETONESubjectHandeler)
	route.GET("Subject/GetMySubject/:id",subjectAPI.GETMySubjectHandeler)
	route.POST("Subject/join",subjectAPI.JoinClassHandeler)



	//Attendance	
	attendanceRepository := repository.AttendanceRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	attendanceAPI := api.AttendanceAPI{
		AttendanceRepository: &attendanceRepository,
	}	
	route.GET("Attendance/alllist", attendanceAPI.AllAttendanceListHandler)
	route.POST("Attendance/new", attendanceAPI.CreateAttendanceHandeler)
	route.GET("Attendance/GETONE", attendanceAPI.GETONEAttendanceHandeler)
	route.PUT("Attendance/Checkname",attendanceAPI.ChecknameHandeler)
	

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


	//QRCode
	// qrcodeRepository := repository.QRCodeRepositoryMongo{
	// 	ConnectionDB: connectionDB,
	// }
	// qrcodeAPI := api.QRCodeAPI{
	// 	QRCodeRepository: &sqrcodeRepository,
	// }	
	// route.POST("QRCODE/create", QRcodeAPI.AddSubjectHandeler)
	

}