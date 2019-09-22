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
//Student
type StudentAPI struct {
	StudentRepository repository.StudentRepository
}

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