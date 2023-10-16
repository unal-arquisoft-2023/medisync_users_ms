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
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDoctor(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var docReq dtos.CreateDoctorRequest

	defer cancel()

	if err := c.Bind(&docReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(docReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newUser := models.User{
		Id:           primitive.NewObjectID(),
		Name:         models.Name(docReq.Name),
		Email:        docReq.Email,
		Phone:        docReq.Phone,
		Location:     models.Location(docReq.Location),
		DateOfBirth:  docReq.DateOfBirth,
		RegisterDate: docReq.RegisterDate,
		Status:       models.UserStatus(docReq.Status),
		DNI:          docReq.DNI,
	}

	newDoctor := models.Doctor{
		Id:               primitive.NewObjectID(),
		UserId:           newUser.Id,
		Specialty:        models.DoctorSpecialty(docReq.Specialty),
		MedicalLicenseID: docReq.MedicalLicenseID,
	}

	// Inserting the User

	_, err := db.Collections.Users.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Inserting the Doctor

	result, err := db.Collections.Doctor.InsertOne(ctx, newDoctor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": result.InsertedID})

}

func GetDoctor(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	UserId := c.Param("doctorId")

	var doctor models.Doctor
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(UserId)
	err := db.Collections.Doctor.FindOne(ctx, bson.M{"_id": objId}).Decode(&doctor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var user models.User
	err = db.Collections.Users.FindOne(ctx, bson.M{"_id": doctor.UserId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var DoctorResponse dtos.DoctorResponse
	DoctorResponse.Id = doctor.Id.Hex()
	DoctorResponse.Name = user.Name
	DoctorResponse.Email = user.Email
	DoctorResponse.Phone = user.Phone
	DoctorResponse.Location = user.Location
	DoctorResponse.DateOfBirth = user.DateOfBirth
	DoctorResponse.RegisterDate = user.RegisterDate
	DoctorResponse.Status = user.Status
	DoctorResponse.DNI = user.DNI
	DoctorResponse.Specialty = doctor.Specialty
	DoctorResponse.MedicalLicenseID = doctor.MedicalLicenseID

	return c.JSON(http.StatusOK, DoctorResponse)
}

func UpdateDoctor(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	UserId := c.Param("doctorId")

	var doctor models.Doctor
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(UserId)
	err := db.Collections.Doctor.FindOne(ctx, bson.M{"_id": objId}).Decode(&doctor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var user models.User
	err = db.Collections.Users.FindOne(ctx, bson.M{"_id": doctor.UserId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var docReq dtos.CreateDoctorRequest

	if err := c.Bind(&docReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	if err := validate.Struct(docReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	doctor.MedicalLicenseID = docReq.MedicalLicenseID
	doctor.Specialty = docReq.Specialty

	user.DNI = docReq.DNI
	user.Email = docReq.Email
	user.Location = docReq.Location
	user.Name = docReq.Name
	user.Phone = docReq.Phone
	user.RegisterDate = docReq.RegisterDate
	user.Status = docReq.Status
	user.DateOfBirth = docReq.DateOfBirth

	_, err = db.Collections.Doctor.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": doctor})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = db.Collections.Users.UpdateOne(ctx, bson.M{"_id": doctor.UserId}, bson.M{"$set": user})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "updated")

}

func GetAllDoctors(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	allDoctors := []dtos.DoctorResponse{}
	defer cancel()

	specialty := c.QueryParam("specialty")
	var cursor *mongo.Cursor
	var err error
	if specialty == "" {
		cursor, err = db.Collections.Doctor.Find(ctx, bson.M{})
	} else {
		cursor, err = db.Collections.Doctor.Find(ctx, bson.M{"specialty": specialty})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for cursor.Next(ctx) {
		var doctor models.Doctor
		err := cursor.Decode(&doctor)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var user models.User
		err = db.Collections.Users.FindOne(ctx, bson.M{"_id": doctor.UserId}).Decode(&user)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var docRes dtos.DoctorResponse
		docRes.Id = user.Id.Hex()
		docRes.Name = user.Name
		docRes.Email = user.Email
		docRes.Phone = user.Phone
		docRes.Location = user.Location
		docRes.DateOfBirth = user.DateOfBirth
		docRes.RegisterDate = user.RegisterDate
		docRes.Status = user.Status
		docRes.DNI = user.DNI
		docRes.Specialty = doctor.Specialty
		docRes.MedicalLicenseID = doctor.MedicalLicenseID

		allDoctors = append(allDoctors, docRes)
	}

	return c.JSON(http.StatusOK, allDoctors)
}
