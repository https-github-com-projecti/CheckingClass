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
//Subject
type SubjectAPI struct {
	SubjectRepository repository.SubjectRepository
}
func (api SubjectAPI) SubjectListHandler(context *gin.Context) {
	var subjectsInfo model.SubjectInfo
	subjects, err := api.SubjectRepository.GetAllSubject()
	if err != nil {
		log.Println("error SubjectListHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	subjectsInfo.Subject = subjects
	context.JSON(http.StatusOK, subjectsInfo)
}
func (api SubjectAPI) AddSubjectHandeler(context *gin.Context) {
	var subject model.Subject
	err := context.ShouldBindJSON(&subject)
	if err != nil {
		log.Println("error AddSubjectHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.SubjectRepository.AddSubject(subject)
	if err != nil {
		log.Println("error AddSubjectHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"status": "susess"})
}
func (api SubjectAPI) EditDescriptionHandler(context *gin.Context) {
	var subject model.Subject
	subjectID := context.Param("subject_id")
	err := context.ShouldBindJSON(&subject)
	if err != nil {
		log.Println("error EditDescriptionHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.SubjectRepository.EditDescription(subjectID, subject)
	if err != nil {
		log.Println("error EditDescriptionHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "susess"})
}
func (api SubjectAPI) DeleteSubjectByIDHandler(context *gin.Context) {
	subjectID := context.Param("subject_id")
	err := api.SubjectRepository.DeleteSubjectByID(subjectID)
	if err != nil {
		log.Println("error DeleteSubjectByIDHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	context.JSON(http.StatusNoContent, gin.H{"message": "susess"})
}

func (api SubjectAPI) EditaddstudentHandler(context *gin.Context) {
	var subject model.Subject
	subjectID := context.Param("subject_id")
	err := context.ShouldBindJSON(&subject)
	if err != nil {
		log.Println("error EditaddstudentHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.SubjectRepository.Editaddstudent(subjectID, subject)
	if err != nil {
		log.Println("error EditaddstudentHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "susess"})
}
