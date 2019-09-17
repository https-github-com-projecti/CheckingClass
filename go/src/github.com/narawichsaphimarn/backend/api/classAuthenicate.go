package api

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	b64 "encoding/base64"
	// "fmt"
	"net/http"
	_ "net/http"
	"github.com/narawichsaphimarn/backend/models"
	"log"
)

var newQrs []models.CreateQr

func CreateBarcode(c *gin.Context){
	var png []byte
	var p models.Qrcode
	defer c.Request.Body.Close()
	if c.ShouldBindJSON(&p) != nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(p.Time)
		log.Println(p.User)
		log.Println(p.Pass)
	}

	png, _ = qrcode.Encode(p.Time + p.User + p.Pass, qrcode.Medium, 256)
	sEnc := b64.StdEncoding.EncodeToString([]byte(png))
	sp := models.CreateQr {
		Id 		:	len(newQrs) + 1,
		Qrcode 	: 	sEnc,
		Time 	: 	p.Time,
		User 	: 	p.User,
	}

	newQrs = append(newQrs , sp)
	c.JSON(http.StatusOK, "Success")
}

func AllQr(c *gin.Context){
	defer c.Request.Body.Close()
	c.JSON(http.StatusOK, newQrs)
}

func MyQr(c *gin.Context){
	user := c.Param("user")
	var sp []models.CreateQr
	for _, copy := range newQrs {
		if (copy.User == user){
			sp = append(sp, copy)
		}
	}
	c.JSON(http.StatusOK, sp)
}


