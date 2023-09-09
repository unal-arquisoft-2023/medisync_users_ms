package controllers

import (
	"context"
	"fmt"
	dtos "medysinc_user_ms/DTOs"
	"medysinc_user_ms/models"

	"medysinc_user_ms/configs"
	"medysinc_user_ms/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var patientCollection *mongo.Collection = configs.GetCollection(configs.DB, "patients")
var validate = validator.New()

func CreatePatient(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var person dtos.CreatePatientRequest

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
		Name:         models.Name(newPerson.Name),
		Email:        newPerson.Email,
		Phone:        newPerson.Phone,
		Location:     models.Location(newPerson.Location),
		Title:        newPerson.Title,
		DateOfBirth:  newPerson.DateOfBirth,
		RegisterDate: newPerson.RegisterDate,
		Status:       models.UserStatus(newPerson.Status),
		DNI:          newPerson.DNI,
	}

	newPatient := models.Patient{
		Id:          newPerson.Id,
		UserId:      newPerson.Id,
		Affiliation: models.PatientAffiliation(newPerson.Affiliation),
	}

	// Inserting the user

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	// Inserting the patient

	patientCollection := configs.GetCollection(configs.DB, "patients")
	_, err = patientCollection.InsertOne(ctx, newPatient)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result.InsertedID}})

}

func CreateDoctor(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var person dtos.CreateDoctorRequest

	defer cancel()

	if err := c.Bind(&person); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if err := validate.Struct(person); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	newPerson := dtos.CreateDoctorRequest{
		Id:               primitive.NewObjectID(),
		Name:             person.Name,
		Email:            person.Email,
		Phone:            person.Phone,
		Location:         person.Location,
		Title:            person.Title,
		DateOfBirth:      person.DateOfBirth,
		RegisterDate:     person.RegisterDate,
		Status:           person.Status,
		DNI:              person.DNI,
		Specialty:        person.Specialty,
		MedicalLicenseID: person.MedicalLicenseID,
	}

	newUser := models.User{
		Id:           newPerson.Id,
		Name:         models.Name(newPerson.Name),
		Email:        newPerson.Email,
		Phone:        newPerson.Phone,
		Location:     models.Location(newPerson.Location),
		Title:        newPerson.Title,
		DateOfBirth:  newPerson.DateOfBirth,
		RegisterDate: newPerson.RegisterDate,
		Status:       models.UserStatus(newPerson.Status),
		DNI:          newPerson.DNI,
	}

	newDoctor := models.Doctor{
		Id:               newPerson.Id,
		UserId:           newPerson.Id,
		Specialty:        models.DoctorSpecialty(newPerson.Specialty),
		MedicalLicenseID: newPerson.MedicalLicenseID,
	}

	// Inserting the user

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	// Inserting the doctor

	doctorCollection := configs.GetCollection(configs.DB, "doctors")
	_, err = doctorCollection.InsertOne(ctx, newDoctor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result.InsertedID}})

}

func CreateStaff(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User

	defer cancel()

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if err := validate.Struct(user); err != nil {
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
		RegisterDate: user.RegisterDate,
		Status:       user.Status,
		DNI:          user.DNI,
	}

	// Inserting the user

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result.InsertedID}})

}

func GetPatient(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("patientId")

	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}
	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": user}})
}

func GetStaffMember(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")

	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}
	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": user}})
}
