package patient

import (
	"medysinc_user_ms/domain"
	"medysinc_user_ms/resources/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PatientController struct {
	patRepo users.PatientRepository
}

// Creates a new patient controller
func NewPatientController(patRepo users.PatientRepository) *PatientController {
	return &PatientController{patRepo}
}

// Creates a new patient
func (pc *PatientController) CreatePatient(c echo.Context, req CreatePatientRequest) error {

	input := users.PatientCreationInput{
		Name: domain.Name{
			FirstName: req.Name.FirstName,
			LastName:  req.Name.LastName,
		},
		Email: req.Email,
		Phone: req.Phone,
		Location: domain.Location{
			Country: req.Location.Country,
			City:    req.Location.City,
			Address: req.Location.Address,
		},
		DateOfBirth:      req.DateOfBirth,
		RegistrationDate: req.RegistrationDate,
		Status:           domain.UserStatus(req.Status),
		CardID:           req.CardID,
		Affiliation:      domain.PatientAffiliation(req.Affiliation),
	}

	patient, err := pc.patRepo.Create(c.Request().Context(), input)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, patient)
}
