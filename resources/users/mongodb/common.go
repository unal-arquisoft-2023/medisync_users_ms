package mongodb

import (
	"medysinc_user_ms/domain"
	"medysinc_user_ms/resources/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoName struct {
	FirstName string `bson:"firstName,omitempty"`
	LastName  string `bson:"lastName,omitempty"`
}

func (mn *MongoName) toDomain() domain.Name {
	return domain.Name{
		FirstName: mn.FirstName,
		LastName:  mn.LastName,
	}
}

func mongoNameFromDomain(name domain.Name) MongoName {
	return MongoName{
		FirstName: name.FirstName,
		LastName:  name.LastName,
	}
}

type MongoLocation struct {
	Country string `bson:"country,omitempty"`
	City    string `bson:"city,omitempty"`
	Address string `bson:"address,omitempty"`
}

func (mn *MongoLocation) toDomain() domain.Location {
	return domain.Location{
		Country: mn.Country,
		City:    mn.City,
		Address: mn.Address,
	}
}

func mongoLocFromDomain(loc domain.Location) MongoLocation {
	return MongoLocation{
		Country: loc.Country,
		City:    loc.City,
		Address: loc.Address,
	}
}

// A struct to manage users in the mongo database
// main difference, the id is changed from string to primitive.ObjectID
type MongoUser struct {
	Id               primitive.ObjectID `bson:"_id,omitempty"`
	Name             MongoName          `bson:"name,omitempty"`
	Email            string             `bson:"email,omitempty"`
	Phone            string             `bson:"phone,omitempty"`
	Location         MongoLocation      `bson:"location,omitempty"`
	DateOfBirth      string             `bson:"dateOfBirth,omitempty"`
	RegistrationDate string             `bson:"registrationDate,omitempty"`
	Status           domain.UserStatus  `bson:"status,omitempty"`
	CardId           string             `bson:"CardId,omitempty"`
}

func (mu *MongoUser) toDomain() domain.User {
	return domain.User{
		ID:               mu.Id.Hex(),
		Name:             mu.Name.toDomain(),
		Email:            mu.Email,
		Phone:            mu.Phone,
		Location:         mu.Location.toDomain(),
		DateOfBirth:      mu.DateOfBirth,
		RegistrationDate: mu.RegistrationDate,
		Status:           mu.Status,
		CardID:           mu.CardId,
	}
}

func mongoUserFromDomain(user domain.User) (MongoUser, error) {
	id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return MongoUser{}, err
	}
	return MongoUser{
		Id:               id,
		Name:             mongoNameFromDomain(user.Name),
		Email:            user.Email,
		Phone:            user.Phone,
		Location:         mongoLocFromDomain(user.Location),
		DateOfBirth:      user.DateOfBirth,
		RegistrationDate: user.RegistrationDate,
		Status:           user.Status,
		CardId:           user.CardID,
	}, nil
}

// Utility function to easily create an invalid id error
func MakeInvalidIdError(id string) users.UserRepositoryError {
	return &users.InvalidPropertyError[string]{
		Property: "id",
		Value:    id,
		Reason:   "invalid id",
	}
}

// Utility function to easily create an id not found error
func MakeNotFoundIdError(resource string, id string) users.UserRepositoryError {
	return &users.NotFoundError[string]{
		Resource: resource,
		Property: "id",
		Value:    id,
	}
}
