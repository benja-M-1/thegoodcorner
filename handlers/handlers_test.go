package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFizzbuzzHandlerShouldReturnBadRequestWithGET(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest(http.MethodGet, "/fizzbuzz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzbuzzHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(http.StatusBadRequest, rr.Code, "Fizzbuzz handler should not accept GET method")
}

func TestFizzbuzzHandlerShouldReturnStringWithPOST(t *testing.T) {
	assert := assert.New(t)

	data := `{
		"int1": 3,
		"int2": 5,
		"str1": "fizz",
		"str2": "buzz",
		"limit": 16
	}`

	req, err := http.NewRequest(http.MethodPost, "/fizzbuzz", strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzbuzzHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(http.StatusOK, rr.Code, "Fizzbuzz handler should accept POST request")

	expected := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16"
	assert.Equal(expected, rr.Body.String(), "Fizzbuzz handler should return the list of integers with 3 replaced by fizz and 5 replaced by buzz")
}
