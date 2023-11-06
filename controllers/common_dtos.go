package controllers

import "medysinc_user_ms/domain"

type NameDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type LocationDTO struct {
	Country string `json:"country" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type UserStatusDTO string

const (
	Active    UserStatusDTO = UserStatusDTO(domain.Active)
	Suspended UserStatusDTO = UserStatusDTO(domain.Suspended)
	Insurance UserStatusDTO = UserStatusDTO(domain.Insurance)
)

type UserCreationRequest struct {
	Name             NameDTO       `json:"name" validate:"required"`
	Email            string        `json:"email" validate:"required,email"`
	Phone            string        `json:"phone" validate:"required"`
	Location         LocationDTO   `json:"location" validate:"required"`
	DateOfBirth      string        `json:"dateOfBirth" validate:"required"`
	RegistrationDate string        `json:"registrationDate" validate:"required"`
	Status           UserStatusDTO `json:"status" validate:"required"`
	CardID           string        `json:"cardId" validate:"required"`
}

type UserUpdateRequest struct {
	ID               string        `json:"id" validate:"required"`
	Name             NameDTO       `json:"name" validate:"required"`
	Email            string        `json:"email" validate:"required,email"`
	Phone            string        `json:"phone" validate:"required"`
	Location         LocationDTO   `json:"location" validate:"required"`
	DateOfBirth      string        `json:"dateOfBirth" validate:"required"`
	RegistrationDate string        `json:"registrationDate" validate:"required"`
	Status           UserStatusDTO `json:"status" validate:"required"`
	CardID           string        `json:"cardId" validate:"required"`
}

type UserResponse struct {
	ID               string        `json:"id"`
	Name             NameDTO       `json:"name"`
	Email            string        `json:"email"`
	Phone            string        `json:"phone"`
	Location         LocationDTO   `json:"location"`
	DateOfBirth      string        `json:"dateOfBirth"`
	RegistrationDate string        `json:"registrationDate"`
	Status           UserStatusDTO `json:"status"`
	CardID           string        `json:"cardId"`
}