package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type EntryLog struct {
	Data string `json:"data"`
}

func entrylog(router chi.Router) {
	router.Post("/", writeEntryLog)
}

func writeEntryLog(w http.ResponseWriter, r *http.Request) {
	var entryLog EntryLog
	json.NewDecoder(r.Body).Decode(&entryLog)
	log.Printf("Received: %s", entryLog.Data)
}
