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

	TIMESTAMP_REGEX = `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`
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
		Body(`{
			"user_id": "test_user",
			"expo_token": "test_token"
		}`).
		Status(http.StatusOK).
		End()
}

func TestAllUsers(t *testing.T) {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(testHandler).
		Get("/users").
		Expect(t).
		Body(`{
			"users": [
				{
					"user_id": "test_user",
					"expo_token": "test_token"
				}
			]
		}`).
		Status(http.StatusOK).
		End()
}

func TestAllLocations(t *testing.T) {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(testHandler).
		Get("/locations").
		Expect(t).
		Body(`{
			"locations": [
				{
					"location_id": 1,
					"name": "The Spire",
					"coords": "0101000020E61000009A5B21ACC6AC4A40DC9C4A06800A19C0"
				},
				{
					"location_id": 2,
					"name": "DCU Nubar",
					"coords": "0101000020E6100000C36169E047B14A40067FBF982D0919C0"
				},
				{
					"location_id": 3,
					"name": "The Academy",
					"coords": "0101000020E6100000D6FCF84B8BAC4A40A0A52BD8460C19C0"
				}
			]
		}`).
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
				Equal(`name`, `DCU Nubar`).
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
				Matches(`entry_id`, `^\d+$`).
				Equal(`user_id`, `test_user`).
				Equal(`location_id`, float64(2)).
				Matches(`entry_time`, TIMESTAMP_REGEX).
				End(),
		).
		Status(http.StatusOK).
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
		Status(http.StatusOK).
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
				Matches(`status_text`, `^user '.+' registered successfully$`).
				End(),
		).
		Status(http.StatusOK).
		End()
}
