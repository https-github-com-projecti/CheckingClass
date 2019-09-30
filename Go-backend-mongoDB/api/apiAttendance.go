package api

import (
	"Go-backend-mongoDB/model"
	"Go-backend-mongoDB/repository"
	"log"
	"net/http"
	// b64 "encoding/base64"
	// "github.com/skip2/go-qrcode"
	"fmt"
	"github.com/gin-gonic/gin"
)
//Attendance
type AttendanceAPI struct {
	AttendanceRepository repository.AttendanceRepository
}
func (api AttendanceAPI) AllAttendanceListHandler(context *gin.Context) {
	var attendancesInfo model.AttendanceInfo
	attendances, err := api.AttendanceRepository.GetAllAttendance()
	if err != nil {
		log.Println("error AttendanceListHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	attendancesInfo.Attendance = attendances
	context.JSON(http.StatusOK, attendancesInfo)
}
func (api AttendanceAPI) CreateAttendanceHandeler(context *gin.Context) {
	var attendance model.Attendance
	err := context.ShouldBindJSON(&attendance)
	if err != nil {
		log.Println("error CreateAttendanceHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.AttendanceRepository.CreateAttendance(attendance)
	if err != nil {
		log.Println("error CreateAttendanceHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"status": "susess"})

}
func (api AttendanceAPI) GETONEAttendanceHandeler(context *gin.Context) {
	var attendance model.Attendance
	err := context.ShouldBindJSON(&attendance)
	if err != nil {
		log.Println("error GETONEAttendanceHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(attendance)
	oneattendance, err2:= api.AttendanceRepository.GetAttendance(attendance.ASpassword)
	if err != nil {
		log.Println("error GetAttendance", err2.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err2.Error()})
		return
	}
	context.JSON(http.StatusOK, oneattendance)
}
func (api AttendanceAPI) ChecknameHandeler(context *gin.Context) {
	// var subject model.JoinSubject
	// err := context.ShouldBindJSON(&subject)
	// fmt.Println(subject)
	// if err != nil {
	// 	log.Println("error ChecknameHandeler", err.Error())
	// 	context.JSON(http.StatusInternalServerError, gin.H{"message1": err.Error()})
	// 	return
	// }
	// err = api.AttendanceRepository.JoinClass(subject.JTSID, subject.Jpassword, subject.JSID)
	// if err != nil {
	// 	log.Println("error JoinClassHandeler", err.Error())
	// 	context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 	return
	// }
	context.JSON(http.StatusOK, gin.H{"status": "susess"})
}

