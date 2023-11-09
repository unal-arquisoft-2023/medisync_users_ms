package doctor

import (
	"medysinc_user_ms/controllers"
	"medysinc_user_ms/domain"
	"medysinc_user_ms/resources/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DoctorController struct {
	docRepo users.DoctorRepository
}

// Creates a new doctor controller
func NewDoctorController(docRepo users.DoctorRepository) *DoctorController {
	return &DoctorController{docRepo}
}

// Creates a new doctor
func (dc *DoctorController) CreateDoctor(c echo.Context, req CreateDoctorRequest) error {

	input := users.DoctorCreationInput{
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
		Specialty:        domain.DoctorSpecialty(req.Specialty),
		MedicalLicenseID: req.MedicalLicenseID,
	}

	doctor, err := dc.docRepo.Create(c.Request().Context(), input)

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	docRes := makeDoctorResponse(*doctor)

	return c.JSON(http.StatusCreated, docRes)
}

func (dc *DoctorController) GetDoctor(c echo.Context, idReq controllers.UserIdRequest) error {
	doctor, err := dc.docRepo.FindOne(c.Request().Context(), idReq.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	docRes := makeDoctorResponse(*doctor)

	return c.JSON(http.StatusOK, docRes)
}

func (dc *DoctorController) UpdateDoctor(c echo.Context, req UpdateDoctorRequest) error {
	input := users.DoctorUpdateInput{
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
		Specialty:        domain.DoctorSpecialty(req.Specialty),
		MedicalLicenseID: req.MedicalLicenseID,
	}

	doctor, err := dc.docRepo.Update(c.Request().Context(), input)

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	docRes := makeDoctorResponse(*doctor)

	return c.JSON(http.StatusOK, docRes)
}

// Route to susped a doctor
func (dc *DoctorController) SuspendDoctor(c echo.Context, idReq controllers.UserIdRequest) error {
	_, err := dc.docRepo.Suspend(c.Request().Context(), idReq.ID)

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// Route to activate a doctor
func (dc *DoctorController) ActivateDoctor(c echo.Context, idReq controllers.UserIdRequest) error {
	_, err := dc.docRepo.Activate(c.Request().Context(), idReq.ID)

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// Route to get all doctors
func (dc *DoctorController) GetAllDoctors(c echo.Context) error {
	doctors, err := dc.docRepo.FindAll(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	docRes := make([]DoctorResponse, 0)

	for _, doctor := range doctors {
		docRes = append(docRes, makeDoctorResponse(doctor))
	}

	return c.JSON(http.StatusOK, docRes)
}

// Route to get all doctors by specialty
func (dc *DoctorController) GetAllDoctorsBySpecialty(c echo.Context, specialtyReq controllers.SpecialtyRequest) error {
	doctors, err := dc.docRepo.FindBySpecialty(c.Request().Context(), domain.DoctorSpecialty(specialtyReq.Specialty))

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	docRes := make([]DoctorResponse, 0)

	for _, doctor := range doctors {
		docRes = append(docRes, makeDoctorResponse(doctor))
	}

	return c.JSON(http.StatusOK, docRes)
}
