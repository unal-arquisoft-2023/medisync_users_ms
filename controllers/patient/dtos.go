package patient

import (
	"medysinc_user_ms/controllers"
	"medysinc_user_ms/domain"
)

type PatientAffiliationRequest string

const (
	Private   PatientAffiliationRequest = PatientAffiliationRequest(domain.Private)
	Public    PatientAffiliationRequest = PatientAffiliationRequest(domain.Public)
	Insurance PatientAffiliationRequest = PatientAffiliationRequest(domain.Insurance)
)

type CreatePatientRequest struct {
	controllers.UserCreationRequest
	Affiliation string `json:"affiliation" validate:"required"`
}
