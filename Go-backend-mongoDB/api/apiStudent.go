package api

import (
	"Go-backend-mongoDB/model"
	"Go-backend-mongoDB/repository"
	"log"
	"net/http"
	// b64 "encoding/base64"
	// "github.com/skip2/go-qrcode"
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)
//StudentAPI is ...
type StudentAPI struct {
	StudentRepository repository.StudentRepository
	AttendanceRepository repository.AttendanceRepository
	
}
//AllStudentListHandler is ...
func (api StudentAPI) AllStudentListHandler(context *gin.Context) {
	var studentInfo model.StudentInfo
	student, err := api.StudentRepository.GetAllStudent()
	if err != nil {
		log.Println("error UserListHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	studentInfo.Student = student
	context.JSON(http.StatusOK, studentInfo)
}
//CreateStudentHandeler is ...
func (api StudentAPI) CreateStudentHandeler(context *gin.Context) {
	var student model.Student 
	err := context.ShouldBindJSON(&student)
	if err != nil {
		log.Println("error CreateStudentHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.StudentRepository.CreateStudent(student)
	if err != nil {
		log.Println("error CreateStudentHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"status": "susess"})

}
//ShowsubjectlistStudentHandeler is ...
func (api StudentAPI) ShowsubjectlistStudentHandeler(context *gin.Context) {
	var studentInfo model.StudentInfo
	student, err := api.StudentRepository.GetAllStudent()
	if err != nil {
		log.Println("error UserListHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	studentInfo.Student = student
	context.JSON(http.StatusOK, studentInfo)
}
//StudentLoginHandler is ...
func (api StudentAPI) StudentLoginHandler(context *gin.Context) {
	var user model.Login
	err := context.ShouldBindJSON(&user)	
	oneStudent, err:= api.StudentRepository.LoginStudent(user.UserName)
	if err != nil {
		log.Println("error LoginHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	for _, copy := range oneStudent {
		if copy.Susername == user.UserName && copy.Spassword == user.Password {
			context.JSON(http.StatusOK, oneStudent)
		}else{
			context.JSON(http.StatusOK, "Wrong")
		}
	}
	
	
}
//StudentinfoListHandler is ...
func (api StudentAPI) StudentinfoListHandler(context *gin.Context) {
	var user model.Student
	username:= context.Param("id")
	fmt.Println(user)
	fmt.Println(username)
	oneuser, err2:= api.StudentRepository.Getbyid(username)
	if err2 != nil {
		log.Println("error GetUser", err2.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message2": err2.Error()})
		return
	}
	context.JSON(http.StatusOK, oneuser)
}

//StudenSubjectListHandler is ...
func (api StudentAPI) StudenSubjectListHandler(context *gin.Context) {
	var user model.Student
	sid:= context.Param("sid")
	fmt.Println(user)
	fmt.Println(sid)
	oneuser, err:= api.StudentRepository.GetsubjectbySid(sid)
	if err != nil {
		log.Println("error GetsubjectbySid", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	var users model.NewSubjectName
	var newUser []model.NewSubjectName
	for _, copy := range oneuser {
		users.TSNames = copy.TSName
		users.TSpassword = copy.TSpassword
		newUser = append(newUser, users)
	}
	fmt.Println(newUser)
	context.JSON(http.StatusOK, newUser)
}
// GetAttendancebyPassandID is ...
func (api StudentAPI) GetAttendancebyPassandID(context *gin.Context) {
	var attendancesInfo model.AttendanceInfo
	pass := context.Param("pass")
	sid := context.Param("sid")
	defer context.Request.Body.Close()
	attendances, err := api.StudentRepository.GetStudentAttendance(pass)
	if err != nil {
		log.Println("error GetAttendance", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	i, err := strconv.Atoi(pass)
	if err != nil {
		log.Println("error Atoi", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	studentinclass, err := api.StudentRepository.GETStudenClasses(i)
	if err != nil {
		log.Println("error GETStudentinClass", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	
	var NewDataAttendancebyPassIn model.NewDataAttendancebyPass
	var NewDataAttendancebyPassOut []model.NewDataAttendancebyPass

	var AuthenStudentIn model.AuthenStudent
	var AuthenStudentOut []model.AuthenStudent

	attendancesInfo.Attendance = attendances
	for _, copy := range attendances {
		// fmt.Println(copy.Date)
		AuthenStudentIn.Date = copy.Date
		AuthenStudentOut = append(AuthenStudentOut, AuthenStudentIn)
	}
	NewDataAttendancebyPassIn.AuthenStudents = AuthenStudentOut

	for _, copy := range studentinclass {
		for _, copy2 := range copy.TstudentInfo {
			NewDataAttendancebyPassIn.IDStudent = copy2.StudentID
			NewDataAttendancebyPassIn.NameStudent = copy2.SfirstName + "  " + copy2.SlastName
			NewDataAttendancebyPassOut = append(NewDataAttendancebyPassOut, NewDataAttendancebyPassIn)
		}
	}
	var AuthenStudentOut2 []model.AuthenStudent
	for j := 0; j < len(NewDataAttendancebyPassOut); j++ {
		// fmt.Println("NewDataAttendancebyPassOut")
		len := len(NewDataAttendancebyPassOut[j].AuthenStudents)
		for i := 0; i < len; i++ {
			dataStudent := NewDataAttendancebyPassOut[j].AuthenStudents[i]
			data := attendances
			dataCheck := data[i]

			// fmt.Println("dataStudent = ", dataStudent)
			// fmt.Println("dataCheck = ", dataCheck.Date)
			data2 := data[i].Astudent

			state := 0
			if dataCheck.Date == dataStudent.Date {
				for _, copyAstudent := range data2 {
					NameStudent := copyAstudent.SfirstName + "  " + copyAstudent.SlastName
					if (NameStudent == NewDataAttendancebyPassOut[j].NameStudent) && (copyAstudent.StudentID == NewDataAttendancebyPassOut[j].IDStudent) {
						state = 1
					}
				}
			}
			if state == 1 {
				dataStudent.StateAuthen = true
			}
		
			AuthenStudentOut2 = append(AuthenStudentOut2, dataStudent)
		}
		NewDataAttendancebyPassOut[j].No = j+1
		NewDataAttendancebyPassOut[j].AuthenStudents = AuthenStudentOut2
		AuthenStudentOut2 = nil
		// fmt.Println("AuthenStudentOut2 = ", AuthenStudentOut2)
	}
	for _, copy := range NewDataAttendancebyPassOut {
			if sid == copy.IDStudent {
				fmt.Println("NewDataAttendancebyPassOut = ", copy)
				context.JSON(http.StatusOK, copy)
			}
		}
}
//DeleteStudentHandler is ...
func (api StudentAPI) DeleteStudentHandler(context *gin.Context) {
	id := context.Param("id")
	err := api.StudentRepository.DeleteStudent(id)
	if err != nil {
		log.Println("error DeleteStudentHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	context.JSON(http.StatusNoContent, gin.H{"message": "Success"})
}
