package responses

import (
	"encoding/json"
	"net/http"
)

func RespondOK(w http.ResponseWriter, v interface{}) {
	respond(w, v, http.StatusOK)
}

func RespondError(w http.ResponseWriter, err error, statusCode int) {
	response := &ErrorResponse{statusCode, err.Error()}
	respond(w, response, statusCode)
}

func respond(w http.ResponseWriter, v interface{}, statusCode int) {
	response, _ := json.Marshal(v)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
