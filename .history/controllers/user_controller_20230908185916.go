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
var doctorCollection *mongo.Collection = configs.GetCollection(configs.DB, "doctors")
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

func GetStaffMember(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")

	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}
	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": user}})
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

func UpdateStaffMember(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")

	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": user})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: fmt.Sprintf(" error %s", objId), Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": user}})
}

// update patient in patient collection and user in user collection at the same time using CreatePatientRequest struct

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

func updateDoctor(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("doctorId")

	var doctor models.Doctor
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := patientCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&doctor)

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
