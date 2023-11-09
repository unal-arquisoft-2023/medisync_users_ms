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
	// FindAll returns all patients
	FindAll(ctx context.Context) ([]domain.Patient, UserRepositoryError)
}

type DoctorCreationInput struct {
	UserCreationInput
	Specialty        domain.DoctorSpecialty
	MedicalLicenseID string
}

type DoctorUpdateInput struct {
	UserUpdateInput
	Specialty        domain.DoctorSpecialty
	MedicalLicenseID string
}

type DoctorRepository interface {
	// FindOne returns a doctor given its id
	FindOne(ctx context.Context, id string) (*domain.Doctor, UserRepositoryError)
	// Create creates a new Doctor and returns it
	Create(ctx context.Context, input DoctorCreationInput) (*domain.Doctor, UserRepositoryError)
	// Update updates a Doctor and returns the updated object
	Update(ctx context.Context, input DoctorUpdateInput) (*domain.Doctor, UserRepositoryError)
	// Suspends a Doctor and returns the updated object
	Suspend(ctx context.Context, id string) (*domain.Doctor, UserRepositoryError)
	// Activates a Doctor and returns the updated object
	Activate(ctx context.Context, id string) (*domain.Doctor, UserRepositoryError)
	// FindAll returns all Staffs
	FindAll(ctx context.Context) ([]domain.Doctor, UserRepositoryError)
	// Find all doctors by specialty
	FindBySpecialty(ctx context.Context, specialty domain.DoctorSpecialty) ([]domain.Doctor, UserRepositoryError)
}
type StaffRepository interface {
	// FindOne returns a patient given its id
	FindOne(ctx context.Context, id string) (*domain.Staff, UserRepositoryError)
	// Create creates a new Staff and returns it
	Create(ctx context.Context, input UserCreationInput) (*domain.Staff, UserRepositoryError)
	// Update updates a Staff and returns the updated object
	Update(ctx context.Context, input UserUpdateInput) (*domain.Staff, UserRepositoryError)
	// Suspends a Staff and returns the updated object
	Suspend(ctx context.Context, id string) (*domain.Staff, UserRepositoryError)
	// Activates a Staff and returns the updated object
	Activate(ctx context.Context, id string) (*domain.Staff, UserRepositoryError)
	// FindAll returns all Staffs
	FindAll(ctx context.Context) ([]domain.Staff, UserRepositoryError)
}
