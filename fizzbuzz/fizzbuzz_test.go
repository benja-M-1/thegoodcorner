package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"}

	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	fizzbuzz := Replace(n, 3, "fizz", 5, "buzz")

	assert.Equal(t, expected, fizzbuzz, "Should replace 3 by fizz and 5 by buzz")
}
