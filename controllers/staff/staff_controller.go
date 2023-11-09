package staff

import (
	"medysinc_user_ms/controllers"
	"medysinc_user_ms/domain"
	"medysinc_user_ms/resources/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StaffController struct {
	staffRepo users.StaffRepository
}

// Creates a new Staff controller
func NewStaffController(staffRepo users.StaffRepository) *StaffController {
	return &StaffController{staffRepo}
}

// Creates a new Staff
func (sc *StaffController) CreateStaff(c echo.Context, req CreateStaffRequest) error {

	input := users.UserCreationInput{

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
	}

	staff, err := sc.staffRepo.Create(c.Request().Context(), input)

	if err != nil {

		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	staffRes := makeStaffResponse(*staff)

	return c.JSON(http.StatusCreated, staffRes)
}

func (sc *StaffController) GetStaff(c echo.Context, idReq controllers.UserIdRequest) error {
	staff, err := sc.staffRepo.FindOne(c.Request().Context(), idReq.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	staffRes := makeStaffResponse(*staff)

	return c.JSON(http.StatusOK, staffRes)
}

func (sc *StaffController) UpdateStaff(c echo.Context, req UpdateStaffRequest) error {
	input := users.UserUpdateInput{

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
	}

	staff, err := sc.staffRepo.Update(c.Request().Context(), input)

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	staffRes := makeStaffResponse(*staff)

	return c.JSON(http.StatusOK, staffRes)
}

// Route to susped a staff
func (sc *StaffController) SuspendStaff(c echo.Context, idReq controllers.UserIdRequest) error {
	_, err := sc.staffRepo.Suspend(c.Request().Context(), idReq.ID)

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// Route to activate a staff
func (sc *StaffController) ActivateStaff(c echo.Context, idReq controllers.UserIdRequest) error {
	_, err := sc.staffRepo.Activate(c.Request().Context(), idReq.ID)

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// Route to get all Staffs
func (sc *StaffController) GetAllStaffs(c echo.Context) error {
	staffs, err := sc.staffRepo.FindAll(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(err.HttpStatusCode(), err.Error())
	}

	staffRes := make([]StaffResponse, 0)

	for _, staff := range staffs {
		staffRes = append(staffRes, makeStaffResponse(staff))
	}

	return c.JSON(http.StatusOK, staffRes)
}
