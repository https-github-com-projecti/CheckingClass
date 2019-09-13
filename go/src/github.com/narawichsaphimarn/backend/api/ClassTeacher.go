package api

import (
	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/backend/models"
	"net/http"
)

var testClass []models.TClass

func CreatClass(c *gin.Context){
	var p models.TClass
	defer c.Request.Body.Close()
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, "Can't creat class")
	}

	newClass := models.TClass{
		Id:                len(testClass) + 1,
		TClassName:        p.TClassName,
		TClassDescription: p.TClassDescription,
		TClassId:          p.TClassId,
		TUser:             p.TUser,
	}

	testClass = append(testClass, newClass)
	c.JSON(http.StatusOK, "Success")
}

func AllClass(c *gin.Context){
	c.JSON(http.StatusOK, testClass)
	defer c.Request.Body.Close()
}

func MyClass(c *gin.Context){
	user := c.Param("user")
	defer c.Request.Body.Close()
	var sp []models.TClass
	for _ , copy := range testClass{
		if(copy.TUser == user){
			sp = append(sp, copy)
		}
	}
	c.JSON(http.StatusOK, sp)
}

//func AddDataClass(){
//	addNewClass := models.TClass{
//		len(testClass) + 1,
//		"Project 1",
//		"สำหรับลงโปรเเจ็ค",
//		"523495",
//		"NS",
//	}
//	testClass = append(testClass, addNewClass)
//}
