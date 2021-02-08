package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers 1.1")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article-all", returnAllArticles)
	myRouter.HandleFunc("/article-query/{id}", returnSingleArticle)
	myRouter.HandleFunc("/article-creation", returnCreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/article-delete/{id}", returnDeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article-update", returnUpdateArticle).Methods("PUT")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func returnCreateNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	/*
		fmt.Println(r)
		fmt.Println(r.Body)
		fmt.Println(r.Header)
		fmt.Println(r.GetBody)
	*/
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func returnDeleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our
		// articles
		if article.Id == id {
			// updates our Articles array to remove the
			// article
			Articles = append(Articles[:index], Articles[index+1:]...)
			w.Write([]byte("deletion is successful"))
		}
	}
}

func returnUpdateArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var reqUpdate Article
	json.Unmarshal(reqBody, &reqUpdate)

	id := reqUpdate.Id

	for index, article := range Articles {
		// fmt.Println(Articles[index])
		if article.Id == id {
			Articles[index] = reqUpdate
			break
		}
	}

	//Articles = append(Articles, article)
	//json.NewEncoder(w).Encode(article)
	w.Write([]byte("updation is successful"))
}
