package api

import (
	"Go-backend-mongoDB/model"
	"Go-backend-mongoDB/repository"
	"log"
	"net/http"
	// "encoding/json"
	// b64 "encoding/base64"
	// "github.com/skip2/go-qrcode"
	"github.com/gin-gonic/gin"
	// "github.com/globalsign/mgo/bson"
	"encoding/json"
	"fmt"
	"strconv"
	"math/rand"
	"time"
	"strings"
	
)
//SubjectAPI is ...
type SubjectAPI struct {
	SubjectRepository repository.SubjectRepository
}
//SubjectListHandler is ...
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
//AddSubjectHandeler is ...
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
//EditDescriptionHandler is ...
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
//DeleteSubjectHandler is ...
func (api SubjectAPI) DeleteSubjectHandler(context *gin.Context) {
	id := context.Param("id")
	err := api.SubjectRepository.DeleteSubject(id)
	if err != nil {
		log.Println("error DeleteSubjectHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	context.JSON(http.StatusNoContent, gin.H{"message": "Success"})
}
//GETMySubjectHandeler is ...
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
//JoinClassHandeler is ...
func (api SubjectAPI) JoinClassHandeler(context *gin.Context) {
	var join model.JoinSubject
	err := context.ShouldBindJSON(&join)
	fmt.Println(join)
	subjectUN, err := json.Marshal(join)
	fmt.Println(subjectUN)
	str := string(subjectUN)
	fmt.Println(str)
	if err != nil {
		log.Println("error JoinClassHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message1": err.Error()})
		return
	}
	i, _  := strconv.Atoi(join.Password)
	fmt.Println(i)
	if err != nil {
		log.Println("error JoinClass", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	timeinclass, err:= api.SubjectRepository.GETTime(i)
	if err != nil {
		log.Println("error GETTime", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	for _, copy := range timeinclass {
		for _, copy2 := range copy.TstudentInfo {
			if copy2.StudentID  == join.StudentID{
			fmt.Println("joined")
			return
		}
		}

	}


			
	fmt.Println("joining")
	err = api.SubjectRepository.JoinClass(i, join)
	
	context.JSON(http.StatusOK, "WTF!")
}
//GETONESubjectHandeler is ...
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
//RandPassClass is ...
func (api SubjectAPI) RandPassClass() (i int){
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
	fmt.Println("subjects = ",subjects)
	num1 := classnum
	if(num1 < 1000){
		api.RandPassClass()
	}

	subjectsInfo.Subject = subjects
	// var state int = 0
	for _, copy := range subjects{
		
		fmt.Println(copy.TSpassword)
		if (copy.TSpassword == classnum){
			api.RandPassClass()
		}else{
			return classnum
		}
	}
	
	return classnum
}
//GETStudentinClassHandeler is ...
func (api SubjectAPI) GETStudentinClassHandeler(context *gin.Context) {
	id:= context.Param("pass")
	fmt.Println(id)
	i, err := strconv.Atoi(id)
	studentinclass, err:= api.SubjectRepository.GETStudentinClass(i)
	if err != nil {
		log.Println("error GETStudentinClassHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(studentinclass)
	for _, copy := range studentinclass {
		context.JSON(http.StatusOK, copy.TstudentInfo)
	}
}
//GETStudentinfoClassHandeler is ...
func (api SubjectAPI) GETStudentinfoClassHandeler(context *gin.Context) {
	id:= context.Param("pass")
	fmt.Println(id)
	i, err := strconv.Atoi(id)
	studentinclass, err:= api.SubjectRepository.GETStudentinClass(i)
	if err != nil {
		log.Println("error GETStudentinClassHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(studentinclass)
	for _, copy := range studentinclass {
		context.JSON(http.StatusOK, copy.TstudentInfo)
	}
}
// UserJoinHandler is ...
func (api SubjectAPI) UserJoinHandler(context *gin.Context) {
	var user model.JoinUser
	err := context.ShouldBindJSON(&user)
	fmt.Println(user)
	if err != nil {
		log.Println("error UserJoinHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.SubjectRepository.UserJoin(user)
	if err != nil {
		log.Println("error UserJoin", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, "Success")

}
//GETTimeHandeler is ...
func (api SubjectAPI) GETTimeHandeler(context *gin.Context) {
	id:= context.Param("pass")
	fmt.Println(id)
	i, err := strconv.Atoi(id)
	timeinclass, err:= api.SubjectRepository.GETTime(i)
	
	if err != nil {
		log.Println("error GETTimeHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(timeinclass)
	for _, copy := range timeinclass {
		context.JSON(http.StatusOK, copy.TStimesubject)
	}
}
//EditTimelimitHandler is ...
func (api SubjectAPI) EditTimelimitHandler(context *gin.Context) {
	var subject model.Subject
	var time model.TimeLimitandtimeOut
	err := context.ShouldBindJSON(&time)
	if err != nil {
		log.Println("error EditTimelimitHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	pass:= context.Param("pass")
	fmt.Println(pass)
	id, err := strconv.Atoi(pass)
	timeinclass, err:= api.SubjectRepository.GETTime(id)
	if err != nil {
		log.Println("error GETTime", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	
	for _, copy := range timeinclass {
		time.TSlimit = copy.TStimesubject
		time.TStimeout = copy.TStimesubject
		
		
	}
	fmt.Println("time",time)
	fmt.Println("time.TSlimit",time.TSlimit)
	fmt.Println("time.TStimeout",time.TStimeout)
	// time.Timetemp = "10"
	// time.TimeoutTemp = "90"
	var temptime model.TimeLimitandtimeOut
	for _, copy := range time.TSlimit {
		fmt.Println(copy)
		str := strings.Split(copy.Time, "-")
		fmt.Println(str)
		str2 := strings.Split(str[0], ".")
		fmt.Println(str2)
		h, err := strconv.Atoi(str2[0])
		m, err := strconv.Atoi(str2[1])
		j, err := strconv.Atoi(time.Timetemp)
		fmt.Println(err)
		var h1 int
		if j/60 !=0 {
			h1 = j/60
			
		}
		j1 := j%60
		h = h+h1
		x :=m+j1
		s0 := strconv.Itoa(x)
		s1 := strconv.Itoa(h)
		copy.Time = s1+"."+s0+"-"+str[1]
		temptime.TSlimit = append(temptime.TSlimit,copy) 
		fmt.Println(copy.Time)
	}
	time.TSlimit = temptime.TSlimit

	for _, copy := range time.TStimeout {
		str := strings.Split(copy.Time, "-")
		str2 := strings.Split(str[0], ".")
		h, err := strconv.Atoi(str2[0])
		m, err := strconv.Atoi(str2[1])
		j, err := strconv.Atoi(time.TimeoutTemp)
		fmt.Println(err)
		var h1 int
		if j/60 !=0 {
			h1 = j/60
			
		}
		j1 := j%60
		h = h+h1
		x :=m+j1
		s0 := strconv.Itoa(x)
		s1 := strconv.Itoa(h)
		copy.Time = s1+"."+s0+"-"+str[1]
		temptime.TStimeout = append(temptime.TStimeout,copy) 
		fmt.Println(copy.Time)
	}
	time.TStimeout = temptime.TStimeout
	
	fmt.Println("time.Limit",time.TSlimit)
	fmt.Println("time.Timeout",time.TStimeout)
	
	

	err = api.SubjectRepository.EditTimeLimit(id ,subject,time)
	if err != nil {
		log.Println("error EditTimelimitHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, "Success")
}
