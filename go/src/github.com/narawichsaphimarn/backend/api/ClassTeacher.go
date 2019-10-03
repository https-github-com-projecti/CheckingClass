package api

import (
	"flag"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/backend/models"
	// "fmt"
)

var testClass []models.TClass

func CreatClass(c *gin.Context) {
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

func AllClass(c *gin.Context) {
	c.JSON(http.StatusOK, testClass)
	defer c.Request.Body.Close()
}

func MyClass(c *gin.Context) {
	user := c.Param("user")
	defer c.Request.Body.Close()
	var sp []models.TClass
	for _, copy := range testClass {
		if copy.TUser == user {
			sp = append(sp, copy)
		}
	}
	c.JSON(http.StatusOK, sp)
}

func OneClass(c *gin.Context) {
	flag.Parse()
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	defer c.Request.Body.Close()
	var copy []models.TClass
	for _, sp := range testClass {
		if sp.Id == i {
			copy = append(copy, sp)
		}
	}
	c.JSON(http.StatusOK, copy)
}

func RandPassClass() (i int) {
	rand.Intn(10000)
	rand.Seed(time.Now().UnixNano())
	classnum := rand.Intn(10000)
	num := classnum
	var count int = 0
	for num > 0 {
		num = num / 10
		count = count + 1
	}
	state := true
	for _, copy := range testClass {
		if classnum == copy.TClassPass {
			state = false
		}
	}
	if state == false {
		RandPassClass()
	}
	if count < 4 {
		RandPassClass()
	}
	i = classnum
	return
}
