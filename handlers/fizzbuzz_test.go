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

func TestFizzbuzzHandler(t *testing.T) {
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
			"int1=3&int2=5&str1=fizz&str2=buzz&limit=16",
			"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16",
		},
		{
			"int1=3&str1=fizz&str2=buzz&limit=11",
			"1,2,fizz,4,5,fizz,7,8,fizz,10,11",
		},
		{
			"int1=2&int2=4&str2=buzz&&limit=8",
			"1,,3,buzz,5,,7,buzz",
		},
		{
			"int1=2&int2=4&str1=fizz&&limit=8",
			"1,fizz,3,fizz,5,fizz,7,fizz",
		},
	}


	h := createFizzBuzzHandler()
	handler := http.HandlerFunc(h.Handle)

	for _, c := range cases {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/fizzbuzz?%v", c.input), new(bytes.Buffer))
		handler.ServeHTTP(rr, req)

		assert.Equal(http.StatusOK, rr.Code)
		assert.Equalf(c.expected, rr.Body.String(), "Fizzbuzz handler should return the list of integers with request %v", c.input)
	}
}
