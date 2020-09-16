package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct{
	Title string `json:"Title"`
	Desc string `json:"desc"`
}

var Articles []Article

func homePage(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "Welcome to our webpage")
	fmt.Println("Endpoint hit: Homepage")
}

func createNewArticle(writer http.ResponseWriter, request *http.Request){
	// Create a new article
	reqBody, _ := ioutil.ReadAll(request.Body)
	fmt.Fprintf(writer, string(reqBody))
///make the checks for the json object then push it to the array of articles or push it to the DB
}


func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000" , myRouter))
}


func returnAllArticles(writer http.ResponseWriter, request *http.Request){
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Desc"},
		Article{Title: "Hi", Desc: "Article Desc"},
	}
	json.NewEncoder(writer).Encode(Articles)
}

func main(){

	handleRequests()
}

