package datelp

import (
	"testing"
	"fmt"
)

func TestStripWord(t *testing.T) {
	testCases := []struct{
		input string
		output string
	} {
		{"THIS....", "this",},
	}

	for _, tc := range testCases {
		if word := StripWord(tc.input); word != tc.output {
			t.Error(fmt.Sprintf("StripWord failed. Expected %s -> %s ... Got: %s", tc.input, tc.output, word))
		}
	}
}

func TestToInteger(t *testing.T) {
	testCases := []struct{
		output uint
		input string
	}{
		{1, "1st",},
		{1, "first",},
		{2, "2nd",},
	}

	for _, tc := range testCases {
		val, _ := ToInteger(tc.input)

		if val != tc.output {
			t.Error(fmt.Sprintf("ToInteger conversion failed. expected: %d actual: %d, input: %s", tc.output, val, tc.input))
			return
		}
	}
}
