package api

import (
	b64 "encoding/base64"
	"fmt"
	"log"
	"net/http"
	_ "net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/backend/models"
	"github.com/skip2/go-qrcode"

	// "strconv"
	"flag"
)

var newQrs []models.CreateQr

func CreateBarcode(c *gin.Context) {
	var png []byte
	var p models.Qrcode
	defer c.Request.Body.Close()
	if c.ShouldBindJSON(&p) != nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(p.Time)
		log.Println(p.User)
		log.Println(p.Pass)
	}

	str := strconv.Itoa(p.Pass)
	var copyMyQr []models.CreateQr
	var timeAuthens int = 0

	for _, sp := range newQrs {
		if p.Pass == sp.Pass {
			copyMyQr = append(copyMyQr, sp)
		}
	}

	timeAuthens = len(copyMyQr) + 1
	str2 := strconv.Itoa(timeAuthens)

	png, _ = qrcode.Encode(p.Time+";"+p.User+";"+str+";"+str2, qrcode.Medium, 512)
	sEnc := b64.StdEncoding.EncodeToString([]byte(png))
	sp := models.CreateQr{
		Id:         len(newQrs) + 1,
		Qrcode:     sEnc,
		Time:       p.Time,
		Pass:       p.Pass,
		TimeAuthen: timeAuthens,
	}

	fmt.Println(sp)
	newQrs = append(newQrs, sp)
	c.JSON(http.StatusOK, "Success")
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
	var sp []models.CreateQr
	for _, copy := range newQrs {
		if copy.Pass == i {
			sp = append(sp, copy)
		}
	}

	var sp2 []models.CreateQr
	index := len(sp)
	for i := index; i >= 0; i-- {
		for _, copy := range sp {
			if index == copy.TimeAuthen {
				sp2 = append(sp2, copy)
				index = index-1
			}
		}
	}

	c.JSON(http.StatusOK, sp2)
}

func GetshowQrCode(c *gin.Context) {
	pass := c.Param("passOfcouse")
	defer c.Request.Body.Close()
	flag.Parse()
	i, _ := strconv.Atoi(pass)
	var sp []models.CreateQr
	for _, copy := range newQrs {
		if copy.Pass == i {
			sp = append(sp, copy)
		}
	}

	len := len(sp)
	fmt.Println(sp)
	var sp2 []models.CreateQr
	for _, copy := range sp {
		if copy.TimeAuthen == len {
			sp2 = append(sp2, copy)
		}
	}
	c.JSON(http.StatusOK, sp2)
	fmt.Println(len)
	fmt.Println(sp2)
}
