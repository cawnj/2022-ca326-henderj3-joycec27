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

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func entrylog(router chi.Router) {
	router.Post("/", writeEntryLog)
}

func writeEntryLog(w http.ResponseWriter, r *http.Request) {
	var entryLog EntryLog
	json.NewDecoder(r.Body).Decode(&entryLog)
	log.Printf("Received: %s", entryLog.Data)

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "log successful"})
}
