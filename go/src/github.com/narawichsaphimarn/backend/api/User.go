package api

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/backend/models"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var testUsers []models.User

//var pare models.User

func NewUser(c *gin.Context) {
	var p models.User
	defer c.Request.Body.Close()
	if err := c.ShouldBindJSON(&p); err != nil {
		return
	}
	//fmt.Println(p)
	addUser := models.User{
		len(testUsers) + 1,
		p.TFirstName,
		p.TLastName,
		p.UserName,
		p.TId,
		p.TEmail,
		p.TWorkPlace,
		p.TPassword,
		p.TPicture,
	}

	testUsers = append(testUsers, addUser)
	fmt.Println(testUsers)

}

func AllUsers(c *gin.Context) {
	defer c.Request.Body.Close()
	c.JSON(http.StatusOK, testUsers)
}

func UserLogin(c *gin.Context) {
	var Ui models.UserLogin
	defer c.Request.Body.Close()
	if c.ShouldBindJSON(&Ui) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(Ui.Username)
		log.Println(Ui.Password)
	} 
	for _, copy := range testUsers {
		if Ui.Username == copy.UserName && Ui.Password == copy.TPassword {
			c.JSON(http.StatusOK, copy)
			fmt.Println(copy)
		}
	}
	
	fmt.Println(testUsers)
	fmt.Println(Ui)
}
