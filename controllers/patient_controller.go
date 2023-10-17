package controllers

import (
	"context"
	dtos "medysinc_user_ms/DTOs"
	db "medysinc_user_ms/db"
	"medysinc_user_ms/models"

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
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(person); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newUser := models.User{
		Id:               primitive.NewObjectID(),
		Name:             models.Name(person.Name),
		Email:            person.Email,
		Phone:            person.Phone,
		Location:         models.Location(person.Location),
		DateOfBirth:      person.DateOfBirth,
		RegistrationDate: person.RegistrationDate,
		Status:           models.UserStatus(person.Status),
		CardId:           person.RegistrationDate,
	}

	newPatient := models.Patient{
		Id:          primitive.NewObjectID(),
		UserId:      newUser.Id,
		Affiliation: models.PatientAffiliation(person.Affiliation),
	}

	//Inserting the User

	_, err := db.Collections.Users.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	//Inserting the patient

	result, err := db.Collections.Patient.InsertOne(ctx, newPatient)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": result.InsertedID})

}

func GetPatient(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	UserId := c.Param("patientId")

	var patient models.Patient
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(UserId)
	err := db.Collections.Patient.FindOne(ctx, bson.M{"_id": objId}).Decode(&patient)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var user models.User
	err = db.Collections.Users.FindOne(ctx, bson.M{"_id": patient.UserId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var userResponse dtos.PatientResponse
	userResponse.Id = patient.Id.Hex()
	userResponse.Name = user.Name
	userResponse.Email = user.Email
	userResponse.Phone = user.Phone
	userResponse.Location = user.Location
	userResponse.DateOfBirth = user.DateOfBirth
	userResponse.RegistrationDate = user.RegistrationDate
	userResponse.Status = user.Status
	userResponse.CardId = user.CardId
	userResponse.Affiliation = patient.Affiliation

	return c.JSON(http.StatusOK, userResponse)
}

func UpdatePatient(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	UserId := c.Param("patientId")

	var patient models.Patient
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(UserId)
	err := db.Collections.Patient.FindOne(ctx, bson.M{"_id": objId}).Decode(&patient)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var user models.User
	err = db.Collections.Users.FindOne(ctx, bson.M{"_id": patient.UserId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var PatReq dtos.CreatePatientRequest
	if err := c.Bind(&PatReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	if err := validate.Struct(PatReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	user.CardId = PatReq.CardId
	user.Email = PatReq.Email
	user.Location = PatReq.Location
	user.Name = PatReq.Name
	user.Phone = PatReq.Phone
	user.RegistrationDate = PatReq.RegistrationDate
	user.Status = PatReq.Status
	user.DateOfBirth = PatReq.DateOfBirth

	patient.Affiliation = models.PatientAffiliation(PatReq.Affiliation)

	_, err = db.Collections.Patient.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": patient})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = db.Collections.Users.UpdateOne(ctx, bson.M{"_id": patient.UserId}, bson.M{"$set": user})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "updated")

}

func GetAllPatients(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var allPatients []dtos.PatientResponse
	defer cancel()

	cursor, err := db.Collections.Patient.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for cursor.Next(ctx) {
		var patient models.Patient
		err := cursor.Decode(&patient)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var user models.User
		err = db.Collections.Users.FindOne(ctx, bson.M{"_id": patient.UserId}).Decode(&user)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var patientResponse dtos.PatientResponse
		patientResponse.Id = patient.Id.Hex()
		patientResponse.Name = user.Name
		patientResponse.Email = user.Email
		patientResponse.Phone = user.Phone
		patientResponse.Location = user.Location
		patientResponse.DateOfBirth = user.DateOfBirth
		patientResponse.RegistrationDate = user.RegistrationDate
		patientResponse.Status = user.Status
		patientResponse.CardId = user.CardId
		patientResponse.Affiliation = patient.Affiliation

		allPatients = append(allPatients, patientResponse)
	}

	return c.JSON(http.StatusOK, allPatients)
}
