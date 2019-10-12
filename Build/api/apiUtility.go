package api

import (
	
	"Go-backend-mongoDB/repository"

	"Go-backend-mongoDB/model"
	b64 "encoding/base64"
	"github.com/skip2/go-qrcode"
	"github.com/gin-gonic/gin"
	// "fmt"
	"flag"
	"net/http"
	
	"strconv"
	// "log"
)	
//UtilityAPI is ...
type UtilityAPI struct {
	AttendanceRepository repository.AttendanceRepository
	UserRepository repository.UserRepository
	SubjectRepository repository.SubjectRepository
	StudentRepository repository.StudentRepository
}

var newQrs []model.CreateQr
//CreateBarcode is ...
func (api UtilityAPI) CreateBarcode(context *gin.Context)  (string){
	var png []byte
	// var qrIN model.CreateQr
	var qr model.Qrcode
	defer context.Request.Body.Close()
	
	str := strconv.Itoa(qr.Pass)
	var copyMyQr []model.CreateQr
	var timeAuthens int = 0


	timeAuthens = len(copyMyQr) + 1
	str2 := strconv.Itoa(timeAuthens)

	png, _ = qrcode.Encode(qr.Time+";"+qr.User+";"+str+";"+str2, qrcode.Medium, 512)

	sEnc := b64.StdEncoding.EncodeToString([]byte(png))
	
	
	return sEnc

	
	
}
//AllQr is ...
func AllQr(c *gin.Context) {
	defer c.Request.Body.Close()
	c.JSON(http.StatusOK, newQrs)
}
//MyQr is ...
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

