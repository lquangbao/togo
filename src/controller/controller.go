package controller

import (
	"db"
	"encoding/json"
	"entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Serve - call this to start the server
func Serve(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/todo", getTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todo", createTodo).Methods("POST")
	router.HandleFunc("/todo", deleteTodo).Methods("DELETE")
	log.Println("Serve at http://localhost:" + port + "/todo")
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, router))
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todoChannel = make(chan entity.Todo)
	go db.Select(params["id"], todoChannel)
	json.NewEncoder(w).Encode(<-todoChannel)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos := db.SelectAll()
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {

}

func deleteTodo(w http.ResponseWriter, r *http.Request) {

}
