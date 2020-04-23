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

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	fmt.Println(Articles)
	json.NewEncoder(w).Encode(Articles)
}
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	// fmt.Fprintf(w, "key:"+key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	reqbody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqbody, &article)
	// json.NewDecoder(r.Body).Decode(&article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
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
		}
	}

}
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {

		if article.Id == id {

			reqbody, _ := ioutil.ReadAll(r.Body)
			var newarticle Article
			json.Unmarshal(reqbody, &newarticle)
			Articles[index] = newarticle
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article/create", CreateArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", UpdateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", DeleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println(Articles)
	handleRequests()
}
