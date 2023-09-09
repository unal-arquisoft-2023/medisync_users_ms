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

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func GetDoctor(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("doctorId")

	var doctor models.Doctor
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := doctorCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&doctor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	var doctorResponse dtos.CreateDoctorRequest
	doctorResponse.Id = user.Id
	doctorResponse.Name = dtos.Name(user.Name)
	doctorResponse.Email = user.Email
	doctorResponse.Phone = user.Phone
	doctorResponse.Location = dtos.Location(user.Location)
	doctorResponse.Title = user.Title
	doctorResponse.DateOfBirth = user.DateOfBirth
	doctorResponse.RegisterDate = user.RegisterDate
	doctorResponse.Status = dtos.UserStatus(user.Status)
	doctorResponse.DNI = user.DNI
	doctorResponse.Specialty = dtos.DoctorSpecialty(doctor.Specialty)
	doctorResponse.MedicalLicenseID = doctor.MedicalLicenseID

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": doctorResponse}})
}

func UpdateDoctor(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("doctorId")

	var doctor models.Doctor
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := doctorCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&doctor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	var doctorRequest dtos.CreateDoctorRequest
	doctorRequest.Id = user.Id
	doctorRequest.Name = dtos.Name(user.Name)
	doctorRequest.Email = user.Email
	doctorRequest.Phone = user.Phone
	doctorRequest.Location = dtos.Location(user.Location)
	doctorRequest.Title = user.Title
	doctorRequest.DateOfBirth = user.DateOfBirth
	doctorRequest.RegisterDate = user.RegisterDate
	doctorRequest.Status = dtos.UserStatus(user.Status)
	doctorRequest.DNI = user.DNI
	doctorRequest.Specialty = dtos.DoctorSpecialty(doctor.Specialty)
	doctorRequest.MedicalLicenseID = doctor.MedicalLicenseID

	if err := c.Bind(&doctorRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})

	}

	if err := validate.Struct(doctorRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})

	}

	newDoctor := models.Doctor{
		Id:               doctorRequest.Id,
		UserId:           doctorRequest.Id,
		Specialty:        models.DoctorSpecialty(doctorRequest.Specialty),
		MedicalLicenseID: doctorRequest.MedicalLicenseID,
	}

	newUser := models.User{
		Id:           doctorRequest.Id,
		Name:         models.Name(doctorRequest.Name),
		Email:        doctorRequest.Email,
		Phone:        doctorRequest.Phone,
		Location:     models.Location(doctorRequest.Location),
		Title:        doctorRequest.Title,
		DateOfBirth:  doctorRequest.DateOfBirth,
		RegisterDate: doctorRequest.RegisterDate,
		Status:       models.UserStatus(doctorRequest.Status),
		DNI:          doctorRequest.DNI,
	}

	_, err = doctorCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": newDoctor})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": newUser})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": newUser}})

}

func GetAllDoctors(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var doctorRequest []dtos.CreateDoctorRequest
	defer cancel()

	cursor, err := doctorCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	for cursor.Next(ctx) {
		var doctor models.Doctor
		err := cursor.Decode(&doctor)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		var user models.User
		err = userCollection.FindOne(ctx, bson.M{"_id": doctor.UserId}).Decode(&user)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		var doctorResponse dtos.CreateDoctorRequest
		doctorResponse.Id = user.Id
		doctorResponse.Name = dtos.Name(user.Name)
		doctorResponse.Email = user.Email
		doctorResponse.Phone = user.Phone
		doctorResponse.Location = dtos.Location(user.Location)
		doctorResponse.Title = user.Title
		doctorResponse.DateOfBirth = user.DateOfBirth
		doctorResponse.RegisterDate = user.RegisterDate
		doctorResponse.Status = dtos.UserStatus(user.Status)
		doctorResponse.DNI = user.DNI
		doctorResponse.Specialty = dtos.DoctorSpecialty(doctor.Specialty)
		doctorResponse.MedicalLicenseID = doctor.MedicalLicenseID

		doctorRequest = append(doctorRequest, doctorResponse)
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": doctorRequest}})
}
