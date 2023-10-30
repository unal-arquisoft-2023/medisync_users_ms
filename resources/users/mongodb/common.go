package mongodb

import (
	"medysinc_user_ms/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// A struct to manage users in the mongo database
// main difference, the id is changes from string to primitive.ObjectID
type mongoUser struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name             domain.Name        `json:"name,omitempty" validate:"required"`
	Email            string             `json:"email,omitempty" validate:"required,email"`
	Phone            string             `json:"phone,omitempty" validate:"required"`
	Location         domain.Location    `json:"location,omitempty" validate:"required"`
	DateOfBirth      string             `json:"dateOfBirth,omitempty" validate:"required"`
	RegistrationDate string             `json:"registrationDate,omitempty" validate:"required"`
	Status           domain.UserStatus  `json:"status,omitempty" validate:"required"`
	CardId           string             `json:"CardId,omitempty" validate:"required"`
}
