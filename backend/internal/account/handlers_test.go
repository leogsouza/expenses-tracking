package account

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/leogsouza/expenses-tracking/backend/internal/entity"
	mocks "github.com/leogsouza/expenses-tracking/backend/mocks/account"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	t.Run("GetAll - Sucess", func(t *testing.T) {
		mockService := new(mocks.Service)

		mockService.On("FindAll").Return(getAccounts(), nil)

		rr := httptest.NewRecorder()

		h := NewHandler(mockService)

		request, err := http.NewRequest("GET", "accounts", nil)
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

}

func getAccount() entity.Account {
	return entity.Account{ID: "1uBp2CH2furqtZoM0lgWqcu9WRE", Name: "Wallet"}
}

func getAccounts() []entity.Account {
	return []entity.Account{getAccount()}
}
