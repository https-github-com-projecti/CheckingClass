package api

import (
	"Go-backend-mongoDB/model"
	"Go-backend-mongoDB/repository"
	"log"
	"net/http"
	// b64 "encoding/base64"
	// "github.com/skip2/go-qrcode"

	"github.com/gin-gonic/gin"
)
//StudentAPI is ...
type StudentAPI struct {
	StudentRepository repository.StudentRepository
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
			context.JSON(http.StatusOK, "Success")
		}
	}
	
	
}