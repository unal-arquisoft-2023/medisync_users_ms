package mongodb

import (
	"context"
	"medysinc_user_ms/domain"
	"medysinc_user_ms/resources/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoPatientRepository struct {
	coll *mongo.Collection
}

// Creates a new mongo patient repository
func NewMongoPatientRepository(coll *mongo.Collection) users.PatientRepository {
	return &MongoPatientRepository{coll}
}

type mongoPatient struct {
	Affiliation domain.PatientAffiliation `json:"affiliation,omitempty" validate:"required"`
	mongoUser
}

// Lets you find a patient by its id
func (r *MongoPatientRepository) FindOne(ctx context.Context, id string) (*domain.Patient, error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var mongoPat mongoPatient
	var patient domain.Patient

	err = r.coll.FindOne(ctx, primitive.M{"_id": objId}).Decode(&mongoPat)

	if err != nil {
		return nil, err
	}

	patient = domain.Patient{
		User: domain.User{
			ID:               mongoPat.Id.Hex(),
			Name:             mongoPat.Name,
			Email:            mongoPat.Email,
			Phone:            mongoPat.Phone,
			Location:         mongoPat.Location,
			DateOfBirth:      mongoPat.DateOfBirth,
			RegistrationDate: mongoPat.RegistrationDate,
			Status:           mongoPat.Status,
			CardID:           mongoPat.CardId,
		},
		Affiliation: mongoPat.Affiliation,
	}

	return &patient, nil
}

// Lets you create a patient
func (r *MongoPatientRepository) Create(
	ctx context.Context,
	input users.PatientCreationInput,
) (*domain.Patient, error) {

	mongoPat := mongoPatient{
		mongoUser: mongoUser{
			Id:               primitive.NewObjectID(),
			Name:             input.Name,
			Email:            input.Email,
			Phone:            input.Phone,
			Location:         input.Location,
			DateOfBirth:      input.DateOfBirth,
			RegistrationDate: input.RegistrationDate,
			Status:           input.Status,
			CardId:           input.CardID,
		},
		Affiliation: input.Affiliation,
	}
	_, err := r.coll.InsertOne(ctx, mongoPat)

	if err != nil {
		return nil, err
	}

	patient := domain.Patient{
		User: domain.User{
			ID:               mongoPat.Id.Hex(),
			Name:             mongoPat.Name,
			Email:            mongoPat.Email,
			Phone:            mongoPat.Phone,
			Location:         mongoPat.Location,
			DateOfBirth:      mongoPat.DateOfBirth,
			RegistrationDate: mongoPat.RegistrationDate,
			Status:           mongoPat.Status,
			CardID:           mongoPat.CardId,
		},
		Affiliation: mongoPat.Affiliation,
	}
	return &patient, nil
}

// Lets you update a patient
func (r *MongoPatientRepository) Update(
	ctx context.Context,
	input users.PatientUpdateInput,
) (*domain.Patient, error) {
	objId, err := primitive.ObjectIDFromHex(input.ID)

	if err != nil {
		return nil, err
	}

	mongoPat := mongoPatient{
		mongoUser: mongoUser{
			Id:               objId,
			Name:             input.Name,
			Email:            input.Email,
			Phone:            input.Phone,
			Location:         input.Location,
			DateOfBirth:      input.DateOfBirth,
			RegistrationDate: input.RegistrationDate,
			Status:           input.Status,
			CardId:           input.CardID,
		},
		Affiliation: input.Affiliation,
	}

	_, err = r.coll.UpdateOne(ctx, primitive.M{"_id": objId}, primitive.M{"$set": mongoPat})

	if err != nil {
		return nil, err
	}

	patient := domain.Patient{
		User: domain.User{
			ID:               mongoPat.Id.Hex(),
			Name:             mongoPat.Name,
			Email:            mongoPat.Email,
			Phone:            mongoPat.Phone,
			Location:         mongoPat.Location,
			DateOfBirth:      mongoPat.DateOfBirth,
			RegistrationDate: mongoPat.RegistrationDate,
			Status:           mongoPat.Status,
			CardID:           mongoPat.CardId,
		},
		Affiliation: mongoPat.Affiliation,
	}
	return &patient, nil
}

// function that lets you delete a patient
func (r *MongoPatientRepository) Delete(ctx context.Context, id string) (*domain.Patient, error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var mongoPat mongoPatient

	err = r.coll.FindOneAndDelete(ctx, primitive.M{"_id": objId}).Decode(&mongoPat)

	if err != nil {
		return nil, err
	}

	patient := domain.Patient{
		User: domain.User{
			ID:               mongoPat.Id.Hex(),
			Name:             mongoPat.Name,
			Email:            mongoPat.Email,
			Phone:            mongoPat.Phone,
			Location:         mongoPat.Location,
			DateOfBirth:      mongoPat.DateOfBirth,
			RegistrationDate: mongoPat.RegistrationDate,
			Status:           mongoPat.Status,
			CardID:           mongoPat.CardId,
		},
		Affiliation: mongoPat.Affiliation,
	}
	return &patient, nil
}
