package controllers

import (
	"medysinc_user_ms/domain"
	"medysinc_user_ms/resources/validation"
	"strings"
	"time"

	"github.com/go-playground/validator"
)

type NameDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type LocationDTO struct {
	Country string `json:"country" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type UserStatusDTO string

const (
	Active    UserStatusDTO = UserStatusDTO(domain.Active)
	Suspended UserStatusDTO = UserStatusDTO(domain.Suspended)
)

type UserIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type UserCreationRequest struct {
	Name        NameDTO       `json:"name" validate:"required"`
	Email       string        `json:"email" validate:"required,email"`
	Phone       string        `json:"phone" validate:"required"`
	Location    LocationDTO   `json:"location" validate:"required"`
	DateOfBirth CivilTime     `json:"dateOfBirth" validate:"required"`
	Status      UserStatusDTO `json:"status" validate:"required,isUserStatus"`
	CardID      string        `json:"cardId" validate:"required"`
}

type UserUpdateRequest struct {
	ID          string        `param:"id" validate:"required"`
	Name        NameDTO       `json:"name" validate:"required"`
	Email       string        `json:"email" validate:"required,email"`
	Phone       string        `json:"phone" validate:"required"`
	Location    LocationDTO   `json:"location" validate:"required"`
	DateOfBirth CivilTime     `json:"dateOfBirth" validate:"required"`
	Status      UserStatusDTO `json:"status" validate:"required,isUserStatus"`
	CardID      string        `json:"cardId" validate:"required"`
}

type UserResponse struct {
	ID               string        `json:"id"`
	Name             NameDTO       `json:"name"`
	Email            string        `json:"email"`
	Phone            string        `json:"phone"`
	Location         LocationDTO   `json:"location"`
	DateOfBirth      CivilTime     `json:"dateOfBirth"`
	RegistrationDate time.Time     `json:"registrationDate"`
	Status           UserStatusDTO `json:"status"`
	CardID           string        `json:"cardId"`
}

// Why is this necesary?
// see: https://romangaranin.net/posts/2021-02-19-json-time-and-golang/
type CivilTime time.Time

func (c *CivilTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*c = CivilTime(t) //set result using the pointer
	return nil
}

func (c CivilTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("2006-01-02") + `"`), nil
}

func (c CivilTime) Time() time.Time {
	return time.Time(c)
}

func AddCustomDTOValidations(
	val *validation.MedisyncValidator,
) {

	val.Validator.RegisterValidation("isUserStatus", func(fl validator.FieldLevel) bool {
		status := UserStatusDTO(fl.Field().String())
		switch status {
		case Active, Suspended:
			return true
		default:
			return false
		}
	})

}
