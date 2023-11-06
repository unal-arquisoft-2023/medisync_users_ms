package controllers

import "medysinc_user_ms/domain"

type NameRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type LocationRequest struct {
	Country string `json:"country" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type UserStatusRequest string

const (
	Active    UserStatusRequest = UserStatusRequest(domain.Active)
	Suspended UserStatusRequest = UserStatusRequest(domain.Suspended)
	Insurance UserStatusRequest = UserStatusRequest(domain.Insurance)
)

type UserCreationRequest struct {
	Name             NameRequest       `json:"name" validate:"required"`
	Email            string            `json:"email" validate:"required,email"`
	Phone            string            `json:"phone" validate:"required"`
	Location         LocationRequest   `json:"location" validate:"required"`
	DateOfBirth      string            `json:"dateOfBirth" validate:"required"`
	RegistrationDate string            `json:"registrationDate" validate:"required"`
	Status           UserStatusRequest `json:"status" validate:"required"`
	CardID           string            `json:"cardId" validate:"required"`
}
