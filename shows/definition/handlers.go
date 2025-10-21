package definition

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var show Show
	var err = json.NewDecoder(r.Body).Decode(&show)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	_, err = Create(&show)
	if err != nil {
		log.Printf("error creating show: %v", err)
		http.Error(w, "Failed to create show", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
