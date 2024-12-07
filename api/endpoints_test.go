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

func TestExpenseVat(t *testing.T) {
	w := helpers.PerformRequest(router, "GET", "/expenses/1/vat")

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Amounts
	helpers.GetJsonResponse(w, t, &response)

	assert.Equal(t, 20.0, response.ComputeVat())
}

func TestExpenseVatNotFound(t *testing.T) {
	w := helpers.PerformRequest(router, "GET", "/expenses/invalid_id/vat")

	assert.Equal(t, http.StatusNotFound, w.Code)
}
