package handlers

import (
	"errors"
	"fmt"
	"github.com/benja-M-1/thegoodcorner/app"
	"github.com/benja-M-1/thegoodcorner/fizzbuzz"
	"github.com/benja-M-1/thegoodcorner/models"
	"log"
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
	if http.MethodGet != r.Method {
		http.Error(w, "Only GET requests are allowed.", http.StatusMethodNotAllowed)
		return
	}


	keys := r.URL.Query()

	err := validate(keys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	int1, _ := strconv.Atoi(keys.Get("int1"))
	int2, _ := strconv.Atoi(keys.Get("int2"))
	limit, _ := strconv.Atoi(keys.Get("limit"))
	str1 := keys.Get("str1")
	str2 := keys.Get("str2")

	fizzbuzzRequest := models.FizzbuzzRequest{0, int1, int2, str1, str2}

	_, err = h.container.DB.CreateRequest(&fizzbuzzRequest)
	if err != nil {
		log.Fatal(err.Error())
	}

	n := fizzbuzz.Replace(listGenerator(limit), fizzbuzzRequest)

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

// Validates that the values contained in the query of an url contains the required keys
func validate(v url.Values) error {
	keys := []string{"int1", "int2", "str1", "str2", "limit"}
	missingParams := []string{}

	for _, key := range keys {
		if _, ok := v[key]; !ok {
			missingParams = append(missingParams, key)
		}
	}

	if len(missingParams) > 0 {
		return errors.New(fmt.Sprintf("Missing '%v' parameter(s)", strings.Join(missingParams, ", ")))
	}

	return nil
}