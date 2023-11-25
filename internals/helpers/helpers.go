package helpers

import (
	"encoding/json"
	"net/http"
)

// a reponse writer function.
// It is used to send a json response
func WriteJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

type Message struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data.omitempty"`
}

func StatusInternalServerError(w http.ResponseWriter, message string) {
	WriteJson(w, http.StatusInternalServerError, Message{
		Message: message,
	})
}

func StatusUnauthorized(w http.ResponseWriter, message string) {
	WriteJson(w, http.StatusUnauthorized, Message{
		Message: message,
	})
}

func StatusBadRequest(w http.ResponseWriter, message string) {
	WriteJson(w, http.StatusBadRequest, Message{
		Message: message,
	})
}

func StatusNotFound(w http.ResponseWriter, message string) {
	WriteJson(w, http.StatusNotFound, Message{
		Message: message,
	})
}

func StatusOk(w http.ResponseWriter, message string) {
	WriteJson(w, http.StatusOK, Message{
		Message: message,
	})
}
