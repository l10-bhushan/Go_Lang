package main

import (
	"example/practice-be/handler"
	"example/practice-be/repository"
	"example/practice-be/service"
	"fmt"
	"net/http"
)

func main() {

	repo := &repository.UserRepository{}
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	http.HandleFunc("/user", handler.GetUser)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
