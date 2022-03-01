package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"sonic-server/db"

	"github.com/joho/godotenv"
	"github.com/steinfletcher/apitest"
)

var testHandler http.Handler
var testServer *httptest.Server

func TestMain(m *testing.M) {
	// load env vars for db connection while testing locally
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := db.Initialize()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Conn.Close()

	testHandler = NewHandler(db)
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
		Status(http.StatusOK).
		End()
}
