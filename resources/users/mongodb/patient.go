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
	Affiliation domain.PatientAffiliation `bson:"affiliation,omitempty"`
	MongoUser   `bson:",inline"`
}

// Converts a mongo patient to a domain patient
func (mp *mongoPatient) toDomain() domain.Patient {
	return domain.Patient{
		User:        mp.MongoUser.toDomain(),
		Affiliation: mp.Affiliation,
	}
}

// Lets you find a patient by its id
func (r *MongoPatientRepository) FindOne(ctx context.Context, id string) (*domain.Patient, users.UserRepositoryError) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, &users.InvalidPropertyError[string]{
			Property: "id",
			Value:    id,
			Reason:   "invalid id",
		}
	}

	var mongoPat mongoPatient

	err = r.coll.FindOne(ctx, primitive.M{"_id": objId}).Decode(&mongoPat)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &users.NotFoundError[string]{
				Resource: "patient",
				Property: "id",
				Value:    id,
			}
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	patient := mongoPat.toDomain()
	return &patient, nil
}

// Lets you create a patient
func (r *MongoPatientRepository) Create(
	ctx context.Context,
	input users.PatientCreationInput,
) (*domain.Patient, users.UserRepositoryError) {

	mongoPat := mongoPatient{
		MongoUser: MongoUser{
			Id:               primitive.NewObjectID(),
			Name:             mongoNameFromDomain(input.Name),
			Email:            input.Email,
			Phone:            input.Phone,
			Location:         mongoLocFromDomain(input.Location),
			DateOfBirth:      input.DateOfBirth,
			RegistrationDate: input.RegistrationDate,
			Status:           input.Status,
			CardId:           input.CardID,
		},
		Affiliation: input.Affiliation,
	}

	_, err := r.coll.InsertOne(ctx, mongoPat)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, &users.AlreadyExistsError[string]{
				Property: "id",
				Value:    mongoPat.Id.Hex(),
			}
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	patient := mongoPat.toDomain()
	return &patient, nil
}

// Lets you update a patient
func (r *MongoPatientRepository) Update(
	ctx context.Context,
	input users.PatientUpdateInput,
) (*domain.Patient, users.UserRepositoryError) {
	objId, err := primitive.ObjectIDFromHex(input.ID)

	if err != nil {
		return nil, &users.InvalidPropertyError[string]{
			Property: "id",
			Value:    input.ID,
			Reason:   "invalid id",
		}
	}

	mongoPat := mongoPatient{
		MongoUser: MongoUser{
			Id:               objId,
			Name:             mongoNameFromDomain(input.Name),
			Email:            input.Email,
			Phone:            input.Phone,
			Location:         mongoLocFromDomain(input.Location),
			DateOfBirth:      input.DateOfBirth,
			RegistrationDate: input.RegistrationDate,
			Status:           input.Status,
			CardId:           input.CardID,
		},
		Affiliation: input.Affiliation,
	}

	var updatedMongoPat mongoPatient
	err = r.coll.FindOneAndUpdate(ctx, primitive.M{"_id": objId}, primitive.M{"$set": mongoPat}).Decode(&updatedMongoPat)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &users.NotFoundError[string]{
				Resource: "patient",
				Property: "id",
				Value:    input.ID,
			}
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	patient := updatedMongoPat.toDomain()
	return &patient, nil
}

// function that lets you change the status of a patient to suspended
func (r *MongoPatientRepository) Suspend(ctx context.Context, id string) (*domain.Patient, users.UserRepositoryError) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, &users.InvalidPropertyError[string]{
			Property: "id",
			Value:    id,
			Reason:   "invalid id",
		}
	}

	var mongoPat mongoPatient

	err = r.coll.FindOneAndUpdate(
		ctx,
		primitive.M{"_id": objId},
		primitive.M{"$set": primitive.M{"status": domain.Suspended}},
	).Decode(&mongoPat)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &users.NotFoundError[string]{
				Resource: "patient",
				Property: "id",
				Value:    id,
			}
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	patient := mongoPat.toDomain()
	return &patient, nil

}

func (r *MongoPatientRepository) Activate(ctx context.Context, id string) (*domain.Patient, users.UserRepositoryError) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, &users.InvalidPropertyError[string]{
			Property: "id",
			Value:    id,
			Reason:   "invalid id",
		}
	}

	var mongoPat mongoPatient

	err = r.coll.FindOneAndUpdate(
		ctx,
		primitive.M{"_id": objId},
		primitive.M{"$set": primitive.M{"status": domain.Active}},
	).Decode(&mongoPat)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &users.NotFoundError[string]{
				Resource: "patient",
				Property: "id",
				Value:    id,
			}
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	patient := mongoPat.toDomain()
	return &patient, nil
}
