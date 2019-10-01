package api

import (
	"Go-backend-mongoDB/model"
	"Go-backend-mongoDB/repository"
	"log"
	"net/http"
	b64 "encoding/base64"
	"github.com/skip2/go-qrcode"
	"fmt"
	// "flag"
	"strconv"
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
	defer context.Request.Body.Close()
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
	defer context.Request.Body.Close()
	err := context.ShouldBindJSON(&attendance)
	if err != nil {
		log.Println("error GETONEAttendanceHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(attendance)
	oneattendance, err2:= api.AttendanceRepository.GetAttendance(attendance.ASpassword)
	if err2 != nil {
		log.Println("error GetAttendance", err2.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err2.Error()})
		return
	}
	context.JSON(http.StatusOK, oneattendance)
}
func (api AttendanceAPI) ChecknameHandeler(context *gin.Context) {
	
	context.JSON(http.StatusOK, gin.H{"status": "susess"})
}


func (api AttendanceAPI) CreateQrcodeAndAttendanceHandeler(context *gin.Context) {
	var png []byte
	var qr model.Qrcode
	defer context.Request.Body.Close()
	err := context.ShouldBindJSON(&qr)
	if err != nil {
		log.Println("error CreateQrcodeAndAttendanceHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(qr)
	str := strconv.Itoa(qr.Pass)
	var timeAuthens int = 0
	var attendancesInfo model.AttendanceInfo
	
	attendances, err := api.AttendanceRepository.GetONEAllAttendance(str)
	if err != nil {
		log.Println("error GetONEAllAttendance", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	attendancesInfo.Attendance = attendances
	fmt.Println(attendances)
	timeAuthens = len(attendances) + 1
	str2 := strconv.Itoa(timeAuthens)
	png, _ = qrcode.Encode(qr.Time+";"+qr.User+";"+str+";"+str2, qrcode.Medium, 256)
	sEnc := b64.StdEncoding.EncodeToString([]byte(png))
	fmt.Println(sEnc)

	var NewAttendance model.Attendance
	NewAttendance.Date = qr.Time
	NewAttendance.ASpassword = str
	NewAttendance.AQRcode = sEnc
	NewAttendance.ATimeAuthen = timeAuthens

	err = api.AttendanceRepository.CreateAttendance(NewAttendance)
	if err != nil {
		log.Println("error CreateAttendance", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated,  "Success")


}
func (api AttendanceAPI) GETQRcoderHandeler(context *gin.Context) {
	var attendancesInfo model.AttendanceInfo
	pass:= context.Param("pass")
	defer context.Request.Body.Close()
	attendances, err := api.AttendanceRepository.GetONEAllAttendance(pass)
	if err != nil {
		log.Println("error AttendanceListHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	
	attendancesInfo.Attendance = attendances
	context.JSON(http.StatusOK, attendances)
	
	
}

