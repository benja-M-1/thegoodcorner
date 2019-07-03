package handlers

import (
	"bytes"
	"github.com/benja-M-1/thegoodcorner/app"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createStatisticsHandler() StatisticsHandler {
	container := &app.Container{DB: &mockDB{}}
	h := NewStatisticsHandler(container)

	return h
}


func TestStatisticsHandler(t *testing.T) {
	assert := assert.New(t)

	sh := createStatisticsHandler()
	handler := http.HandlerFunc(sh.Handle)

	req, _ := http.NewRequest(http.MethodGet, "/statistics", new(bytes.Buffer))
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(http.StatusOK, rr.Code, "Statistics handler sent a 200 status code")
}
