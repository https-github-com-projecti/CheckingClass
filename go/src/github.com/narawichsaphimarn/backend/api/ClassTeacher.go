package api

import (
	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/backend/models"
	"net/http"
	"strconv"
	"flag"
	"math/rand"
	// "fmt"
)

var testClass []models.TClass

func CreatClass(c *gin.Context){
	var p models.TClass
	defer c.Request.Body.Close()
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, "Can't creat class")
	}
	num := RandPassClass()

	newClass := models.TClass{
		Id:                len(testClass) + 1,
		TClassName:        p.TClassName,
		TClassDescription: p.TClassDescription,
		TClassId:          p.TClassId,
		TUser:             p.TUser,
		TClassPass:        num,
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

func OneClass(c *gin.Context){
	flag.Parse()
	id := c.Param("id")
	i, _  := strconv.Atoi(id)
	defer c.Request.Body.Close()
	var copy []models.TClass
	for _, sp := range testClass {
		if (sp.Id == i ){
			copy = append(copy, sp)
		} 
	}
	c.JSON(http.StatusOK, copy)
}

func RandPassClass() (i int){
	classnum := rand.Intn(100000000)
	for _, copy := range testClass {
		if (classnum == copy.TClassPass) {
			RandPassClass()
		}
	}
	i = classnum
	return 
}



