package api

import (
	"Go-backend-mongoDB/model"
	"Go-backend-mongoDB/repository"
	"log"
	"net/http"
	"encoding/json"
	// b64 "encoding/base64"
	// "github.com/skip2/go-qrcode"
	"github.com/gin-gonic/gin"
	// "github.com/globalsign/mgo/bson"
	// "encoding/json"
	"fmt"
	"strconv"
	"math/rand"
	
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

	subject.TSpassword = api.RandPassClass() 

	err = api.SubjectRepository.AddSubject(subject)
	if err != nil {
		log.Println("error AddSubjectHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, "Success")
}
func (api SubjectAPI) EditDescriptionHandler(context *gin.Context) {
	var subject model.Subject
	err := context.ShouldBindJSON(&subject)
	if err != nil {
		log.Println("error EditDescriptionHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.SubjectRepository.EditDescription(subject.TSID ,subject)
	if err != nil {
		log.Println("error EditDescriptionHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "susess"})
}
func (api SubjectAPI) DeleteSubjectHandler(context *gin.Context) {
	var subject model.Subject
	err := context.ShouldBindJSON(&subject)
	fmt.Println(subject)
	if err != nil {
		log.Println("error DeleteSubjectHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.SubjectRepository.DeleteSubject(subject.TSID)
	if err != nil {
		log.Println("error DeleteSubjectHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	context.JSON(http.StatusNoContent, gin.H{"message": "susess"})
}

func (api SubjectAPI) GETMySubjectHandeler(context *gin.Context) {
	id:= context.Param("id")
	fmt.Println(id)
	onesubject, err:= api.SubjectRepository.GetSubject(id)
	if err != nil {
		log.Println("error GETONESubjectHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, onesubject)
}
func (api SubjectAPI) JoinClassHandeler(context *gin.Context) {
	var subject model.JoinSubject
	err := context.ShouldBindJSON(&subject)
	fmt.Println(subject)
	subjectUN, err := json.Marshal(subject)
	// fmt.Println(subjectUN)
	str := string(subjectUN)
	fmt.Println(str)
	if err != nil {
		log.Println("error DeleteSubjectHandler1", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message1": err.Error()})
		return
	}
	i, _  := strconv.Atoi(subject.Jpassword)
	err = api.SubjectRepository.JoinClass(i, subject.JSID)
	if err != nil {
		log.Println("error JoinClassHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, "WTF!")
}
func (api SubjectAPI) GETONESubjectHandeler(context *gin.Context) {
	id:= context.Param("id")
	fmt.Println(id)
	onesubject, err:= api.SubjectRepository.GetOneSubject(id)
	if err != nil {
		log.Println("error GETONESubjectHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, onesubject)
}


///TEST
func (api SubjectAPI) RandPassClass() (i int){
	classnum := rand.Intn(10000)
	var subjectsInfo model.SubjectInfo
	subjects, err := api.SubjectRepository.GetAllSubject()
	if err != nil {
		log.Println("error SubjectListHandler", err.Error())
		
		return
	}
	subjectsInfo.Subject = subjects
	for _, copy := range subjects {
		if (classnum == copy.TSpassword) {
			api.RandPassClass()
		}
	}
	i = classnum
	return 
}

