package users

import (
	"context"
	"medysinc_user_ms/domain"
	"time"
)

type UserCreationInput struct {
	Name        domain.Name
	Email       string
	Phone       string
	Location    domain.Location
	DateOfBirth time.Time
	Status      domain.UserStatus
	CardID      string
}

type PatientCreationInput struct {
	UserCreationInput
	Affiliation domain.PatientAffiliation
}

type UserUpdateInput struct {
	ID          string
	Name        domain.Name
	Email       string
	Phone       string
	Location    domain.Location
	DateOfBirth time.Time
	Status      domain.UserStatus
	CardID      string
}

type PatientUpdateInput struct {
	UserUpdateInput
	Affiliation domain.PatientAffiliation
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
