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
//USER
type UserAPI struct {
	UserRepository repository.UserRepository
}

func (api UserAPI) UserListHandler(context *gin.Context) {
	var usersInfo model.UserInfo
	users, err := api.UserRepository.GetAllUser()
	if err != nil {
		log.Println("error UserListHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	usersInfo.User = users
	context.JSON(http.StatusOK, usersInfo)
}

func (api UserAPI) AddUserHandeler(context *gin.Context) {
	var user model.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.Println("error AddUserHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.UserRepository.AddUser(user)
	if err != nil {
		log.Println("error AddUserHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, "susess")
}

func (api UserAPI) EditPasswordHandler(context *gin.Context) {
	var user model.User
	err := context.ShouldBindJSON(&user)
	fmt.Println(user)
	if err != nil {
		log.Println("error EditPasswordHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.UserRepository.EditPassword(user.UserName, user.TPassword)
	if err != nil {
		log.Println("error EditPasswordHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "susess"})
}
func (api UserAPI) DeleteUserByIDHandler(context *gin.Context) {
	var user model.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.Println("error EditPasswordHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.UserRepository.DeleteUser(user.UserName)
	if err != nil {
		log.Println("error DeleteUserHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	context.JSON(http.StatusNoContent, gin.H{"message": "susess"})
}
func (api UserAPI) GETONEUserHandeler(context *gin.Context) {
	var user model.User
	err := context.ShouldBindJSON(&user)
	fmt.Println(user)
	onesubject, err:= api.UserRepository.GetUser(user.UserName)
	if err != nil {
		log.Println("error GETONEUserHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, onesubject)
}
func (api UserAPI) LoginHandler(context *gin.Context) {
	var user model.Login
	err := context.ShouldBindJSON(&user)	
	onesubject, err:= api.UserRepository.GetUser(user.UserName)
	if err != nil {
		log.Println("error LoginHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	for _, copy := range onesubject {
		if copy.UserName == user.UserName && copy.TPassword == user.TPassword {
			context.JSON(http.StatusOK, "Success")
		}
	}
	
	
}
func (api UserAPI) GETONEnameHandeler(context *gin.Context) {
	var user model.User
	username:= context.Param("username")
	fmt.Println(user)
	fmt.Println(username)
	oneuser, err2:= api.UserRepository.GetUser(username)
	if err2 != nil {
		log.Println("error GetUser", err2.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err2.Error()})
		return
	}
	context.JSON(http.StatusOK, oneuser)
}


func (api UserAPI) GETPictureUserHandeler(context *gin.Context) {
	var user model.User
	id:= context.Param("id")
	fmt.Println(id)
	pic, err:= api.UserRepository.Getbyid(id)
	fmt.Println(user.TPicture)
	if err != nil {
		log.Println("error GetUser", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	for _, copy := range pic {
		
			context.JSON(http.StatusOK, copy.TPicture)
		
	}
	
}



//QRCODE
// type QRCodeAPI struct {
// 	//QRCodeRepository repository.QRCodeRepository
// }
// var newQrs []models.CreateQr
// func CreateBarcode(c *gin.Context){
// 	var png []byte
// 	var p models.Qrcode
// 	defer c.Request.Body.Close()
// 	if err := c.ShouldBindJSON(&p); err != nil {
// 		c.JSON(http.StatusOK, "Can't creat class")
// 	}

// 	newQr := models.Qrcode{
// 		Time: p.Time,
// 		User: p.User,
// 		Pass: p.Pass,
// 	}

// 	fmt.Println(newQr)
// 	png, _ = qrcode.Encode(p.Time + p.User + p.Pass, qrcode.Medium, 256)
// 	sEnc := b64.StdEncoding.EncodeToString([]byte(png))
// 	fmt.Println("test = " + sEnc)
// 	sp := models.CreateQr{
// 		Id:     len(newQrs)+1,
// 		Qrcode: sEnc,
// 		Time: p.Time,
// 	}

// 	newQrs = append(newQrs , sp)
// 	c.JSON(http.StatusOK, newQrs)
// }


