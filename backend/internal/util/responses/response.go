package responses

import (
	"net/http"

	"github.com/go-chi/render"
)

func RespondOK(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.JSON(w, r, v)
}

func RespondError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {

	response := &ErrorResponse{
		Err:            err,
		HTTPStatusCode: statusCode,
		StatusText:     err.Error(),
		ErrorText:      err.Error(),
	}
	render.Render(w, r, response)
}
