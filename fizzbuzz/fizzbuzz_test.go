package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type fizzbuzzTest struct {
	n        []int
	int1     int
	int2     int
	str1     string
	str2     string
	expected []string
}

var fizzbuzzCases = []fizzbuzzTest{
	{
		n:        []int{},
		int1:     3,
		int2:     5,
		str1:     "fizz",
		str2:     "buzz",
		expected: []string{},
	},
	{
		n:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		int1:     3,
		int2:     5,
		str1:     "fizz",
		str2:     "buzz",
		expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"},
	},
	{
		n:        []int{1, 2, 3, 4, 5, 6, 7, 8},
		int1:     2,
		int2:     4,
		str1:     "fizz",
		str2:     "buzz",
		expected: []string{"1", "fizz", "3", "fizzbuzz", "5", "fizz", "7", "fizzbuzz"},
	},
}

func TestFizzbuzz(t *testing.T) {
	for _, test := range fizzbuzzCases {
		fizzbuzz := Replace(test.n, test.int1, test.str1, test.int2, test.str2)
		assert.Equal(t, test.expected, fizzbuzz, "Should replace 3 by fizz and 5 by buzz")
	}
}

func BenchmarkReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range fizzbuzzCases {
			Replace(test.n, test.int1, test.str1, test.int2, test.str2)
		}
	}
}
