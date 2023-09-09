package controllers

import (
	"context"
	"echo-mongo-api/configs"
	"medysinc_user_ms/models"
	"medysinc_user_ms/responses"
	"net/http"
	"os/user"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreatePatient(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var patient models.Patient
	defer cancel()

	if err := c.Bind(&patient); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if err := validate.Struct(patient); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	user := models.User{
		Id:			  primitive.NewObjectID(),
		Name:         patient.Name,
		Email:        patient.Email,
		Phone:        patient.Phone,
		Location:     patient.Location,
		Title:        patient.Title,
		DateOfBirth:  patient.DateOfBirth,
		RegisterDate: time.Now(),
		Status:       models.Active,
		DNI: 		  patient.DNI,
	}
	Name         Name               `json:"name,omitempty" validate:"required"`
	Email        string             `json:"email,omitempty" validate:"required,email"`
	Phone        string             `json:"phone,omitempty" validate:"required"`
	Location     Location           `json:"location,omitempty" validate:"required"`
	Title        string             `json:"title,omitempty" validate:"required"`
	DateOfBirth  time.Time          `json:"dateOfBirth,omitempty" validate:"required"`
	RegisterDate time.Time          `json:"registerDate,omitempty" validate:"required"`
	Status       UserStatus         `json:"status,omitempty" validate:"required"`
	DNI      

}
