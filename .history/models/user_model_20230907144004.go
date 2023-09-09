package models

import (
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

type PatientAffiliation string

const (
	Private   PatientAffiliation = "PRIVATE"
	Public    PatientAffiliation = "PUBLIC"
	Insurance PatientAffiliation = "INSURANCE"
)

type DoctorSpecialty string

const (
	GeneralMedicine  DoctorSpecialty = "General Medicine"
	Pediatrics       DoctorSpecialty = "Pediatrics"
	Cardiology       DoctorSpecialty = "Cardiology"
	Orthopedics      DoctorSpecialty = "Orthopedics"
	Dermatology      DoctorSpecialty = "Dermatology"
	Gastroenterology DoctorSpecialty = "Gastroenterology"
	Neurology        DoctorSpecialty = "Neurology"
	Ophthalmology    DoctorSpecialty = "Ophthalmology"
	Oncology         DoctorSpecialty = "Oncology"
	Otolaryngology   DoctorSpecialty = "Otolaryngology"
	Urology          DoctorSpecialty = "Urology"
	Psychiatry       DoctorSpecialty = "Psychiatry"
	Obstetrics       DoctorSpecialty = "Obstetrics"
	Gynecology       DoctorSpecialty = "Gynecology"
	Anesthesiology   DoctorSpecialty = "Anesthesiology"
	Radiology        DoctorSpecialty = "Radiology"
	Pathology        DoctorSpecialty = "Pathology"
	Emergency        DoctorSpecialty = "Emergency"
	FamilyMedicine   DoctorSpecialty = "Family Medicine"
	InternalMedicine DoctorSpecialty = "Internal Medicine"
	Surgery          DoctorSpecialty = "Surgery"
	Other            DoctorSpecialty = "Other"
)

type User struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	Name         Name               `json:"name,omitempty" validate:"required"`
	Email        string             `json:"email,omitempty" validate:"required,email"`
	Phone        string             `json:"phone,omitempty" validate:"required"`
	Location     Location           `json:"location,omitempty" validate:"required"`
	Title        string             `json:"title,omitempty" validate:"required"`
	DateOfBirth  string             `json:"dateOfBirth,omitempty" validate:"required"`
	RegisterDate string             `json:"registerDate,omitempty" validate:"required"`
	Status       UserStatus         `json:"status,omitempty" validate:"required"`
	DNI          string             `json:"dni,omitempty" validate:"required"`
}

type Patient struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"userId,omitempty" validate:"required"` // Relación con User
	Affiliation PatientAffiliation `json:"affiliation,omitempty" validate:"required"`
}

type Doctor struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId           primitive.ObjectID `json:"userId,omitempty" validate:"required"` // Relación con User
	Specialty        DoctorSpecialty    `json:"specialty,omitempty" validate:"required"`
	MedicalLicenseID string             `json:"medicalLicenseID,omitempty" validate:"required"`
}
