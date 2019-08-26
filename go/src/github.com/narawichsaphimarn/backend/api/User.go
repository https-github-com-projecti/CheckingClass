package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/backend/models"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var testUsers []models.User

func NewUser(c *gin.Context) {
	var p models.User
	if err := c.ShouldBindJSON(&p); err != nil {
		return
	}
	//fmt.Println(p)
	addUser := models.User{
		len(testUsers) + 1,
		p.TName,
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
	c.JSON(http.StatusOK, testUsers)
}
