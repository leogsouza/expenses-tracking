package account

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/leogsouza/expenses-tracking/backend/internal/entity"
	mocks "github.com/leogsouza/expenses-tracking/backend/mocks/account"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	t.Run("GetAll - Success", func(t *testing.T) {
		mockService := new(mocks.Service)

		mockService.On("FindAll").Return(getAccounts(), nil)

		rr := httptest.NewRecorder()

		h := NewHandler(mockService)

		request, err := http.NewRequest("GET", "/api/accounts", nil)
		assert.NoError(t, err)

		handler := http.HandlerFunc(h.GetAll)

		handler.ServeHTTP(rr, request)

		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.Encode(getAccounts())

		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, buf.Bytes(), rr.Body.Bytes())
		mockService.AssertExpectations(t)
	})

	t.Run("GetAll - Error", func(t *testing.T) {
		mockService := new(mocks.Service)

		mockService.On("FindAll").Return(nil, errors.New("Could not find any accounts"))

		rr := httptest.NewRecorder()

		h := NewHandler(mockService)

		request, err := http.NewRequest("GET", "/api/accounts", nil)
		assert.NoError(t, err)

		handler := http.HandlerFunc(h.GetAll)

		handler.ServeHTTP(rr, request)
	})

}

func TestGet(t *testing.T) {
	t.Run("Get - Success", func(t *testing.T) {
		mockService := new(mocks.Service)
		account := getAccount()

		rr := httptest.NewRecorder()

		h := NewHandler(mockService)

		request := httptest.NewRequest("GET", fmt.Sprintf("/api/accounts/%s", account.ID.String()), nil)
		// Because the Get handler is getting the user through context
		// it's needed to set manually the context into request here
		var contextKey string = "account"

		request = request.WithContext(context.WithValue(request.Context(), contextKey, &account))

		handler := http.HandlerFunc(h.Get)

		handler.ServeHTTP(rr, request)

		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.Encode(getAccount())

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, buf.Bytes(), rr.Body.Bytes())
		mockService.AssertExpectations(t)

	})
}

func getAccount() entity.Account {
	return entity.Account{ID: "1uBp2CH2furqtZoM0lgWqcu9WRE", Name: "Wallet"}
}

func getAccounts() []entity.Account {
	return []entity.Account{getAccount()}
}
