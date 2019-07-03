package fizzbuzz

import (
	"fmt"
	"github.com/benja-M-1/thegoodcorner/models"
	"strconv"
)

func Replace(numbers []int, request models.Request) []string {
	replaced := make([]string, len(numbers))

	for index, num := range numbers {
		replacement := strconv.Itoa(num)

		if num%request.Int1 == 0 {
			replacement = request.Str1
		}

		if num%request.Int2 == 0 {
			replacement = request.Str2
		}

		if num%request.Int1 == 0 && num%request.Int2 == 0 {
			replacement = fmt.Sprintf("%v%v", request.Str1, request.Str2)
		}

		replaced[index] = replacement
	}

	return replaced
}
