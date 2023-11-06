package patient

import (
	"medysinc_user_ms/controllers"
	"medysinc_user_ms/domain"
	"medysinc_user_ms/resources/validation"

	"github.com/go-playground/validator"
)

type PatientAffiliationDTO string

const (
	Private   PatientAffiliationDTO = PatientAffiliationDTO(domain.Private)
	Public    PatientAffiliationDTO = PatientAffiliationDTO(domain.Public)
	Insurance PatientAffiliationDTO = PatientAffiliationDTO(domain.Insurance)
)

type CreatePatientRequest struct {
	controllers.UserCreationRequest `json:",inline"`
	Affiliation                     PatientAffiliationDTO `json:"affiliation" validate:"required,isAffiliation"`
}

type UpdatePatientRequest struct {
	controllers.UserUpdateRequest `json:",inline"`
	Affiliation                   PatientAffiliationDTO `json:"affiliation" validate:"required,isAffiliation"`
}

type PatientResponse struct {
	controllers.UserResponse `json:",inline"`
	Affiliation              PatientAffiliationDTO `json:"affiliation"`
}

func makePatientResponse(p domain.Patient) PatientResponse {
	return PatientResponse{
		UserResponse: controllers.UserResponse{
			ID: p.ID,
			Name: controllers.NameDTO{
				FirstName: p.Name.FirstName,
				LastName:  p.Name.LastName,
			},
			Email:            p.Email,
			Phone:            p.Phone,
			Location:         controllers.LocationDTO{Country: p.Location.Country, City: p.Location.City, Address: p.Location.Address},
			DateOfBirth:      controllers.CivilTime(p.DateOfBirth),
			RegistrationDate: p.RegistrationTimeStamp,
			Status:           controllers.UserStatusDTO(p.Status),
			CardID:           p.CardID,
		},
		Affiliation: PatientAffiliationDTO(p.Affiliation),
	}

}

func AddCustomDTOValidations(
	val *validation.MedisyncValidator,
) {
	val.Validator.RegisterValidation("isAffiliation", func(fl validator.FieldLevel) bool {
		affiliation := PatientAffiliationDTO(fl.Field().String())

		switch affiliation {
		case Public, Private, Insurance:
			return true
		default:
			return false
		}
	})
}
