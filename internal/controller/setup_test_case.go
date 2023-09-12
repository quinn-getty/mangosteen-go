package controller

import (
	"mangosteen/config"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	q *queries.Queries
	w *httptest.ResponseRecorder
	r *gin.Engine
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	config.LoadConfig()
	database.Connect()
	w = httptest.NewRecorder()
	q = database.NewQuery()

	r = gin.Default()

	return func(t *testing.T) {
		database.Close()
	}
}
