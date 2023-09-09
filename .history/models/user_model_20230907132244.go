package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStatus string

const (
	Active    UserStatus = "ACTIVE"
	Suspended UserStatus = "SUSPENDED"
)

type Location struct {
	Country string `json:"country,omitempty" validate:"required"`
	City    string `json:"city,omitempty" validate:"required"`
	Adress  string `json:"adress,omitempty" validate:"required"`
}

type Name struct {
	FirstName string `json:"firstName,omitempty" validate:"required"`
	LastName  string `json:"lastName,omitempty" validate:"required"`
}

type User struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        Name               `json:"name,omitempty" validate:"required"`
	Email       string             `json:"email,omitempty" validate:"required,email"`
	Phone       string             `json:"phone,omitempty" validate:"required"`
	Location    Location           `json:"location,omitempty" validate:"required"`
	Title       string             `json:"title,omitempty" validate:"required"`
	DateOfBirth time.Time          `json:"dateOfBirth,omitempty" validate:"required"`
	Status      UserStatus         `json:"status,omitempty" validate:"required"`
	DNI         string             `json:"dni,omitempty" validate:"required"`
}
