package handlers

import (
	"bytes"
	"github.com/benja-M-1/thegoodcorner/app"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func createFizzBuzzHandler() FizzBuzzHandler {
	container := &app.Container{DB: &mockDB{}}
	h := NewFizzBuzzHandler(container)

	return h
}

func TestFizzbuzzHandlerShouldAllowOnlyPost(t *testing.T) {
	assert := assert.New(t)

	rr := httptest.NewRecorder()

	fbH := createFizzBuzzHandler()
	handler := http.HandlerFunc(fbH.Handle)

	req, _ := http.NewRequest(http.MethodGet, "/fizzbuzz", new(bytes.Buffer))
	handler.ServeHTTP(rr, req)

	assert.NotEqual(http.StatusOK, rr.Code, "Fizzbuzz handler should not accept GET method")
}

func TestFizzbuzzHandlerShouldReturnABadRequest(t *testing.T) {
	assert := assert.New(t)

	cases := []string{
		`{
			"request": {
				int1: 1
			}
			"limit": 0
		}`,
		"",
	}
	rr := httptest.NewRecorder()

	fbH := createFizzBuzzHandler()
	handler := http.HandlerFunc(fbH.Handle)

	for _, c := range cases {
		req, _ := http.NewRequest(http.MethodPost, "/fizzbuzz", strings.NewReader(c))
		handler.ServeHTTP(rr, req)

		assert.Equalf(http.StatusBadRequest, rr.Code, "Fizzbuzz handler should not accept POST request with body %v", c)
	}
}

type successfulCases struct {
	input    string
	expected string
}

func TestFizzbuzzHandler(t *testing.T) {
	assert := assert.New(t)

	cases := []successfulCases{
		{
			`{}`,
			"",
		},
		{
			`{
				"request": {}
			}`,
			"",
		},
		{
			`{
				"limit": 0
			}`,
			"",
		},
		{
			`{
				"request": {
					"int1": 3,
					"int2": 5,
					"str1": "fizz",
					"str2": "buzz"
				},
				"limit": 16
			}`,
			"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16",
		},
	}

	rr := httptest.NewRecorder()

	fbH := createFizzBuzzHandler()
	handler := http.HandlerFunc(fbH.Handle)

	for _, c := range cases {
		req, _ := http.NewRequest(http.MethodPost, "/fizzbuzz", strings.NewReader(c.input))
		handler.ServeHTTP(rr, req)

		assert.Equal(http.StatusOK, rr.Code, "Fizzbuzz handler should accept POST request")

		assert.Equal(c.expected, rr.Body.String(), "Fizzbuzz handler should return the list of integers with 3 replaced by fizz and 5 replaced by buzz")

	}
}
