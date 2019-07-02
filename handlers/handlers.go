package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/benja-M-1/thegoodcorner/fizzbuzz"
	"net/http"
	"strings"
)

type FizzBuzz struct {
	Int1  int    `json:int1`
	Int2  int    `json:int1`
	Str1  string `json:str1`
	Str2  string `json:str1`
	Limit int    `json:limit`
}

func FizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
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

	f := fizzbuzz.Replace(listGenerator(payload), payload.Int1, payload.Str1, payload.Int2, payload.Str2)

	fmt.Fprintf(w, strings.Join(f, ","))
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