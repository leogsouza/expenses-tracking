package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/leogsouza/expenses-tracking/server/internal/entity"

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
	r.Post("/", h.Save)
	return r
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	out, err := h.service.FindAll()
	if err != nil {
		responses.RespondError(w, fmt.Errorf("could not retrieve the users: %v", err), http.StatusInternalServerError)
		return
	}
	responses.RespondOK(w, out)
}

type userInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *handler) Save(w http.ResponseWriter, r *http.Request) {
	var in userInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		responses.RespondError(w, fmt.Errorf("could not read the request body: %v", err), http.StatusBadRequest)
		return
	}
	createdAt := time.Now().UTC()

	user := &entity.User{
		ID:        entity.GenerateID(),
		Name:      in.Name,
		Email:     in.Email,
		Password:  in.Password,
		CreatedAt: createdAt,
	}
	out, err := h.service.Store(user)
	if err != nil {
		responses.RespondError(w, fmt.Errorf("could not save the user: %v", err), http.StatusInternalServerError)
		return
	}
	responses.RespondOK(w, out)
}
