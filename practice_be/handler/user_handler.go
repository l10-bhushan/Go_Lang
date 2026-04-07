package handler

import (
	"encoding/json"
	"example/practice-be/service"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	user := h.service.GetUser()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
