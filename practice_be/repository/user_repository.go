package repository

import "example/practice-be/model"

type UserRepository struct{}

func (r *UserRepository) GetUser() model.User {
	return model.User{
		ID:      1,
		Name:    "Bhushan",
		Age:     28,
		Addr:    "Nashik, Maharashtra",
		Country: "India",
	}
}
