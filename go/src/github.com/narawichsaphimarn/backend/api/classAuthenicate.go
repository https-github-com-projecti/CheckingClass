package api

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	b64 "encoding/base64"
	"fmt"
	"net/http"
	_ "net/http"
	"github.com/narawichsaphimarn/backend/models"
)

var newQrs []models.CreateQr

func CreateBarcode(c *gin.Context){
	var png []byte
	var p models.Qrcode
	defer c.Request.Body.Close()
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, "Can't creat class")
	}

	newQr := models.Qrcode{
		Time: p.Time,
		User: p.User,
		Pass: p.Pass,
	}

	fmt.Println(newQr)
	png, _ = qrcode.Encode(p.Time + p.User + p.Pass, qrcode.Medium, 256)
	sEnc := b64.StdEncoding.EncodeToString([]byte(png))
	fmt.Println("test = " + sEnc)
	sp := models.CreateQr{
		Id:     len(newQrs)+1,
		Qrcode: sEnc,
		Time: p.Time,
	}

	newQrs = append(newQrs , sp)
	c.JSON(http.StatusOK, newQrs)
}


