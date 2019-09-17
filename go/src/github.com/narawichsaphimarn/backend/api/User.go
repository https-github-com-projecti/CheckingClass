package api

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/backend/models"
	"net/http"
	"strconv"
	"flag"
)

var testUsers []models.User

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
	if c.ShouldBindJSON(&Ui) != nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(Ui.Username)
		log.Println(Ui.Password)
	}
	for _, copy := range testUsers {
		if Ui.Username == copy.UserName && Ui.Password == copy.TPassword {
			c.JSON(http.StatusOK, "Success")
		}
	}
}

func GetId(c *gin.Context) {
	user := c.Param("user")
	defer c.Request.Body.Close()
	// var sp []models.User
	for _, copy := range testUsers {
		if copy.UserName == user {
			// sp = append(sp, copy)
			c.JSON(http.StatusOK, copy.ID)
		}
	}
}

func GetPicture(c *gin.Context){
	flag.Parse()
	id := c.Param("id");
	i, _  := strconv.Atoi(id)
	defer c.Request.Body.Close()
	for _, mypic := range testUsers {
		if (mypic.ID == i){
			c.JSON(http.StatusOK, mypic.TPicture)
			// fmt.Println("MyPic = " + mypic.TPicture)
		}
	}
}
