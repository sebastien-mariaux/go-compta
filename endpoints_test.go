package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to perform a GET request and return the response recorder
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
    req, _ := http.NewRequest(method, path, nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    return w
}

// Helper function to unmarshal JSON response
func unmarshalResponse(w *httptest.ResponseRecorder, v interface{}) error {
    return json.Unmarshal(w.Body.Bytes(), v)
}

func getJsonResponse(w *httptest.ResponseRecorder, t *testing.T, v interface{}) {
    err := unmarshalResponse(w, v)
    assert.NoError(t, err)
}

func TestGetExpenses(t *testing.T) {
    r := SetupRouter()
    w := performRequest(r, "GET", "/expenses")

    assert.Equal(t, http.StatusOK, w.Code)

    var response []expense
    getJsonResponse(w, t, &response)

    // Test we have 3 elements in response
    assert.Equal(t, 3, len(response))
}

func TestGetRevenues(t *testing.T) {
    r := SetupRouter()
    w := performRequest(r, "GET", "/revenues")

    assert.Equal(t, http.StatusOK, w.Code)

    var response []revenue
    getJsonResponse(w, t, &response)

    // Test we have 3 elements in response
    assert.Equal(t, 3, len(response))
}