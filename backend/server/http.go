package server

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	// ErrTextBadRequest "Bad request"
	ErrTextBadRequest = "Bad request"
	// ErrTextInternalServerError "An error occurred"
	ErrTextInternalServerError = "An error occurred"
	// ErrTextUnauthorized "Unauthorized"
	ErrTextUnauthorized = "Unauthorized"
)

// SendJSONResponse marshals "obj", sets the Content-Type header to
// "application/json" and writes the JSON to the response writer.
//
// In case of an error, http.Error will be called and the error will be logged.
func SendJSONResponse(w http.ResponseWriter, obj interface{}) {

	jsn, err := json.Marshal(obj)
	if err != nil {
		log.Printf("json.Marshal error for obj %v: %v", obj, err)
		http.Error(w, ErrTextInternalServerError, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
}
