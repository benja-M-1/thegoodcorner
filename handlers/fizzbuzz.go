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

	var err error

	keys := r.URL.Query()

	err = validate(keys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	int1, err := strconv.Atoi(keys.Get("int1"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Cannot convert the value of the int1 parameter", http.StatusBadRequest)
		return
	}

	int2, _ := strconv.Atoi(keys.Get("int2"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Cannot convert the value of the int2 parameter", http.StatusBadRequest)
		return
	}

	limit, _ := strconv.Atoi(keys.Get("limit"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Cannot convert the value of the limit parameter", http.StatusBadRequest)
		return
	}

	str1 := keys.Get("str1")
	str2 := keys.Get("str2")

	fizzbuzzRequest := models.FizzbuzzRequest{0, int1, int2, str1, str2}

	_, err = h.container.DB.CreateRequest(&fizzbuzzRequest)
	// We don't return an http error here
	// If the save of the fizzbuzz request fails we can send a response to the user anyway
	if err != nil {
		log.Println(err.Error())
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