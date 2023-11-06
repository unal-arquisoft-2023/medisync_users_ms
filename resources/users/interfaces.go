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
	FindOne(ctx context.Context, id string) (*domain.Patient, UserRepositoryError)
	Create(ctx context.Context, input PatientCreationInput) (*domain.Patient, UserRepositoryError)
	Update(ctx context.Context, input PatientUpdateInput) (*domain.Patient, UserRepositoryError)
	Suspend(ctx context.Context, id string) (*domain.Patient, UserRepositoryError)
	Activate(ctx context.Context, id string) (*domain.Patient, UserRepositoryError)
}
