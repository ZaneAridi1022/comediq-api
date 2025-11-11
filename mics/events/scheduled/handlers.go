package scheduled

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/comediq-api/validation"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var mic Mic
	var err = json.NewDecoder(r.Body).Decode(&mic)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if validation.Validate.Struct(mic) != nil {
		http.Error(w, "Invalid JSON schema", http.StatusBadRequest)
		return
	}

	_, err = Create(&mic)
	if err != nil {
		log.Printf("error creating mic: %v", err)
		http.Error(w, "Failed to create mic", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
