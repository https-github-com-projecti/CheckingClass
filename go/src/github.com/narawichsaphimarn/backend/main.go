package main

import (
	"github.com/narawichsaphimarn/backend/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
	// 	"log"
	// 	"net/http"
)

func welcome (c *gin.Context){
	c.JSON(200, gin.H{
		"massage" : "Hello API",

	})
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080/"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATH"},
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
	}
	r.Run(":8080")
}
