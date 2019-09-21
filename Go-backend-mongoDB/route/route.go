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
	route.GET("user/list", userAPI.UserListHandler)
	route.POST("user/Add", userAPI.AddUserHandeler)
	route.PUT("user/Edit/:user_id", userAPI.EditPasswordHandler)
	route.DELETE("user/Delete/:user_id", userAPI.DeleteUserByIDHandler)
	

	
	//SUBject
	subjectRepository := repository.SubjectRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	subjectAPI := api.SubjectAPI{
		SubjectRepository: &subjectRepository,
	}	
	route.GET("Subject/list", subjectAPI.SubjectListHandler)
	route.POST("Subject/Add", subjectAPI.AddSubjectHandeler)
	route.PUT("Subject/Edit/:subject_id", subjectAPI.EditDescriptionHandler)
	route.DELETE("Subject/Delete/:subject_id", subjectAPI.DeleteSubjectByIDHandler)
	route.PUT("Subject/Editaddstudent/:subject_id", subjectAPI.EditaddstudentHandler)

	//QRCode
	// qrcodeRepository := repository.QRCodeRepositoryMongo{
	// 	ConnectionDB: connectionDB,
	// }
	// qrcodeAPI := api.QRCodeAPI{
	// 	QRCodeRepository: &sqrcodeRepository,
	// }	
	// route.POST("QRCODE/create", QRcodeAPI.AddSubjectHandeler)
	

	//Attendance	
	attendanceRepository := repository.AttendanceRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	attendanceAPI := api.AttendanceAPI{
		AttendanceRepository: &attendanceRepository,
	}	
	route.GET("Attendance/alllist", attendanceAPI.AllAttendanceListHandler)
	route.POST("Attendance/new", attendanceAPI.CreateAttendanceHandeler)
	route.GET("Attendance/list/:AName", attendanceAPI.FindbyANameAttendanceHandler)

	//Student
	studentRepository := repository.StudentRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	studentAPI := api.StudentAPI{
		StudentRepository: &studentRepository,
	}	
	route.GET("Student/alllist", studentAPI.AllStudentListHandler)


}