package fizzbuzz

import (
	"fmt"
	"github.com/benja-M-1/thegoodcorner/models"
	"strconv"
)

// Replaces every multiples of a Request values by the corresponding string
func Replace(numbers []int, request models.FizzbuzzRequest) []string {
	replaced := make([]string, len(numbers))

	for index, num := range numbers {
		replacement := strconv.Itoa(num)

		if shouldReplace(request.Int1, num) {
			replacement = request.Str1
		}

		if shouldReplace(request.Int2, num) {
			replacement = request.Str2
		}

		if shouldReplace(request.Int1, num) && shouldReplace(request.Int2, num) {
			replacement = fmt.Sprintf("%v%v", request.Str1, request.Str2)
		}

		replaced[index] = replacement
	}

	return replaced
}

func shouldReplace(i int, num int) bool {
	return i > 0 && num%i == 0
}
