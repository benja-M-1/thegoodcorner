package handlers

import (
	"errors"
	"fmt"
	"github.com/benja-M-1/thegoodcorner/app"
	"github.com/benja-M-1/thegoodcorner/fizzbuzz"
	"github.com/benja-M-1/thegoodcorner/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type FizzBuzzHandler struct {
	container *app.Container
}

func NewFizzBuzzHandler(c *app.Container) FizzBuzzHandler {
	h := FizzBuzzHandler{c}

	return h
}

func (h *FizzBuzzHandler) Handle(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	err := validate(keys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	int1, _ := strconv.Atoi(keys.Get("int1"))
	int2, _ := strconv.Atoi(keys.Get("int2"))
	limit, _ := strconv.Atoi(keys.Get("limit"))
	str1 := keys.Get("str1")
	str2 := keys.Get("str2")

	req := models.Request{0, int1, int2, str1, str2}

	h.container.DB.CreateRequest(&req)

	n := fizzbuzz.Replace(listGenerator(limit), req)

	fmt.Fprintf(w, strings.Join(n, ","))
}

func listGenerator(limit int) []int {
	start := 1
	l := make([]int, limit)

	for i := range l {
		l[i] = start
		start += 1
	}

	return l
}

func validate(v url.Values) error {
	keys := []string{"int1", "int2", "str1", "str2", "limit"}

	for _, key := range keys {
		if _, ok := v[key]; !ok {
			return errors.New(fmt.Sprintf("Missing '%v' parameter", key))
		}
	}

	return nil
}