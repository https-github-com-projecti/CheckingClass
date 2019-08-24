package api

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	// "github.com/gorilla/mux"
	"net/http"
	// "gopkg.in/mgo.v2/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)


type user struct{
	// ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	tName string `json:"tName"`
	userName string `json:"userName"`
	tId string `json:"tId"`
	tEmail string `json:"tEmail"`
	tWorkPlace string `json:"tWorkPlace"`
	tPassword string `json:"tPassword"`
}

type users []user
// var sp []user

func Logins(w http.ResponseWriter,r *http.Request){
	fmt.Println("Welcom to func login")
}

func AddUserTeacher(w http.ResponseWriter, r *http.Request){
	fmt.Println("Welcom to add user")

	// vars := mux.Vars(r)
	// fmt.Println(vars)

	// var tNames, userNames, tIds, tEmails, tWorkPlaces, tPasswords  string
	// tNames = string(vars["tName"])
	// userNames =  string(vars["userName"])
	// tIds =  string(vars["tId"])
	// tEmails =  string(vars["tEmail"])
	// tWorkPlaces =  string(vars["tWorkPlace"])
	// tPasswords =  string(vars["tPassword"])

	// sp = []user{
	// 	user{tName : tNames,
	// 		userName : userNames,
	// 		tId : tIds,
	// 		tEmail : tEmails,
	// 		tWorkPlace : tWorkPlaces,
	// 		tPassword : tPasswords},
	// }
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// Unmarshal
	var msg user
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// var userCreate user
	// userCreate.tName = msg.tName
	// userCreate.tId = msg.tId
	// userCreate.userName = msg.userName
	// userCreate.tEmail = msg.tEmail
	// userCreate.tWorkPlace = msg.tWorkPlace
	// userCreate.tPassword = msg.tPassword
	fmt.Println(msg.tName)
	fmt.Println(msg.userName)
	fmt.Println(msg.tId)
	fmt.Println(msg.tEmail)
	fmt.Println(msg.tWorkPlace)
	fmt.Println(msg.tPassword)

}

// func AllUser (w http.ResponseWriter, r *http.Request){
// 	json.NewEncoder(w).Encode(users)
// }
