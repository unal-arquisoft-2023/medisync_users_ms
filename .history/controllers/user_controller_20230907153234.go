package controllers

import (
	"context"
	dtos "medysinc_user_ms/DTOs"
	"medysinc_user_ms/models"

	"medysinc_user_ms/configs"
	"medysinc_user_ms/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreatePatient(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var person dtos.CreatePatientRequest
	var user models.User
	var patient models.Patient

	defer cancel()

	if err := c.Bind(&person); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if err := validate.Struct(person); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	newPerson := dtos.CreatePatientRequest{
		Id:           primitive.NewObjectID(),
		Name:         person.Name,
		Email:        person.Email,
		Phone:        person.Phone,
		Location:     person.Location,
		Title:        person.Title,
		DateOfBirth:  person.DateOfBirth,
		RegisterDate: person.RegisterDate,
		Status:       person.Status,
		DNI:          person.DNI,
		Affiliation:  person.Affiliation,
	}

	newUser := models.User{
		Id:           newPerson.Id,
		Name:         newPerson.Name,
		Email:        newPerson.Email,
		Phone:        newPerson.Phone,
		Location:     newPerson.Location,
		Title:        newPerson.Title,
		DateOfBirth:  newPerson.DateOfBirth,
		RegisterDate: newPerson.RegisterDate,
		Status:       newPerson.Status,
		DNI:          newPerson.DNI,
	}

	newPatient := models.Patient{
		Id:          newPerson.Id,
		UserId:      newPerson.Id,
		Affiliation: newPerson.Affiliation,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}
