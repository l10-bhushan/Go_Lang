package service

import (
	"example/practice-be/model"
	"example/practice-be/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUser() model.User {
	// business logic lives here
	return s.repo.GetUser()
}
