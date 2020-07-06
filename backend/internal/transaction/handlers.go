package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"context"

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
	r.Get("/type/{type}", h.GetAllByType)
	r.Get("/{id}", h.Get)
	r.Post("/", h.Save)
	r.Put("/{id}", h.Update)
	r.Get("/", h.GetAll)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(h.TransactionCtx)
		r.Get("/", h.Get)
		r.Put("/", h.Update)
	})
	r.Post("/", h.Save)
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

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	out := r.Context().Value("transaction").(*entity.Transaction)
	responses.RespondOK(w, out)
}

func (h *handler) TransactionCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var transaction entity.Transaction
		var err error

		if transactionID := GetURLParam(r, "id"); transactionID != "" {
			transaction, err = h.service.Find(entity.ID(transactionID))
		} else {
			responses.RespondError(w, fmt.Errorf("transaction not found"), http.StatusNotFound)
			return
		}

		if err != nil {
			responses.RespondError(w, fmt.Errorf("could not retrieve a transaction: %v", err), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "transaction", &transaction)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *handler) GetAllByType(w http.ResponseWriter, r *http.Request) {
	out, err := h.service.FindAllByType(GetURLParam(r, "type"))
	if err != nil {
		responses.RespondError(w, fmt.Errorf("could not retrieve transactions: %v", err), http.StatusInternalServerError)
		return
	}
	responses.RespondOK(w, out)
}

type transactionInput struct {
	UserID      entity.ID                `json:"user_id"`
	Type        entity.TypeTransaction   `json:"type"`
	AccountID   entity.ID                `json:"account_id"`
	CategoryID  entity.ID                `json:"category_id"`
	Description string                   `json:"description"`
	Amount      float64                  `json:"amount"`
	Date        time.Time                `json:"date"`
	Status      entity.StatusTransaction `json:"status"`
}

func (h *handler) Save(w http.ResponseWriter, r *http.Request) {
	var in transactionInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		responses.RespondError(w, fmt.Errorf("could not read the transaction body: %v", err), http.StatusBadRequest)
		return
	}
	createdAt := time.Now().UTC()

	transaction := &entity.Transaction{
		ID:          entity.GenerateID(),
		UserID:      in.UserID,
		Type:        in.Type,
		AccountID:   in.AccountID,
		CategoryID:  in.CategoryID,
		Description: in.Description,
		Amount:      in.Amount,
		Date:        in.Date,
		Status:      in.Status,
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt,
	}
	out, err := h.service.Store(transaction)
	if err != nil {
		responses.RespondError(w, fmt.Errorf("could not save the transaction: %v", err), http.StatusInternalServerError)
		return
	}
	responses.RespondOK(w, out)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	transaction := r.Context().Value("transaction").(*entity.Transaction)
	var in transactionInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		responses.RespondError(w, fmt.Errorf("could not read the account body: %v", err), http.StatusBadRequest)
		return
	}
	updatedAt := time.Now().UTC()
	transaction.Type = in.Type
	transaction.AccountID = in.AccountID
	transaction.CategoryID = in.CategoryID
	transaction.Description = in.Description
	transaction.Amount = in.Amount
	transaction.Date = in.Date
	transaction.Status = in.Status
	transaction.UpdatedAt = updatedAt

	if err := h.service.Update(transaction); err != nil {
		responses.RespondError(w, fmt.Errorf("could not update the transaction: %v", err), http.StatusInternalServerError)
		return
	}

	out, err := h.service.Find(entity.ID(transaction.ID))
	if err != nil {
		responses.RespondError(w, fmt.Errorf("could not retrieve a account: %v", err), http.StatusNotFound)
		return
	}
	responses.RespondOK(w, out)
}
