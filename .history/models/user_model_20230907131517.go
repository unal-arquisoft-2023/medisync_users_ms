package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	Country string `json:"country,omitempty" validate:"required"`
	City    string `json:"city,omitempty" validate:"required"`
	Adress  string `json:"adress,omitempty" validate:"required"`
}

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required,email"`
	Phone    string             `json:"phone,omitempty" validate:"required"`
	Location Location           `json:"location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" validate:"required"`
}
