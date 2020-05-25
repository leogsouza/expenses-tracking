package responses

import (
	"encoding/json"
	"net/http"
)

func RespondOK(w http.ResponseWriter, v interface{}) {

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		RespondError(w, err, http.StatusInternalServerError)
		return
	}
	respond(w, v, http.StatusOK)
}

func RespondError(w http.ResponseWriter, err error, statusCode int) {
	response := &ErrorResponse{statusCode, err.Error()}
	respond(w, response, statusCode)
}

func respond(w http.ResponseWriter, v interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}
