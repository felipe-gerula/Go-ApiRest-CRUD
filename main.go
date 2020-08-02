package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"

	"github.com/gorilla/mux"
)


type task struct{
	ID int `json:"ID"`
	Name string `json:"Name"`  
	Content string `json:"Content"`
}  

type allTask []task 

var tasks = allTask{
	{
		ID:1,
		Name:"Task One",
		Content:"Some Content",
	},
}

func getTasks(w http.ResponseWriter , r *http.Request){
	json.NewEncoder(w).Encode(tasks)
}

func indexRoute(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w,"welcome to my API")
}

func main(){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/",indexRoute)
	router.HandleFunc("/tasks",getTasks)

	log.Fatal(http.ListenAndServe(":3000",router)) 
}

