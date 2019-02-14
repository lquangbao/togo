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
	router.HandleFunc("/todo/{id}", getTodo).Methods("GET") // id is the title
	router.HandleFunc("/todo", createTodo).Methods("POST")
	router.HandleFunc("/todo", createTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	log.Println("Serve at http://localhost:" + port + "/todo")
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, router))
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todoChannel = make(chan entity.Todo)
	go db.Select(params["id"], todoChannel)
	json.NewEncoder(w).Encode(<-todoChannel)
	w.WriteHeader(http.StatusOK)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos := db.SelectAll()
	json.NewEncoder(w).Encode(todos)
	w.WriteHeader(http.StatusOK)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo entity.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	db.Insert(todo)
	json.NewEncoder(w).Encode(todo)
	w.WriteHeader(http.StatusOK)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	var todo entity.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	db.Update(todo)
	json.NewEncoder(w).Encode(todo)
	w.WriteHeader(http.StatusOK)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db.Remove(params["id"])
	w.WriteHeader(http.StatusOK)
}
