package patient

import (
	"medysinc_user_ms/controllers"
	"medysinc_user_ms/domain"
)

type PatientAffiliationDTO string

const (
	Private   PatientAffiliationDTO = PatientAffiliationDTO(domain.Private)
	Public    PatientAffiliationDTO = PatientAffiliationDTO(domain.Public)
	Insurance PatientAffiliationDTO = PatientAffiliationDTO(domain.Insurance)
)

type CreatePatientRequest struct {
	controllers.UserCreationRequest `json:",inline"`
	Affiliation                     string `json:"affiliation" validate:"required"`
}

type UpdatePatientRequest struct {
	controllers.UserUpdateRequest `json:",inline"`
	Affiliation                   string `json:"affiliation" validate:"required"`
}

type PatientResponse struct {
	controllers.UserResponse `json:",inline"`
	Affiliation              PatientAffiliationDTO `json:"affiliation" validate:"required"`
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
			DateOfBirth:      p.DateOfBirth,
			RegistrationDate: p.RegistrationDate,
			Status:           controllers.UserStatusDTO(p.Status),
			CardID:           p.CardID,
		},
		Affiliation: PatientAffiliationDTO(p.Affiliation),
	}

}
