package doctor

import (
	"medysinc_user_ms/controllers"
	"medysinc_user_ms/domain"
	"medysinc_user_ms/resources/validation"

	"github.com/go-playground/validator"
)

type DoctorSpecialtyDTO string

const (
	GeneralMedicine  DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.GeneralMedicine)
	Pediatrics       DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Pediatrics)
	Cardiology       DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Cardiology)
	Orthopedics      DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Orthopedics)
	Dermatology      DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Dermatology)
	Gastroenterology DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Gastroenterology)
	Neurology        DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Neurology)
	Ophthalmology    DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Ophthalmology)
	Oncology         DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Oncology)
	Otolaryngology   DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Otolaryngology)
	Urology          DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Urology)
	Psychiatry       DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Psychiatry)
	Obstetrics       DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Obstetrics)
	Gynecology       DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Gynecology)
	Anesthesiology   DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Anesthesiology)
	Radiology        DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Radiology)
	Pathology        DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Pathology)
	Emergency        DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Emergency)
	FamilyMedicine   DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.FamilyMedicine)
	InternalMedicine DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.InternalMedicine)
	Surgery          DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Surgery)
	Other            DoctorSpecialtyDTO = DoctorSpecialtyDTO(domain.Other)
)

type CreateDoctorRequest struct {
	controllers.UserCreationRequest `json:",inline"`
	Specialty                       DoctorSpecialtyDTO `json:"specialty" validate:"required,isSpecialty"`
	MedicalLicenseID                string             `json:"medicalLicenseID" validate:"required"`
}

type UpdateDoctorRequest struct {
	controllers.UserUpdateRequest `json:",inline"`
	Specialty                     DoctorSpecialtyDTO `json:"specialty" validate:"required,isSpecialty"`
	MedicalLicenseID              string             `json:"medicalLicenseID" validate:"required"`
}

type DoctorResponse struct {
	controllers.UserResponse `json:",inline"`
	Specialty                DoctorSpecialtyDTO `json:"specialty"`
	MedicalLicenseID         string             `json:"medicalLicenseID" validate:"required"`
}

func makeDoctorResponse(d domain.Doctor) DoctorResponse {
	return DoctorResponse{
		UserResponse: controllers.UserResponse{
			ID: d.ID,
			Name: controllers.NameDTO{
				FirstName: d.Name.FirstName,
				LastName:  d.Name.LastName,
			},
			Email:            d.Email,
			Phone:            d.Phone,
			Location:         controllers.LocationDTO{Country: d.Location.Country, City: d.Location.City, Address: d.Location.Address},
			DateOfBirth:      controllers.CivilTime(d.DateOfBirth),
			RegistrationDate: d.RegistrationTimeStamp,
			Status:           controllers.UserStatusDTO(d.Status),
			CardID:           d.CardID,
		},
		Specialty:        DoctorSpecialtyDTO(d.Specialty),
		MedicalLicenseID: d.MedicalLicenseID,
	}

}

func AddCustomDTOValidations(
	val *validation.MedisyncValidator,
) {
	val.Validator.RegisterValidation("isSpecialty", func(fl validator.FieldLevel) bool {
		specialty := DoctorSpecialtyDTO(fl.Field().String())

		switch specialty {
		case GeneralMedicine, Pediatrics, Cardiology, Orthopedics, Dermatology, Gastroenterology, Neurology, Ophthalmology, Oncology, Otolaryngology, Urology, Psychiatry, Obstetrics, Gynecology, Anesthesiology, Radiology, Pathology, Emergency, FamilyMedicine, InternalMedicine, Surgery, Other:
			return true
		default:
			return false
		}
	})
}
