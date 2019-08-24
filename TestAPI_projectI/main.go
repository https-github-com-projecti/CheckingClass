package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/narawichsaphimarn/TestAPI_projectI/api"
	
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticle(w http.ResponseWriter, r *http.Request){
	article := Articles{
		Article{Title:"Test API1",Desc:"Test Description1",Content:"Test content1"},
		Article{Title:"Test API2",Desc:"Test Description2",Content:"Test content2"},
	}
	fmt.Println("Endpoint Hit: All Article Endpoint")
	json.NewEncoder(w).Encode(article)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello new API")
}

func testPostArticle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello test post")
}

func handleRequests(){
	myrouter := mux.NewRouter().StrictSlash(true)

	myrouter.HandleFunc("/", homePage)
	myrouter.HandleFunc("/allArticle", allArticle).Methods("GET")
	myrouter.HandleFunc("/allArticle", testPostArticle).Methods("POST")

	myrouter.HandleFunc("/login", api.Logins).Methods("POST")
	myrouter.HandleFunc("/newUser", api.AddUserTeacher).Methods("POST")
	// myrouter.HandleFunc("/allUser", api.AllUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(myrouter)))
}

func main(){
	handleRequests()
}
