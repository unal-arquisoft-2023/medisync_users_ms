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

type MongoStaffRepository struct {
	coll *mongo.Collection
}

// Creates a new mongo Staff repository
func NewMongoStaffRepository(coll *mongo.Collection) users.StaffRepository {
	return &MongoStaffRepository{coll}
}

type mongoStaff struct {
	MongoUser `bson:",inline"`
}

// Converts a mongo Staff to a domain Staff
func (mp *mongoStaff) toDomain() domain.Staff {
	return domain.Staff{
		User: mp.MongoUser.toDomain(),
	}
}

// Lets you find a Staff by its id
func (r *MongoStaffRepository) FindOne(ctx context.Context, id string) (*domain.Staff, users.UserRepositoryError) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, MakeInvalidIdError(id)
	}

	var mongoPat mongoStaff

	err = r.coll.FindOne(ctx, primitive.M{"_id": objId}).Decode(&mongoPat)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, MakeNotFoundIdError("Staff", id)
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	Staff := mongoPat.toDomain()
	return &Staff, nil
}

// Lets you create a Staff
func (r *MongoStaffRepository) Create(
	ctx context.Context,
	input users.UserCreationInput,
) (*domain.Staff, users.UserRepositoryError) {

	mongoPat := mongoStaff{
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

	staff := mongoPat.toDomain()
	return &staff, nil
}

// Lets you update a staff
func (r *MongoStaffRepository) Update(
	ctx context.Context,
	input users.UserUpdateInput,
) (*domain.Staff, users.UserRepositoryError) {
	objId, err := primitive.ObjectIDFromHex(input.ID)

	if err != nil {
		return nil, MakeInvalidIdError(input.ID)
	}

	mongoPat := mongoStaff{
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
	}

	var updatedMongoPat mongoStaff
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = r.coll.FindOneAndUpdate(
		ctx,
		primitive.M{"_id": objId},
		primitive.M{"$set": mongoPat},
		opts,
	).Decode(&updatedMongoPat)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, MakeNotFoundIdError("staff", input.ID)
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	staff := updatedMongoPat.toDomain()
	return &staff, nil
}

// function that lets you change the status of a staff to suspended
func (r *MongoStaffRepository) Suspend(ctx context.Context, id string) (*domain.Staff, users.UserRepositoryError) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, MakeInvalidIdError(id)
	}

	var mongoPat mongoStaff

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = r.coll.FindOneAndUpdate(
		ctx,
		primitive.M{"_id": objId},
		primitive.M{"$set": primitive.M{"status": domain.Suspended}},
		opts,
	).Decode(&mongoPat)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, MakeNotFoundIdError("staff", id)
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	staff := mongoPat.toDomain()
	return &staff, nil

}

func (r *MongoStaffRepository) Activate(ctx context.Context, id string) (*domain.Staff, users.UserRepositoryError) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, MakeInvalidIdError(id)
	}

	var mongoPat mongoStaff

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = r.coll.FindOneAndUpdate(
		ctx,
		primitive.M{"_id": objId},
		primitive.M{"$set": primitive.M{"status": domain.Active}},
		opts,
	).Decode(&mongoPat)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, MakeNotFoundIdError("staff", id)
		} else {
			return nil, &users.InternalError{
				Message: err.Error(),
			}
		}
	}

	staff := mongoPat.toDomain()
	return &staff, nil
}

// Returns all staffs
func (r *MongoStaffRepository) FindAll(ctx context.Context) ([]domain.Staff, users.UserRepositoryError) {
	var mongoPats []mongoStaff

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

	staffs := make([]domain.Staff, len(mongoPats))
	for i, mongoPat := range mongoPats {
		staffs[i] = mongoPat.toDomain()
	}

	return staffs, nil
}
