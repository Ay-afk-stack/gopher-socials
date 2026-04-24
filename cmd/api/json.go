package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func WriteJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Success bool `json:"succes"`
		Error string `json:"error"`
	}

	return WriteJSON(w, status, &envelope{
		Success: false,
		Error: message,
		},
	)
}

func (app *application) jsonResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Success bool `json:"success"`
		Data any `json:"data"`
	}

	return WriteJSON(w, status, &envelope{
		Success: true,
		Data: data,
	})
}