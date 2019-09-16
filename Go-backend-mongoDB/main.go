package main

import (
	"time"
	"github.com/gin-contrib/cors"
	"Go-backend-mongoDB/route"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

const (
	mongoDBEnPint  = "mongodb://localhost:27017"
	portWebServie = ":8080"
)

func main() {
	connectionDB, err := mgo.Dial(mongoDBEnPint)
	if err != nil {
		log.Panic("Can no connect Database", err.Error())
	}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
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
	route.NewRouteUser(router, connectionDB)
	router.Run(portWebServie)
}