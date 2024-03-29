package fizzbuzz

import (
	"github.com/benja-M-1/thegoodcorner/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type fizzbuzzTest struct {
	n        []int
	r        models.FizzbuzzRequest
	expected []string
}

var fizzbuzzCases = []fizzbuzzTest{
	{
		n: []int{},
		r: models.FizzbuzzRequest{
			Int1: 3,
			Int2: 5,
			Str1: "fizz",
			Str2: "buzz",
		},
		expected: []string{},
	},
	{
		n: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		r: models.FizzbuzzRequest{
			Int1: 3,
			Int2: 5,
			Str1: "fizz",
			Str2: "buzz",
		},
		expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"},
	},
	{
		n: []int{1, 2, 3, 4, 5, 6, 7, 8},
		r: models.FizzbuzzRequest{
			Int1: 2,
			Int2: 0,
			Str1: "",
			Str2: "",
		},
		expected: []string{"1", "", "3", "", "5", "", "7", ""},
	},
	{
		n: []int{1, 2, 3, 4, 5, 6, 7, 8},
		r: models.FizzbuzzRequest{
			Int1: 0,
			Int2: 2,
			Str1: "fizz",
			Str2: "",
		},
		expected: []string{"1", "", "3", "", "5", "", "7", ""},
	},
	{
		n: []int{1, 2, 3, 4, 5, 6, 7, 8},
		r: models.FizzbuzzRequest{
			Int1: 0,
			Int2: 2,
			Str1: "fizz",
			Str2: "buzz",
		},
		expected: []string{"1", "buzz", "3", "buzz", "5", "buzz", "7", "buzz"},
	},
}

func TestFizzbuzz(t *testing.T) {
	for _, test := range fizzbuzzCases {
		fizzbuzz := Replace(test.n, test.r)
		assert.Equalf(t, test.expected, fizzbuzz, "Should replace %v by %v and %v by %v", test.r.Int1, test.r.Str1, test.r.Int2, test.r.Str2)
	}
}

func BenchmarkReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range fizzbuzzCases {
			Replace(test.n, test.r)
		}
	}
}
