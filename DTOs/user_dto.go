package dtos

import (
	m "medysinc_user_ms/models"
)

type CreatePatientRequest struct {
	Name             m.Name               `json:"name,omitempty" validate:"required"`
	Email            string               `json:"email,omitempty" validate:"required,email"`
	Phone            string               `json:"phone,omitempty" validate:"required"`
	Location         m.Location           `json:"location,omitempty" validate:"required"`
	DateOfBirth      string               `json:"dateOfBirth,omitempty" validate:"required"`
	RegistrationDate string               `json:"registrationDate,omitempty" validate:"required"`
	Status           m.UserStatus         `json:"status,omitempty" validate:"required"`
	CardId           string               `json:"cardId,omitempty" validate:"required"`
	Affiliation      m.PatientAffiliation `json:"affiliation,omitempty" validate:"required"`
}

type PatientResponse struct {
	Id               string               `json:"id,omitempty"`
	Name             m.Name               `json:"name,omitempty" validate:"required"`
	Email            string               `json:"email,omitempty" validate:"required,email"`
	Phone            string               `json:"phone,omitempty" validate:"required"`
	Location         m.Location           `json:"location,omitempty" validate:"required"`
	DateOfBirth      string               `json:"dateOfBirth,omitempty" validate:"required"`
	RegistrationDate string               `json:"registrationDate,omitempty" validate:"required"`
	Status           m.UserStatus         `json:"status,omitempty" validate:"required"`
	CardId           string               `json:"cardId,omitempty" validate:"required"`
	Affiliation      m.PatientAffiliation `json:"affiliation,omitempty" validate:"required"`
}

type CreateDoctorRequest struct {
	Name             m.Name            `json:"name,omitempty" validate:"required"`
	Email            string            `json:"email,omitempty" validate:"required,email"`
	Phone            string            `json:"phone,omitempty" validate:"required"`
	Location         m.Location        `json:"location,omitempty" validate:"required"`
	DateOfBirth      string            `json:"dateOfBirth,omitempty" validate:"required"`
	RegistrationDate string            `json:"registrationDate,omitempty" validate:"required"`
	Status           m.UserStatus      `json:"status,omitempty" validate:"required"`
	CardId           string            `json:"cardId,omitempty" validate:"required"`
	Specialty        m.DoctorSpecialty `json:"specialty,omitempty" validate:"required"`
	MedicalLicenseID string            `json:"medicalLicenseID,omitempty" validate:"required"`
}

type DoctorResponse struct {
	Id               string            `json:"id,omitempty"`
	Name             m.Name            `json:"name,omitempty" validate:"required"`
	Email            string            `json:"email,omitempty" validate:"required,email"`
	Phone            string            `json:"phone,omitempty" validate:"required"`
	Location         m.Location        `json:"location,omitempty" validate:"required"`
	DateOfBirth      string            `json:"dateOfBirth,omitempty" validate:"required"`
	RegistrationDate string            `json:"registrationDate,omitempty" validate:"required"`
	Status           m.UserStatus      `json:"status,omitempty" validate:"required"`
	CardId           string            `json:"cardId,omitempty" validate:"required"`
	Specialty        m.DoctorSpecialty `json:"specialty,omitempty" validate:"required"`
	MedicalLicenseID string            `json:"medicalLicenseID,omitempty" validate:"required"`
}

type CreateStaffRequest struct {
	Name             m.Name       `json:"name,omitempty" validate:"required"`
	Email            string       `json:"email,omitempty" validate:"required,email"`
	Phone            string       `json:"phone,omitempty" validate:"required"`
	Location         m.Location   `json:"location,omitempty" validate:"required"`
	DateOfBirth      string       `json:"dateOfBirth,omitempty" validate:"required"`
	RegistrationDate string       `json:"registrationDate,omitempty" validate:"required"`
	Status           m.UserStatus `json:"status,omitempty" validate:"required"`
	CardId           string       `json:"cardId,omitempty" validate:"required"`
	Position         string       `json:"position,omitempty" validate:"required"`
}

type StaffResponse struct {
	Id               string       `json:"id,omitempty"`
	Name             m.Name       `json:"name,omitempty" validate:"required"`
	Email            string       `json:"email,omitempty" validate:"required,email"`
	Phone            string       `json:"phone,omitempty" validate:"required"`
	Location         m.Location   `json:"location,omitempty" validate:"required"`
	Title            string       `json:"title,omitempty" validate:"required"`
	DateOfBirth      string       `json:"dateOfBirth,omitempty" validate:"required"`
	RegistrationDate string       `json:"registrationDate,omitempty" validate:"required"`
	Status           m.UserStatus `json:"status,omitempty" validate:"required"`
	CardId           string       `json:"cardId,omitempty" validate:"required"`
	Position         string       `json:"position,omitempty" validate:"required"`
}
