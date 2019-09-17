package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/narawichsaphimarn/backend/api"
	// 	"log"
	// 	"net/http"
)

func welcome(c *gin.Context) {
	c.JSON(200, "online")
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080/"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4200"
		},
		MaxAge: 12 * time.Hour,
	}))
	c1 := r.Group("/home")
	{
		c1.GET("/ping", welcome)
	}

	c2 := r.Group("/user")
	{
		c2.POST("/newUser", api.NewUser)
		c2.GET("/allUsers", api.AllUsers)
		c2.POST("/loginUser", api.UserLogin)
		c2.GET("/getId/:user", api.GetId)
		c2.GET("/getMyPic/:id", api.GetPicture)
	}

	c3 := r.Group("/class")
	{
		c3.POST("/newClass", api.CreatClass)
		c3.GET("/allClass", api.AllClass)
		c3.GET("/myClass/:user", api.MyClass)
		c3.GET("/selectClass/:id", api.OneClass)
	}

	c4 := r.Group("/qr")
	{
		c4.POST("/createqr", api.CreateBarcode)
		c4.GET("/allQr", api.AllQr)
		c4.GET("/myQr/:user", api.MyQr)
	}
	r.Run(":8080")
}
