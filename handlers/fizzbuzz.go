package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/benja-M-1/thegoodcorner/app"
	"github.com/benja-M-1/thegoodcorner/fizzbuzz"
	"github.com/benja-M-1/thegoodcorner/models"
	"net/http"
	"strings"
)

type FizzBuzz struct {
	Request models.Request `json:request`
	Limit   int            `json:limit`
}

type FizzBuzzHandler struct {
	container *app.Container
}

func NewFizzBuzzHandler(c *app.Container) FizzBuzzHandler {
	f := FizzBuzzHandler{c}

	return f
}

func (f *FizzBuzzHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if http.MethodPost != r.Method {
		http.Error(w, "/fizzbuzz is only accessible with POST method.", http.StatusMethodNotAllowed)
		return
	}

	var payload FizzBuzz
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	payload.Request.Insert(f.container.DB)

	n := fizzbuzz.Replace(listGenerator(payload), payload.Request)

	fmt.Fprintf(w, strings.Join(n, ","))
}

func listGenerator(f FizzBuzz) []int {
	start := 1
	l := make([]int, f.Limit)

	for i := range l {
		l[i] = start
		start += 1
	}

	return l
}
