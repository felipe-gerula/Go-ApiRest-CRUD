package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"

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
	w.Header().Set("Content-Type" , "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTasks(w http.ResponseWriter , r *http.Request){
	var newTask task
	reqBody , err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w,"insert a valid task")
	}

	json.Unmarshal(reqBody , &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks,newTask)
 
	w.Header().Set("Content-Type" , "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func indexRoute(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w,"welcome to my API")
}

func main(){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/",indexRoute)
	router.HandleFunc("/tasks",getTasks).Methods("GET")
	router.HandleFunc("/tasks",createTasks).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000",router)) 
}

