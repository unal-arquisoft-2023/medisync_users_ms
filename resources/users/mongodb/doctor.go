package mongodb

import (
	"context"
	"medysinc_user_ms/domain"
	"medysinc_user_ms/resources/users"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDoctorRepository struct {
	coll *mongo.Collection
}

// Creates a new mongo Doctor repository
func NewMongoDoctorRepository(coll *mongo.Collection) users.DoctorRepository {
	return &MongoDoctorRepository{coll}
}

type mongoDoctor struct {
	Specialty        domain.DoctorSpecialty `bson:"specialty,omitempty"`
	MedicalLicenseID string                 `bson:"medicalLicenseID,omitempty"`
	MongoUser        `bson:",inline"`
}

// Converts a mongo Doctor to a domain Doctor
func (mp *mongoDoctor) toDomain() domain.Doctor {
	return domain.Doctor{
		User:             mp.MongoUser.toDomain(),
		Specialty:        mp.Specialty,
		MedicalLicenseID: mp.MedicalLicenseID,
	}
}

// Lets you find a Doctor by its id
func (r *MongoDoctorRepository) FindOne(ctx context.Context, id string) (*domain.Doctor, users.UserRepositoryError) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, MakeInvalidIdError(id)
	}

	var mongoPat mongoDoctor

	err = r.coll.FindOne(ctx, primitive.M{"_id": objId}).Decode(&mongoPat)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, MakeNotFoundIdError("doctor", id)
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	doctor := mongoPat.toDomain()
	return &doctor, nil
}

// Lets you create a doctor
func (r *MongoDoctorRepository) Create(
	ctx context.Context,
	input users.DoctorCreationInput,
) (*domain.Doctor, users.UserRepositoryError) {

	mongoPat := mongoDoctor{
		MongoUser: MongoUser{
			Id:                    primitive.NewObjectID(),
			Name:                  mongoNameFromDomain(input.Name),
			Email:                 input.Email,
			Phone:                 input.Phone,
			Location:              mongoLocFromDomain(input.Location),
			DateOfBirth:           primitive.NewDateTimeFromTime(input.DateOfBirth),
			RegistrationTimeStamp: primitive.NewDateTimeFromTime(time.Now()),
			Status:                input.Status,
			CardId:                input.CardID,
		},
		Specialty:        input.Specialty,
		MedicalLicenseID: input.MedicalLicenseID,
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

	doctor := mongoPat.toDomain()
	return &doctor, nil
}

// Lets you update a doctor
func (r *MongoDoctorRepository) Update(
	ctx context.Context,
	input users.DoctorUpdateInput,
) (*domain.Doctor, users.UserRepositoryError) {
	objId, err := primitive.ObjectIDFromHex(input.ID)

	if err != nil {
		return nil, MakeInvalidIdError(input.ID)
	}

	mongoPat := mongoDoctor{
		MongoUser: MongoUser{
			Id:          objId,
			Name:        mongoNameFromDomain(input.Name),
			Email:       input.Email,
			Phone:       input.Phone,
			Location:    mongoLocFromDomain(input.Location),
			DateOfBirth: primitive.NewDateTimeFromTime(input.DateOfBirth),
			Status:      input.Status,
			CardId:      input.CardID,
		},
		Specialty:        input.Specialty,
		MedicalLicenseID: input.MedicalLicenseID,
	}

	var updatedMongoPat mongoDoctor
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = r.coll.FindOneAndUpdate(
		ctx,
		primitive.M{"_id": objId},
		primitive.M{"$set": mongoPat},
		opts,
	).Decode(&updatedMongoPat)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, MakeNotFoundIdError("doctor", input.ID)
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	doctor := updatedMongoPat.toDomain()
	return &doctor, nil
}

// function that lets you change the status of a doctor to suspended
func (r *MongoDoctorRepository) Suspend(ctx context.Context, id string) (*domain.Doctor, users.UserRepositoryError) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, MakeInvalidIdError(id)
	}

	var mongoPat mongoDoctor

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = r.coll.FindOneAndUpdate(
		ctx,
		primitive.M{"_id": objId},
		primitive.M{"$set": primitive.M{"status": domain.Suspended}},
		opts,
	).Decode(&mongoPat)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, MakeNotFoundIdError("doctor", id)
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	doctor := mongoPat.toDomain()
	return &doctor, nil

}

func (r *MongoDoctorRepository) Activate(ctx context.Context, id string) (*domain.Doctor, users.UserRepositoryError) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, MakeInvalidIdError(id)
	}

	var mongoPat mongoDoctor

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = r.coll.FindOneAndUpdate(
		ctx,
		primitive.M{"_id": objId},
		primitive.M{"$set": primitive.M{"status": domain.Active}},
		opts,
	).Decode(&mongoPat)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, MakeNotFoundIdError("doctor", id)
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	doctor := mongoPat.toDomain()
	return &doctor, nil
}

// Returns all doctors
func (r *MongoDoctorRepository) FindAll(ctx context.Context) ([]domain.Doctor, users.UserRepositoryError) {
	var mongoPats []mongoDoctor

	cursor, err := r.coll.Find(ctx, primitive.M{})

	if err != nil {
		return nil, &users.InternalError{
			Message: err.Error(),
		}
	}

	err = cursor.All(ctx, &mongoPats)

	if err != nil {
		return nil, &users.InternalError{
			Message: err.Error(),
		}
	}

	doctors := make([]domain.Doctor, len(mongoPats))
	for i, mongoPat := range mongoPats {
		doctors[i] = mongoPat.toDomain()
	}

	return doctors, nil
}

// Returns all doctors by specialty
func (r *MongoDoctorRepository) FindBySpecialty(ctx context.Context, specialty domain.DoctorSpecialty) ([]domain.Doctor, users.UserRepositoryError) {
	var mongoPats []mongoDoctor

	cursor, err := r.coll.Find(ctx, primitive.M{"specialty": specialty})

	if err != nil {
		return nil, &users.InternalError{
			Message: err.Error(),
		}
	}

	err = cursor.All(ctx, &mongoPats)

	if err != nil {
		return nil, &users.InternalError{
			Message: err.Error(),
		}
	}

	doctors := make([]domain.Doctor, len(mongoPats))
	for i, mongoPat := range mongoPats {
		doctors[i] = mongoPat.toDomain()
	}

	return doctors, nil
}
