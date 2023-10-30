package users

import (
	"context"
	"medysinc_user_ms/domain"
)

type PatientCreationInput struct {
	Name             domain.Name
	Email            string
	Phone            string
	Location         domain.Location
	DateOfBirth      string
	RegistrationDate string
	Status           domain.UserStatus
	CardID           string
	Affiliation      domain.PatientAffiliation
}

type PatientUpdateInput struct {
	ID               string
	Name             domain.Name
	Email            string
	Phone            string
	Location         domain.Location
	DateOfBirth      string
	RegistrationDate string
	Status           domain.UserStatus
	CardID           string
	Affiliation      domain.PatientAffiliation
}

type PatientRepository interface {
	FindOne(ctx context.Context, id string) (*domain.Patient, error)
	Create(ctx context.Context, input PatientCreationInput) (*domain.Patient, error)
	Update(ctx context.Context, input PatientUpdateInput) (*domain.Patient, error)
	Delete(ctx context.Context, id string) (*domain.Patient, error)
}
