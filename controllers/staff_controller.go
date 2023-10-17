package controllers

import (
	"context"
	"fmt"
	dtos "medysinc_user_ms/DTOs"
	db "medysinc_user_ms/db"
	"medysinc_user_ms/models"

	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

func CreateStaff(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var stf dtos.CreateStaffRequest

	defer cancel()

	if err := c.Bind(&stf); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(stf); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newUser := models.User{
		Id:               primitive.NewObjectID(),
		Name:             stf.Name,
		Email:            stf.Email,
		Phone:            stf.Phone,
		Location:         stf.Location,
		DateOfBirth:      stf.DateOfBirth,
		RegistrationDate: stf.RegistrationDate,
		Status:           stf.Status,
		CardId:           stf.CardId,
	}

	newStaff := models.Staff{
		Id:       primitive.NewObjectID(),
		UserId:   newUser.Id,
		Position: stf.Position,
	}

	//Inserting the user

	_, err := db.Collections.Users.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	resStf, err := db.Collections.Staff.InsertOne(ctx, newStaff)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": resStf.InsertedID})

}

func GetStaffMember(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	staffId := c.Param("staffId")
	defer cancel()

	var staff models.Staff
	var user models.User

	objId, _ := primitive.ObjectIDFromHex(staffId)
	err := db.Collections.Staff.FindOne(ctx, bson.M{"_id": objId}).Decode(&staff)

	if err != nil {
		fmt.Println("Error 1", err)
		if mongo.ErrNilDocument == err {
			return c.JSON(http.StatusNotFound, fmt.Sprintf("Not found %s", objId))
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err2 := db.Collections.Users.FindOne(ctx, bson.M{"_id": staff.UserId}).Decode(&user)

	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, err2.Error())
	}

	var stfRes dtos.StaffResponse

	stfRes.Id = staff.Id.Hex()
	stfRes.Name = user.Name
	stfRes.Email = user.Email
	stfRes.Phone = user.Phone
	stfRes.Location = user.Location
	stfRes.DateOfBirth = user.DateOfBirth
	stfRes.RegistrationDate = user.RegistrationDate
	stfRes.Status = user.Status
	stfRes.CardId = user.CardId
	stfRes.Position = staff.Position

	return c.JSON(http.StatusOK, stfRes)
}

func UpdateStaff(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	UserId := c.Param("staffId")

	var staff models.Staff
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(UserId)
	err := db.Collections.Staff.FindOne(ctx, bson.M{"_id": objId}).Decode(&staff)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var user models.User
	err = db.Collections.Users.FindOne(ctx, bson.M{"_id": staff.UserId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var PatReq dtos.CreateStaffRequest
	if err := c.Bind(&PatReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	if err := validate.Struct(PatReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	user.CardId = PatReq.CardId
	user.Email = PatReq.Email
	user.Location = models.Location(PatReq.Location)
	user.Name = models.Name(PatReq.Name)
	user.Phone = PatReq.Phone
	user.RegistrationDate = PatReq.RegistrationDate
	user.Status = models.UserStatus(PatReq.Status)
	user.DateOfBirth = PatReq.DateOfBirth

	staff.Position = PatReq.Position

	_, err = db.Collections.Staff.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": staff})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = db.Collections.Users.UpdateOne(ctx, bson.M{"_id": staff.UserId}, bson.M{"$set": user})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "updated")

}
