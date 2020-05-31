package transaction

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
	Status      entity.StatusTransaction `json: "status"`
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
		ID:         entity.GenerateID(),
		UserID:     in.UserID,
		Type:       in.Type,
		AccountID:  in.AccountID,
		CategoryID: in.CategoryID,
		Amount:     in.Amount,
		Date:       in.Date,
		Status:     in.Status,
		CreatedAt:  createdAt,
		UpdatedAt:  createdAt,
	}
	out, err := h.service.Store(transaction)
	if err != nil {
		responses.RespondError(w, fmt.Errorf("could not save the transaction: %v", err), http.StatusInternalServerError)
		return
	}
	responses.RespondOK(w, out)
}
