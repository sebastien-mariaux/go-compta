package api

import (
	"go-compta/helpers"
	"go-compta/models"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = SetupRouter()

func TestGetExpenses(t *testing.T) {
	w := helpers.PerformRequest(router, "GET", "/expenses")

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Expense
	helpers.GetJsonResponse(w, t, &response)

	assert.Equal(t, 3, len(response))
}

func TestGetRevenues(t *testing.T) {
	w := helpers.PerformRequest(router, "GET", "/revenues")

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Revenue
	helpers.GetJsonResponse(w, t, &response)

	assert.Equal(t, 3, len(response))
}


type VATResponse struct {
	NetAmount float64 `json:"netAmount"`
	GrossAmount float64 `json:"grossAmount"`
	Id string `json:"id"`
	Vat float64 `json:"vat"`
}
func TestGetExpenseVat(t *testing.T) {
	w := helpers.PerformRequest(router, "GET", "/expenses/1/vat")
	assert.Equal(t, http.StatusOK, w.Code)

	var response VATResponse
	assert.Equal(t, 20, response.Vat)
}

func TestGetExpenseVatNotFound(t *testing.T) {
	w := helpers.PerformRequest(router, "GET", "/expenses/invalid_id/vat")

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetNetPNL(t *testing.T) {
	w := helpers.PerformRequest(router, "GET", "/pnl/net")
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	assert.Equal(t, 4043.0, response["net_pnl"])
}

func TestGetPNL(t *testing.T) {
	w := helpers.PerformRequest(router, "GET", "/pnl")
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	assert.Equal(t, 4043.0, response["pnl"])
}