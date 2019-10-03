package api

import (
	"Go-backend-mongoDB/model"
	"Go-backend-mongoDB/repository"
	"encoding/json"
	"log"
	"net/http"

	// b64 "encoding/base64"
	// "github.com/skip2/go-qrcode"
	"github.com/gin-gonic/gin"
	// "github.com/globalsign/mgo/bson"
	// "encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
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
	err = api.SubjectRepository.EditDescription(subject.TSID, subject)
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
	id := context.Param("id")
	fmt.Println(id)
	onesubject, err := api.SubjectRepository.GetSubject(id)
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
	i, _ := strconv.Atoi(subject.Jpassword)
	err = api.SubjectRepository.JoinClass(i, subject.JSID)
	if err != nil {
		log.Println("error JoinClassHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, "WTF!")
}
func (api SubjectAPI) GETONESubjectHandeler(context *gin.Context) {
	id := context.Param("id")
	fmt.Println(id)
	onesubject, err := api.SubjectRepository.GetOneSubject(id)
	if err != nil {
		log.Println("error GETONESubjectHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, onesubject)
}

///TEST
func (api SubjectAPI) RandPassClass() (i int) {
	// Uses default seed of 1, result will be 81:
	fmt.Println(rand.Intn(10000))
	rand.Intn(10000)
	rand.Seed(time.Now().UnixNano())
	classnum := rand.Intn(10000)

	fmt.Println(classnum)
	var subjectsInfo model.SubjectInfo
	subjects, err := api.SubjectRepository.GetAllSubject()
	if err != nil {
		log.Println("error SubjectListHandler", err.Error())

		return
	}
	fmt.Println("subjects = ", subjects)
	var num1 = classnum
	var num int
	for num1 > 0 {
		num1 = num1 / 10
		num = num + 1
	}
	subjectsInfo.Subject = subjects
	// var state int = 0
	for _, copy := range subjects {

		fmt.Println(copy.TSpassword)

		if num < 4 {
			api.RandPassClass()
		}

		if copy.TSpassword != classnum {
			i = classnum
			return i
		} else {
			api.RandPassClass()
		}
	}

	// if state == 1{
	// 	api.RandPassClass()
	// }
	// i = classnum
	return
}
