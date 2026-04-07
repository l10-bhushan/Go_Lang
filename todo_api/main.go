package main

import (
	"encoding/json"
	"example/todo_api/model"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func errorHandler(w http.ResponseWriter) {
	fmt.Fprintf(w, "Unexpected error occurred %d", http.StatusInternalServerError)
}

var globalTodos []model.Todo = []model.Todo{}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	result := model.Result{
		Status: true,
		Data:   "Server is healthy",
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		errorHandler(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func todoCreateHandler(w http.ResponseWriter, r *http.Request) {
	var data model.TodoRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		errorHandler(w)
		return
	}
	fmt.Println(data)
	id := uuid.New()
	newTodo := model.Todo{
		Id:        id.String(),
		Title:     data.Title,
		Status:    false,
		Timestamp: time.Now().GoString(),
	}
	globalTodos = append(globalTodos, newTodo)
	response := model.CreationSuccess{
		Status:  true,
		Message: "Todo successfully created",
		Data:    globalTodos,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		errorHandler(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)

}
func todoUpdateHandler(w http.ResponseWriter, r *http.Request) {

}
func todoDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var newTodo []model.Todo = []model.Todo{}
	for _, value := range globalTodos {
		if value.Id != id {
			newTodo = append(newTodo, value)
		}
	}
	globalTodos = newTodo
	response := model.CreationSuccess{
		Status:  true,
		Message: "Todo successfully created",
		Data:    globalTodos,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		errorHandler(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func todoGetHandler(w http.ResponseWriter, r *http.Request) {
	result := model.TodoResult{
		Status: true,
		Data:   globalTodos,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		errorHandler(w)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResult)
}

func main() {
	fmt.Println("CRUD API using go...")
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/todo/create", todoCreateHandler).Methods("POST")
	r.HandleFunc("/todo/update", todoUpdateHandler).Methods("PUT", "PATCH")
	r.HandleFunc("/todo/delete/{id}", todoDeleteHandler).Methods("DELETE")
	r.HandleFunc("/todo/get", todoGetHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
