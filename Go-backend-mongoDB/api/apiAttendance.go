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

func (api AttendanceAPI) FindbyANameAttendanceHandler(context *gin.Context) {
	var attendancesInfo model.AttendanceInfo
	aName := context.Param("AName")
	attendances, err := api.AttendanceRepository.FindbyAName(aName)
	if err != nil {
		log.Println("error FindbyANameAttendanceHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	attendancesInfo.Attendance = attendances
	context.JSON(http.StatusNoContent, gin.H{"message": "susess"})
}