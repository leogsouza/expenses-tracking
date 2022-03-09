package account

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/leogsouza/expenses-tracking/backend/internal/entity"

	"github.com/leogsouza/expenses-tracking/backend/internal/util/responses"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// GetURLParam is just an alias for chi.URLParam function
var GetURLParam = chi.URLParam

// Router is an interface that wraps the Routes methods which returns an chi.Router that contains all routes from this package
type Router interface {
	Routes() chi.Router
	GetAll(w http.ResponseWriter, r *http.Request)
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
		r.Use(h.AccountCtx)
		r.Get("/", h.Get)
		r.Put("/", h.Update)
	})
	r.Post("/", h.Save)
	return r
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	out, err := h.service.FindAll()
	if err != nil {
		responses.RespondError(w, r, fmt.Errorf("could not retrieve accounts: %v", err), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, out)

}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	out := r.Context().Value("account").(*entity.Account)
	render.JSON(w, r, out)
}

func (h *handler) AccountCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var account entity.Account
		var err error

		if accountID := GetURLParam(r, "id"); accountID != "" {
			account, err = h.service.Find(entity.ID(accountID))
		} else {
			responses.RespondError(w, r, fmt.Errorf("account not found"), http.StatusNotFound)
			return
		}

		if err != nil {
			responses.RespondError(w, r, fmt.Errorf("could not retrieve a account: %v", err), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "account", &account)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type accountInput struct {
	Name string `json:"name"`
}

func (h *handler) Save(w http.ResponseWriter, r *http.Request) {
	var in accountInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		responses.RespondError(w, r, fmt.Errorf("could not read the account body: %v", err), http.StatusBadRequest)
		return
	}
	createdAt := time.Now().UTC()

	account := &entity.Account{
		ID:        entity.GenerateID(),
		Name:      in.Name,
		CreatedAt: createdAt,
	}
	out, err := h.service.Store(account)
	if err != nil {
		responses.RespondError(w, r, fmt.Errorf("could not save the account: %v", err), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, out)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	account := r.Context().Value("account").(*entity.Account)
	var in accountInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		responses.RespondError(w, r, fmt.Errorf("could not read the account body: %v", err), http.StatusBadRequest)
		return
	}

	account.Name = in.Name

	if err := h.service.Update(account); err != nil {
		responses.RespondError(w, r, fmt.Errorf("could not update the account: %v", err), http.StatusInternalServerError)
		return
	}

	out, err := h.service.Find(entity.ID(account.ID))
	if err != nil {
		responses.RespondError(w, r, fmt.Errorf("could not retrieve a account: %v", err), http.StatusNotFound)
		return
	}

	render.JSON(w, r, out)
}
