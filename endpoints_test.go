package main

import (
    "go-compta/models"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = SetupRouter()

func TestGetExpenses(t *testing.T) {
    w := performRequest(router, "GET", "/expenses")

    assert.Equal(t, http.StatusOK, w.Code)

    var response []models.Expense
    getJsonResponse(w, t, &response)

    assert.Equal(t, 3, len(response))
}

func TestGetRevenues(t *testing.T) {
    w := performRequest(router, "GET", "/revenues")

    assert.Equal(t, http.StatusOK, w.Code)

    var response []models.Revenue
    getJsonResponse(w, t, &response)

    assert.Equal(t, 3, len(response))
}

func TestExpenseVat(t *testing.T) {
    w := performRequest(router, "GET", "/expenses/1/vat")

    assert.Equal(t, http.StatusOK, w.Code)

    var response models.Amounts
    getJsonResponse(w, t, &response)

    assert.Equal(t, 20.0, response.ComputeVat())
}

func TestExpenseVatNotFound(t *testing.T) {
    w := performRequest(router, "GET", "/expenses/10/vat")

    assert.Equal(t, http.StatusNotFound, w.Code)
}