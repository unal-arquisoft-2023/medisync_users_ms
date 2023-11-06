package mongodb_test

import (
	"context"
	"medysinc_user_ms/domain"
	"medysinc_user_ms/resources/configuration"
	. "medysinc_user_ms/resources/users"
	. "medysinc_user_ms/resources/users/mongodb"
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testDependencies struct {
	patRepo     *PatientRepository
	failPatRepo *PatientRepository
}

func setupTestDependencies(ctx context.Context, t *testing.T) *testDependencies {

	configuration, err := configuration.NewConfigurationGodotEnv("../../../.env")

	if err != nil {
		t.Fatal(err)
		return nil
	}

	mongoURI, err := configuration.Get("MONGO_URI")
	if err != nil {
		t.Fatal(err)
		return nil
	}

	clientOptions :=
		options.Client().
			ApplyURI(mongoURI).
			SetServerSelectionTimeout(time.Millisecond * 100)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		t.Fatal(err)
		return nil
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		t.Fatal(err)
		return nil
	}

	mongoURIFail, err := configuration.Get("MONGO_URI_FAIL")

	if err != nil {
		t.Fatal(err)
		return nil
	}

	clientOptionsFail :=
		options.Client().
			ApplyURI(mongoURIFail).
			SetServerSelectionTimeout(time.Millisecond * 100)

	clientFail, err := mongo.Connect(ctx, clientOptionsFail)

	if err != nil {
		t.Fatal(err)
		return nil
	}

	mongoDB, err := configuration.Get("MONGO_DB")
	if err != nil {
		t.Fatal(err)
		return nil
	}

	mongoPatientCollection, err := configuration.Get("MONGO_PATIENT_COLLECTION")

	if err != nil {
		t.Fatal("Fail connection did not fail", err)
		return nil
	}

	db := client.Database(mongoDB)
	coll := db.Collection(mongoPatientCollection)
	patRep := NewMongoPatientRepository(coll)

	dbFail := clientFail.Database(mongoDB)
	collFail := dbFail.Collection(mongoPatientCollection)
	patRepFail := NewMongoPatientRepository(collFail)

	return &testDependencies{
		patRepo:     &patRep,
		failPatRepo: &patRepFail,
	}

}

func PatientTestSuite(t *testing.T) {

	ctx := context.Background()
	deps := setupTestDependencies(ctx, t)

	patIn := PatientCreationInput{
		Name:             domain.Name{FirstName: "test", LastName: "test"},
		Email:            "test@mail.com",
		Phone:            "123456789",
		Location:         domain.Location{Country: "test", City: "test", Address: "test"},
		DateOfBirth:      "test",
		RegistrationDate: "test",
		Status:           domain.Active,
		CardID:           "test",
		Affiliation:      domain.Private,
	}

	var pat *domain.Patient

	t.Run("TestCreatePatient", func(t *testing.T) {
		createdPat := testCreatePatient(ctx, *deps.patRepo, patIn, t)
		pat = &createdPat
	})

	t.Run("TestGetPatient", func(t *testing.T) {
		foundPat := testGetPatient(ctx, *deps.patRepo, pat.ID, pat, t)
		pat = &foundPat
	})

	// We update our local copy of the patient, and then update it in the database
	pat.Affiliation = domain.Public
	pat.Name.FirstName = "test2"
	pat.Name.LastName = "test2"
	pat.Email = "test2@mail.com"
	pat.Phone = "987654321"
	pat.Location.Country = "test2"
	pat.Location.City = "test2"
	pat.Location.Address = "test2"
	pat.DateOfBirth = "test2"
	pat.RegistrationDate = "test2"
	pat.CardID = "test2"

	patUp := PatientUpdateInput{
		ID:               pat.ID,
		Name:             pat.Name,
		Email:            pat.Email,
		Phone:            pat.Phone,
		Location:         pat.Location,
		DateOfBirth:      pat.DateOfBirth,
		RegistrationDate: pat.RegistrationDate,
		Status:           pat.Status,
		CardID:           pat.CardID,
		Affiliation:      pat.Affiliation,
	}

	t.Run("TestUpdatePatient", func(t *testing.T) {
		updatedPat := testUpdatePatient(ctx, *deps.patRepo, patUp, pat, t)
		pat = &updatedPat
	})

	t.Run("TestSuspendPatient", func(t *testing.T) {
		suspendedPat := testSuspendPatient(ctx, *deps.patRepo, pat.ID, t)
		pat = &suspendedPat
	})

	t.Run("TestActivatePatient", func(t *testing.T) {
		activatedPat := testActivatePatient(ctx, *deps.patRepo, pat.ID, t)
		pat = &activatedPat
	})

	// TODO: Add tests for error cases

}

func testCreatePatient(
	ctx context.Context,
	patRepo PatientRepository,
	input PatientCreationInput,
	t *testing.T,
) domain.Patient {

	patient, err := patRepo.Create(ctx, input)

	if err != nil {
		t.Fatal("Could not create patient", err)
	}

	return *patient
}

func testGetPatient(
	ctx context.Context,
	patRepo PatientRepository,
	id string,
	originalPatient *domain.Patient,
	t *testing.T,
) domain.Patient {

	patient, err := patRepo.FindOne(ctx, id)

	if err != nil {
		t.Error("Could not get patient", err)
	}

	// Compare the patients to make sure they are the same
	if !reflect.DeepEqual(patient, originalPatient) {
		t.Errorf("Cound not get original patient\nExpected:\n\t%v\ngot\n\t%v", originalPatient, patient)
	}
	return *patient
}

func testUpdatePatient(
	ctx context.Context,
	patRepo PatientRepository,
	input PatientUpdateInput,
	expectedPatient *domain.Patient,
	t *testing.T,
) domain.Patient {

	patient, err := patRepo.Update(ctx, input)

	if err != nil {
		t.Error("Could not update patient", err)
	}

	// Compare the patients to make sure they are the same
	if !reflect.DeepEqual(patient, expectedPatient) {
		t.Errorf("Cound not update original patient\nExpected:\n\t%v\ngot\n\t%v", expectedPatient, patient)
	}

	return *patient
}

func testSuspendPatient(
	ctx context.Context,
	patRepo PatientRepository,
	id string,
	t *testing.T,
) domain.Patient {

	patient, err := patRepo.Suspend(ctx, id)

	if err != nil {
		t.Error("Could not suspend patient", err)
	}

	// Compare the patients to make sure they are the same
	if patient.Status != domain.Suspended {
		t.Fatal("Patient was not suspended", err)
	}

	return *patient
}

func testActivatePatient(
	ctx context.Context,
	patRepo PatientRepository,
	id string,
	t *testing.T,
) domain.Patient {

	patient, err := patRepo.Activate(ctx, id)

	if err != nil {
		t.Error("Could not activate patient", err)
	}

	// Compare the patients to make sure they are the same
	if patient.Status != domain.Active {
		t.Fatal("Patient was not activated", err)
	}

	return *patient
}
