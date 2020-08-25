package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/leogsouza/expenses-tracking/backend/internal/entity"

	"github.com/leogsouza/expenses-tracking/backend/internal/util/responses"

	"github.com/go-chi/chi"
)

// GetURLParam is just an alias for chi.URLParam function
var GetURLParam = chi.URLParam

// Router is an interface that wraps the Routes methods which returns an chi.Router that contains all routes from this package
type Router interface {
	Routes() chi.Router
}

type handler struct {
	service Service
}

// NewHandler returns a router
func NewHandler(serv Service) Router {
	return &handler{serv}
}

func (h *handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.GetAll)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(h.UserCtx)
		r.Get("/", h.Get)
		r.Put("/", h.Update)
	})
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

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	out, err := h.service.Find(entity.ID(GetURLParam(r, "id")))
	if err != nil {
		responses.RespondError(w, fmt.Errorf("could not retrieve an user: %v", err), http.StatusNotFound)
		return
	}
	responses.RespondOK(w, out)
}

func (h *handler) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user entity.User
		var err error

		if userID := GetURLParam(r, "id"); userID != "" {
			user, err = h.service.Find(entity.ID(userID))
		} else {
			responses.RespondError(w, fmt.Errorf("user not found"), http.StatusNotFound)
			return
		}

		if err != nil {
			responses.RespondError(w, fmt.Errorf("could not retrieve a user: %v", err), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "user", &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*entity.User)
	var in userInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		responses.RespondError(w, fmt.Errorf("could not read the body request: %v", err), http.StatusBadRequest)
		return
	}

	user.Name = in.Name
}
