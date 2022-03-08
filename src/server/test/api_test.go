package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"sonic-server/db"
	"sonic-server/handler"

	"github.com/joho/godotenv"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

const (
	POSTGRES_HOST = "localhost"
	POSTGRES_DB   = "test"

	STRING_REGEX    = `^.+$`
	NUMBER_REGEX    = `^\d+$`
	TIMESTAMP_REGEX = `^$|^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`
)

var testHandler http.Handler
var testServer *httptest.Server

func TestMain(m *testing.M) {
	// load env vars for db connection while testing locally
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := db.Initialize(POSTGRES_HOST, POSTGRES_DB)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Conn.Close()

	testHandler = handler.NewHandler(db)
	testServer = httptest.NewServer(testHandler)
	defer testServer.Close()

	os.Exit(m.Run())
}

func TestHealth(t *testing.T) {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(testHandler).
		Get("/health").
		Expect(t).
		Body("Alive!").
		Status(http.StatusOK).
		End()
}

func TestSingleUser(t *testing.T) {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(testHandler).
		Get("/user").
		JSON(`{
			"user_id": "test_user"
		}`).
		Expect(t).
		Assert(
			jsonpath.Chain().
				Equal(`user_id`, `test_user`).
				Matches(`expo_token`, STRING_REGEX).
				End(),
		).
		Status(http.StatusOK).
		End()
}

func TestAllUsers(t *testing.T) {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(testHandler).
		Get("/users").
		Expect(t).
		Assert(
			jsonpath.Chain().
				Matches(`$.users[0].user_id`, STRING_REGEX).
				Matches(`$.users[0].expo_token`, STRING_REGEX).
				End(),
		).
		Status(http.StatusOK).
		End()
}

func TestAllLocations(t *testing.T) {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(testHandler).
		Get("/locations").
		Expect(t).
		Assert(
			jsonpath.Chain().
				Matches(`$.locations[0].location_id`, NUMBER_REGEX).
				Matches(`$.locations[0].name`, STRING_REGEX).
				Matches(`$.locations[0].coords`, STRING_REGEX).
				End(),
		).
		Status(http.StatusOK).
		End()
}

func TestLatestLocation(t *testing.T) {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(testHandler).
		Post("/latestlocation").
		JSON(`{
			"user_id": "test_user"
		}`).
		Expect(t).
		Assert(
			jsonpath.Chain().
				Present(`name`).
				Matches(`timestamp`, TIMESTAMP_REGEX).
				End(),
		).
		Status(http.StatusOK).
		End()
}

func TestEntryLog(t *testing.T) {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(testHandler).
		Post("/entrylog").
		JSON(`{
			"user_id": "test_user",
			"location_id": 2,
			"timestamp": "3000-01-01T00:00:00Z"
		}`).
		Expect(t).
		Assert(
			jsonpath.Chain().
				Matches(`entry_id`, NUMBER_REGEX).
				Equal(`user_id`, `test_user`).
				Equal(`location_id`, float64(2)).
				Matches(`entry_time`, TIMESTAMP_REGEX).
				Matches(`exit_time`, TIMESTAMP_REGEX).
				End(),
		).
		Status(http.StatusCreated).
		End()
}

func TestTrace(t *testing.T) {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(testHandler).
		Post("/trace").
		JSON(`{
			"user_id": "test_user"
		}`).
		Expect(t).
		Assert(
			jsonpath.Chain().
				Equal(`status_code`, float64(201)).
				Matches(`status_text`, `^notified \d+ close contact\(s\)$`).
				End(),
		).
		Status(http.StatusCreated).
		End()
}

func TestRegister(t *testing.T) {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(testHandler).
		Post("/register").
		JSON(`{
			"user_id": "test_user",
			"expo_token": "test_token"
		}`).
		Expect(t).
		Assert(
			jsonpath.Chain().
				Equal(`status_code`, float64(201)).
				Equal(`status_text`, `user 'test_user' registered successfully`).
				End(),
		).
		Status(http.StatusCreated).
		End()
}
