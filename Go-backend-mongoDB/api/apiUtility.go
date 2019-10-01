package api

import (
	
	"Go-backend-mongoDB/repository"

	"Go-backend-mongoDB/model"
	b64 "encoding/base64"
	"github.com/skip2/go-qrcode"
	"github.com/gin-gonic/gin"
	"fmt"
	"flag"
	"net/http"
	_ "net/http"
	"strconv"
	// "log"
)	
type UtilityAPI struct {
	AttendanceRepository repository.AttendanceRepository
	UserRepository repository.UserRepository
	SubjectRepository repository.SubjectRepository
	StudentRepository repository.StudentRepository
}

var newQrs []model.CreateQr

func (api UtilityAPI) CreateBarcode(context *gin.Context)  (string){
	var png []byte
	// var qrIN model.CreateQr
	var qr model.Qrcode
	defer context.Request.Body.Close()
	
	str := strconv.Itoa(qr.Pass)
	var copyMyQr []model.CreateQr
	var timeAuthens int = 0

// 	// ดึงQrcodeออกมา เพื่อเช็คจำนวนการเช็คชื่อของแต่ละวิชา
// 	// for _, sp := range newQrs {
// 	// 	if p.Pass == sp.Pass {
// 	// 		copyMyQr = append(copyMyQr, sp)
// 	// 	}
// 	// }

	timeAuthens = len(copyMyQr) + 1
	str2 := strconv.Itoa(timeAuthens)

	png, _ = qrcode.Encode(qr.Time+";"+qr.User+";"+str+";"+str2, qrcode.Medium, 512)

	sEnc := b64.StdEncoding.EncodeToString([]byte(png))
	
	
	return sEnc

	
	
}

func AllQr(c *gin.Context) {
	defer c.Request.Body.Close()
	c.JSON(http.StatusOK, newQrs)
}

func MyQr(c *gin.Context) {
	flag.Parse()
	defer c.Request.Body.Close()
	pass := c.Param("pass")
	i, _ := strconv.Atoi(pass)
	var sp []model.CreateQr
	for _, copy := range newQrs {
		if copy.Pass == i {
			sp = append(sp, copy)
		}
	}
	c.JSON(http.StatusOK, sp)
}

func GetshowQrCode(c *gin.Context) {
	pass := c.Param("passOfcouse")
	defer c.Request.Body.Close()
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
	c.JSON(http.StatusOK, sp2)
	fmt.Println(len)
	fmt.Println(sp2)
}

