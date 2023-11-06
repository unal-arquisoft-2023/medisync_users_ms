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
	// FindOne returns a patient given its id
	FindOne(ctx context.Context, id string) (*domain.Patient, UserRepositoryError)
	// Create creates a new patient and returns it
	Create(ctx context.Context, input PatientCreationInput) (*domain.Patient, UserRepositoryError)
	// Update updates a patient and returns the updated object
	Update(ctx context.Context, input PatientUpdateInput) (*domain.Patient, UserRepositoryError)
	// Suspends a patient and returns the updated object
	Suspend(ctx context.Context, id string) (*domain.Patient, UserRepositoryError)
	// Activates a patient and returns the updated object
	Activate(ctx context.Context, id string) (*domain.Patient, UserRepositoryError)
}
