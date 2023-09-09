package controllers

import (
	"context"
	"echo-mongo-api/configs"
	"medysinc_user_ms/models"
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
	var patient models.Patient
	var user models.User
	defer cancel()

	if err := c.Bind(&patient); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if err := validate.Struct(patient); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	newUser := models.User{
		Id:           primitive.NewObjectID(),
		Name:         user.Name,
		Email:        user.Email,
		Phone:        user.Phone,
		Location:     user.Location,
		Title:        user.Title,
		DateOfBirth:  user.DateOfBirth,
		RegisterDate: time.Now(),
		Status:       models.Active,
		DNI:          user.DNI,
	}

	newPatient := models.Patient{
		Id:          primitive.NewObjectID(),
		UserId:      newUser.Id,
		Affiliation: patient.Affiliation,
	}

	// Insertar en la base de datos
	insertionResult, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	// Insertar en la base de datos
	insertionResult, err = userCollection.InsertOne(ctx, newPatient)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": insertionResult.InsertedID}})
}
