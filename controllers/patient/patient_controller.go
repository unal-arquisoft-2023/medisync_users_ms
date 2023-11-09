package patient

import (
	"medysinc_user_ms/controllers"
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
		UserCreationInput: users.UserCreationInput{
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
			DateOfBirth: req.DateOfBirth.Time(),
			Status:      domain.UserStatus(req.Status),
			CardID:      req.CardID,
		},
		Affiliation: domain.PatientAffiliation(req.Affiliation),
	}

	patient, err := pc.patRepo.Create(c.Request().Context(), input)

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	patRes := makePatientResponse(*patient)

	return c.JSON(http.StatusCreated, patRes)
}

func (pc *PatientController) GetPatient(c echo.Context, idReq controllers.UserIdRequest) error {
	patient, err := pc.patRepo.FindOne(c.Request().Context(), idReq.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	patRes := makePatientResponse(*patient)

	return c.JSON(http.StatusOK, patRes)
}

func (pc *PatientController) UpdatePatient(c echo.Context, req UpdatePatientRequest) error {
	input := users.PatientUpdateInput{
		UserUpdateInput: users.UserUpdateInput{
			ID: req.ID,
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
			DateOfBirth: req.DateOfBirth.Time(),
			Status:      domain.UserStatus(req.Status),
			CardID:      req.CardID,
		},
		Affiliation: domain.PatientAffiliation(req.Affiliation),
	}

	patient, err := pc.patRepo.Update(c.Request().Context(), input)

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	patRes := makePatientResponse(*patient)

	return c.JSON(http.StatusOK, patRes)
}

// Route to susped a patient
func (pc *PatientController) SuspendPatient(c echo.Context, idReq controllers.UserIdRequest) error {
	_, err := pc.patRepo.Suspend(c.Request().Context(), idReq.ID)

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// Route to activate a patient
func (pc *PatientController) ActivatePatient(c echo.Context, idReq controllers.UserIdRequest) error {
	_, err := pc.patRepo.Activate(c.Request().Context(), idReq.ID)

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// Route to get all patients
func (pc *PatientController) GetAllPatients(c echo.Context) error {
	patients, err := pc.patRepo.FindAll(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	patRes := make([]PatientResponse, 0)

	for _, patient := range patients {
		patRes = append(patRes, makePatientResponse(patient))
	}

	return c.JSON(http.StatusOK, patRes)
}
