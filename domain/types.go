package domain

import "time"

type UserStatus string

const (
	Active    UserStatus = "ACTIVE"
	Suspended UserStatus = "SUSPENDED"
)

type Location struct {
	Country string
	City    string
	Address string
}

type Name struct {
	FirstName string
	LastName  string
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
	ID                    string
	Name                  Name
	Email                 string
	Phone                 string
	Location              Location
	DateOfBirth           time.Time
	RegistrationTimeStamp time.Time
	Status                UserStatus
	CardID                string
}

type Staff struct {
	Position string
	User
}

type Patient struct {
	Affiliation PatientAffiliation
	User
}

type Doctor struct {
	Specialty        DoctorSpecialty
	MedicalLicenseID string
	User
}
