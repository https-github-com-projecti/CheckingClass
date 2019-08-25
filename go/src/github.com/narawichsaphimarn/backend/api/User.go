package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/backend/models"
	"net/http"
)

func NewUser(c *gin.Context){
	var p models.User
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println(p.ID)
	fmt.Println(p.TName)
	fmt.Println(p.UserName)
	fmt.Println(p.TId)
	fmt.Println(p.TEmail)
	fmt.Println(p.TWorkPlace)
	fmt.Println(p.TPassword)

	fmt.Println(p)
	//c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func AllUsers(c *gin.Context){
	var p models.User
	c.JSON(http.StatusOK, p)
}
