package api

import (
	"Go-backend-mongoDB/model"
	"Go-backend-mongoDB/repository"
	"Go-backend-mongoDB/websocket"
	b64 "encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

var clientid string

//AttendanceAPI is ...
type AttendanceAPI struct {
	AttendanceRepository repository.AttendanceRepository
}

//AllAttendanceListHandler is ...
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

//CreateAttendanceHandeler is ...
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

//GETONEAttendanceHandeler is ...
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
	oneattendance, err2 := api.AttendanceRepository.GetAttendance(attendance.ASpassword)
	if err2 != nil {
		log.Println("error GetAttendance", err2.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err2.Error()})
		return
	}
	context.JSON(http.StatusOK, oneattendance)
}

//CreateQrcodeAndAttendanceHandeler is ...
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
	png, _ = qrcode.Encode(qr.Time+";"+qr.User+";"+str+";"+str2+";"+qr.Clientid, qrcode.Medium, 512)
	sEnc := b64.StdEncoding.EncodeToString([]byte(png))
	aqr := qr.Time + ";" + qr.User + ";" + str + ";" + str2 + ";" + qr.Clientid
	fmt.Println(sEnc)
	fmt.Println(aqr)
	var NewAttendance model.Attendance
	NewAttendance.Date = qr.Time
	NewAttendance.ASpassword = str
	NewAttendance.PicQRcode = sEnc
	NewAttendance.AQRcode = aqr
	NewAttendance.ATimeAuthen = timeAuthens
	NewAttendance.Clientid = qr.Clientid

	err = api.AttendanceRepository.CreateAttendance(NewAttendance)
	if err != nil {
		log.Println("error CreateAttendance", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	fmt.Println("-----------------", clientid, "-----------------")
	context.JSON(http.StatusCreated, "Success")

}

//GETQRcoderHandeler is ...
func (api AttendanceAPI) GETQRcoderHandeler(context *gin.Context) {
	var attendancesInfo model.AttendanceInfo
	pass := context.Param("pass")
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

//GetshowQrCode is ...
func (api AttendanceAPI) GetshowQrCode(context *gin.Context) {
	pass := context.Param("passOfcouse")
	defer context.Request.Body.Close()
	flag.Parse()
	i, _ := strconv.Atoi(pass)
	var sp []model.CreateQr

	for _, copy := range newQrs {
		if copy.Pass == i {
			sp = append(sp, copy)
		}
	}

	len := len(sp)
	fmt.Println(sp)
	var sp2 []model.CreateQr
	for _, copy := range sp {
		if copy.TimeAuthen == len {
			sp2 = append(sp2, copy)
		}
	}
	context.JSON(http.StatusOK, sp2)
	fmt.Println(len)
	fmt.Println(sp2)
}

//CheckQRcodeHandeler is ...
func (api AttendanceAPI) CheckQRcodeHandeler(context *gin.Context) {
	var check model.CheckQRcode
	err := context.ShouldBindJSON(&check)
	onesubject, err := api.AttendanceRepository.Checkcode(check.AQRcode)
	if check.AQRcode == "" {
		fmt.Println("No data")
	}
	if err != nil {
		log.Println("error LoginHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	for _, copy := range onesubject {
		if copy.AQRcode == check.AQRcode {
			context.JSON(http.StatusOK, "Success")
		}
	}

}

//CheckNameClassHandeler is ...
func (api AttendanceAPI) CheckNameClassHandeler(context *gin.Context) {
	var check model.CheckStudent
	err := context.ShouldBindJSON(&check)
	fmt.Println(check)
	if err != nil {
		log.Println("error CheckNameClassHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message1": err.Error()})
		return
	}
	err = api.AttendanceRepository.CheckingName(check)
	if err != nil {
		log.Println("error CheckingName", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, "Success")

}

//TimelimitHandeler is ...
func (api AttendanceAPI) TimelimitHandeler(context *gin.Context) {
	var time model.TimeLimit
	err := context.ShouldBindJSON(&time)
	fmt.Println(time)
	if err != nil {
		log.Println("error CheckNameClassHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message1": err.Error()})
		return
	}

	str := strings.Split(time.Date, ",")
	fmt.Println(str[2])

	str2 := strings.Split(str[2], ":")
	str3 := strings.Split(str2[0], " ")
	fmt.Println(str3)
	i, err := strconv.Atoi(str2[1])
	fmt.Println("str2[1] is -----" + str2[1])
	fmt.Println("i is -----", i)
	s, err := strconv.Atoi(str3[1])
	fmt.Println("str2[0] is -----" + str2[0])
	fmt.Println("j is -----++++", s)
	var s0, s1 string
	if i+30 >= 60 {
		s1 = strconv.Itoa((i + 30) - 60)
		s0 = strconv.Itoa(s + 1)
	} else {
		s0 = strconv.Itoa(s)
		s1 = strconv.Itoa((i + 30))
	}

	limit := s0 + ":" + s1 + ":" + str2[2]
	fmt.Println("Start time is -----" + str[2])
	fmt.Println("limit is -----" + limit)

	context.JSON(http.StatusOK, limit)
}

//GETSelfieHandeler is ...
func (api AttendanceAPI) GETSelfieHandeler(context *gin.Context) {
	var attendancesInfo model.AttendanceInfo
	pass := context.Param("pass")
	date := context.Param("date")
	timeauthens := context.Param("timeauthens")
	inttimeauthens, err := strconv.Atoi(timeauthens)
	defer context.Request.Body.Close()
	attendances, err := api.AttendanceRepository.GetSelfie(pass, date, inttimeauthens)
	if err != nil {
		log.Println("error GETSelfieHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	attendancesInfo.Attendance = attendances
	context.JSON(http.StatusOK, attendances)

}

//GETinfoStudentHandeler is ...
func (api AttendanceAPI) GETinfoStudentHandeler(context *gin.Context) {
	var attendancesInfo model.AttendanceInfo
	pass := context.Param("pass")
	date := context.Param("date")
	timeauthens := context.Param("timeauthens")
	inttimeauthens, err := strconv.Atoi(timeauthens)
	defer context.Request.Body.Close()
	attendances, err := api.AttendanceRepository.GetInfoStudent(pass, date, inttimeauthens)
	if err != nil {
		log.Println("error GETSelfieHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	attendancesInfo.Attendance = attendances
	for _, copy := range attendances {

		context.JSON(http.StatusOK, copy.Astudent)

	}

}
//GETClientHandeler is ...
func (api AttendanceAPI) GETClientHandeler(context *gin.Context) {
	var attendancesInfo model.AttendanceInfo
	clientid := context.Param("clientid")
	defer context.Request.Body.Close()
	attendances, err := api.AttendanceRepository.GetClient(clientid)
	if err != nil {
		log.Println("error GETSelfieHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	attendancesInfo.Attendance = attendances
	for _, copy := range attendances {

		context.JSON(http.StatusOK, copy.Astudent)

	}
}


type newDataSocket struct {
	// Clientid string               `json:"clientId"`
	Astudent []model.CheckStudent `json:"Astudent"`
}

//Message2 ...
type Message2 struct {
	Sender    string          `json:"sender,omitempty"`
	Recipient string          `json:"recipient,omitempty"`
	Content   []newDataSocket `json:"content,omitempty"`
}

//CheckwithSocketHandeler is ...
func (api AttendanceAPI) CheckwithSocketHandeler(context *gin.Context) {
	var check model.CheckStudent
	err := context.ShouldBindJSON(&check)
	fmt.Println(check)
	if err != nil {
		log.Println("error CheckNameClassHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message1": err.Error()})
		return
	}
	attendances, err := api.AttendanceRepository.CheckingNamewithSocket(check)
	if err != nil {
		log.Println("error CheckingName", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, "Success")
	var clientid string
	var data1 newDataSocket
	var data2 []newDataSocket
	for _, copy := range attendances {
		clientid = copy.Clientid
		data1.Astudent = copy.Astudent
		fmt.Println("data1 = ", data1)
		data2 = append(data2, data1)
	}
	// out, err := json.Marshal(data2)
	// fmt.Printf("Out(type) = %T \n", out)
	// if err != nil {
	// 	panic(err)
	// }
	jsonMessage, _ := json.Marshal(&Message2{Sender: clientid, Content: data2})
	websocket.Manager.Broadcast <- jsonMessage
	// context.JSON(http.StatusOK, data2)
}
