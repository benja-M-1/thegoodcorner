package fizzbuzz

import (
	"fmt"
	"strconv"
)

func Replace(numbers []int, int1 int, str1 string, int2 int, str2 string) []string {
	replaced := make([]string, len(numbers))

	for index, num := range numbers {
		replacement := strconv.Itoa(num)

		if num%int1 == 0 {
			replacement = str1
		}

		if num%int2 == 0 {
			replacement = str2
		}

		if num%int1 == 0 && num%int2 == 0 {
			replacement = fmt.Sprintf("%v%v", str1, str2)
		}

		replaced[index] = replacement
	}

	return replaced
}
