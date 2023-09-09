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

	//Inserting the user

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//Inserting the patient

	patientCollection := configs.GetCollection(configs.DB, "patients")
	_, err = patientCollection.InsertOne(ctx, newPatient)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result.InsertedID}})

}

func GetPatient(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("patientId")

	var patient models.Patient
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := patientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&patient)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	var patientResponse dtos.CreatePatientRequest
	patientResponse.Id = user.Id
	patientResponse.Name = dtos.Name(user.Name)
	patientResponse.Email = user.Email
	patientResponse.Phone = user.Phone
	patientResponse.Location = dtos.Location(user.Location)
	patientResponse.Title = user.Title
	patientResponse.DateOfBirth = user.DateOfBirth
	patientResponse.RegisterDate = user.RegisterDate
	patientResponse.Status = dtos.UserStatus(user.Status)
	patientResponse.DNI = user.DNI
	patientResponse.Affiliation = dtos.PatientAffiliation(patient.Affiliation)

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": patientResponse}})
}

func UpdatePatient(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("patientId")

	var patient models.Patient
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := patientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&patient)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	var patientRequest dtos.CreatePatientRequest
	patientRequest.Id = user.Id
	patientRequest.Name = dtos.Name(user.Name)
	patientRequest.Email = user.Email
	patientRequest.Phone = user.Phone
	patientRequest.Location = dtos.Location(user.Location)
	patientRequest.Title = user.Title
	patientRequest.DateOfBirth = user.DateOfBirth
	patientRequest.RegisterDate = user.RegisterDate
	patientRequest.Status = dtos.UserStatus(user.Status)
	patientRequest.DNI = user.DNI
	patientRequest.Affiliation = dtos.PatientAffiliation(patient.Affiliation)

	if err := c.Bind(&patientRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})

	}

	if err := validate.Struct(patientRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})

	}

	newPatient := models.Patient{
		Id:          patientRequest.Id,
		UserId:      patientRequest.Id,
		Affiliation: models.PatientAffiliation(patientRequest.Affiliation),
	}

	newUser := models.User{
		Id:           patientRequest.Id,
		Name:         models.Name(patientRequest.Name),
		Email:        patientRequest.Email,
		Phone:        patientRequest.Phone,
		Location:     models.Location(patientRequest.Location),
		Title:        patientRequest.Title,
		DateOfBirth:  patientRequest.DateOfBirth,
		RegisterDate: patientRequest.RegisterDate,
		Status:       models.UserStatus(patientRequest.Status),
		DNI:          patientRequest.DNI,
	}

	_, err = patientCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": newPatient})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": newUser})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": newUser}})

}

func GetAllPatients(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var patientRequest []dtos.CreatePatientRequest
	defer cancel()

	cursor, err := patientCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	for cursor.Next(ctx) {
		var patient models.Patient
		err := cursor.Decode(&patient)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		var user models.User
		err = userCollection.FindOne(ctx, bson.M{"_id": patient.UserId}).Decode(&user)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		var patientResponse dtos.CreatePatientRequest
		patientResponse.Id = user.Id
		patientResponse.Name = dtos.Name(user.Name)
		patientResponse.Email = user.Email
		patientResponse.Phone = user.Phone
		patientResponse.Location = dtos.Location(user.Location)
		patientResponse.Title = user.Title
		patientResponse.DateOfBirth = user.DateOfBirth
		patientResponse.RegisterDate = user.RegisterDate
		patientResponse.Status = dtos.UserStatus(user.Status)
		patientResponse.DNI = user.DNI
		patientResponse.Affiliation = dtos.PatientAffiliation(patient.Affiliation)

		patientRequest = append(patientRequest, patientResponse)
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": patientRequest}})
}
