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
	h := FizzBuzzHandler{c}

	return h
}

func (h *FizzBuzzHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if http.MethodPost != r.Method {
		http.Error(w, "/fizzbuzz is only accessible with POST method.", http.StatusMethodNotAllowed)
		return
	}

	var f FizzBuzz
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&f)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.container.DB.CreateRequest(&f.Request)

	n := fizzbuzz.Replace(listGenerator(f), f.Request)

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
