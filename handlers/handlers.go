package handlers

import (
	"fmt"
	"github.com/benja-M-1/thegoodcorner/fizzbuzz"
	"net/http"
	"strings"
)

func FizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	if http.MethodPost != r.Method {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	f := fizzbuzz.Replace(n, 3, "fizz", 5, "buzz")

	fmt.Fprintf(w, strings.Join(f, ","))
}

