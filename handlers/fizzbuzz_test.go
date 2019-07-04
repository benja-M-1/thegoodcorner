package handlers

import (
	"bytes"
	"fmt"
	"github.com/benja-M-1/thegoodcorner/app"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)


func createFizzBuzzHandler() FizzBuzzHandler {
	container := &app.Container{DB: &mockDB{}}
	h := NewFizzBuzzHandler(container)

	return h
}

type successfulCases struct {
	input    string
	expected string
}

func TestFizzBuzzHandlerOnlyGET(t *testing.T) {
	assert := assert.New(t)

	h := createFizzBuzzHandler()
	handler := http.HandlerFunc(h.Handle)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/fizzbuzz", new(bytes.Buffer))
	handler.ServeHTTP(rr, req)

	assert.HTTPBodyContains(handler, http.MethodPost, "/fizzbuzz", nil, "Only GET requests are allowed.")
}

func TestFizzbuzzHandler(t *testing.T) {
	assert := assert.New(t)

	h := createFizzBuzzHandler()
	handler := http.HandlerFunc(h.Handle)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/fizzbuzz?int1=3&int2=5&str1=fizz&str2=buzz&limit=16", new(bytes.Buffer))
	handler.ServeHTTP(rr, req)

	assert.Equal(http.StatusOK, rr.Code)
	assert.Equal(
		"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16",
		rr.Body.String(),
		"Fizzbuzz handler should return the list of integers replaced by provided strings",
	)
}

func TestFizzbuzzHandlerErrors(t *testing.T) {
	assert := assert.New(t)

	cases := []successfulCases{
		{
			"",
			"",
		},
		{
			"limit=0",
			"",
		},
		{
			"int1=3&str1=fizz&str2=buzz&limit=11",
			"1,2,fizz,4,5,fizz,7,8,fizz,10,11",
		},
		{
			"int2=3&str1=fizz&str2=buzz&limit=11",
			"1,2,fizz,4,5,fizz,7,8,fizz,10,11",
		},
		{
			"int1=2&int2=4&str1=fizz&&limit=8",
			"1,fizz,3,fizz,5,fizz,7,fizz",
		},
		{
			"int1=2&int2=4&str2=buzz&&limit=8",
			"1,,3,buzz,5,,7,buzz",
		},
		{
			"int1=3&int2=11&str1=fizz&str2=buzz",
			"1,2,fizz,4,5,fizz,7,8,fizz,10,11",
		},
	}


	h := createFizzBuzzHandler()
	handler := http.HandlerFunc(h.Handle)

	for _, c := range cases {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/fizzbuzz?%v", c.input), new(bytes.Buffer))
		handler.ServeHTTP(rr, req)

		assert.Equal(http.StatusBadRequest, rr.Code)
	}
}
