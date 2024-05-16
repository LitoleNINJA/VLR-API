package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type resJSON struct {
	Matches interface{} `json:"matches"`
	Count   int         `json:"count"`
}

func TestGetMatches(t *testing.T) {
	router := gin.Default()
	router.GET("/matches", getLiveMatches)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/matches", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "count", "Expected count in response")
	assert.Contains(t, w.Body.String(), "matches", "Expected matches in response")

	var res resJSON
	err := json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Greater(t, res.Count, 0, "Expected Count to be greater than 0")
}

func TestGetCompletedMatches(t *testing.T) {
	router := gin.Default()
	router.GET("/matches", getCompletedMatches)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/matches?status=completed", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "count", "Expected count in response")
	assert.Contains(t, w.Body.String(), "matches", "Expected matches in response")

	var res resJSON
	err := json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Greater(t, res.Count, 0, "Expected Count to be greater than 0")
}
