package transaction

import (
	"fmt"
	"net/http"

	"github.com/leogsouza/expenses-tracking/server/internal/util/responses"

	"github.com/go-chi/chi"
)

type Router interface {
	Routes() chi.Router
}

type handler struct {
	service Service
}

func NewHandler(serv Service) Router {
	return &handler{serv}
}

func (h *handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.GetAll)
	return r
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	out, err := h.service.FindAll()
	if err != nil {
		responses.RespondError(w, fmt.Errorf("could not retrieve transactions: %v", err), http.StatusInternalServerError)
		return
	}
	responses.RespondOK(w, out)
}
