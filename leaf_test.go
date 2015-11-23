package datelp

import (
	"testing"
)

func TestClassifyAsIntegerBranch(t *testing.T) {
	inputs := map[string]int{
		"two thousand and 15":         2015,
		"two thousand fifteen":        2015,
		"two thousand and thirty one": 2031,
		"fifteen hundred":             1500,
		"nineteen ninety 3":           1993,
		"thirty three":                33,
		"two hundred":                 200,
		"three hundred and 45":        345,
	}

	for input, expected := range inputs {
		i := newWordIterator(input)

		val, _, _ := ClassifyAsIntegerStem(i)
		if val != expected {
			t.Log(val, expected)
			t.Error("Invalid  parsed from input")
		}
	}
}

// TODO: write tests for other classifiers
