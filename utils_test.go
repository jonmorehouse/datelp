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
