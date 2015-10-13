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

func TestWordToInteger(t *testing.T) {
	testCases := []struct{
		output uint
		input string
	}{
		{1, "1st",},
		{1, "first",},
		{2, "2nd",},
		{3, "3",},
		{30, "thirty",},
	}

	for _, tc := range testCases {
		val, _ := WordToInteger(tc.input)

		if val != tc.output {
			t.Error(fmt.Sprintf("ToInteger conversion failed. expected: %d actual: %d, input: %s", tc.output, val, tc.input))
			return
		}
	}
}

func TestWordsToInteger(t *testing.T) {
	testCases := []struct{
		input []string
		output int
		shouldErr bool
	}{ 
		{[]string{"twenty-one"}, 21, false,},
		{[]string{"twenty", "one"}, 21, false,},
		{[]string{"null"}, 0, true,},
	}

	for _, tc := range testCases {
		val, err := WordsToInteger(tc.input)

		if err == nil && tc.shouldErr {
			t.Error("Did not return an error when expected") 
			return
		}

		if err != nil && !tc.shouldErr {
			t.Error("Errored when not expected")
			return
		}

		if val != tc.output {
			t.Error("Wrong value returned")
			return
		}
	}
}

func TestToTimeDelta(t *testing.T) {
	testCases := []struct{
		input []string
		results []int
		shouldErr bool
	}{
		{[]string{"3", "days"}, []int{3, 0, 0}, false,},
		{[]string{"4", "years"}, []int{0, 0, 4}, false,},
		{[]string{"4", "months"}, []int{0, 4, 0}, false,},
		{[]string{"thirty", "days"}, []int{30, 0, 0}, false,},
		{[]string{"one", "days"}, []int{1, 0, 0}, false,},
		{[]string{"thirty two", "days"}, []int{32, 0, 0}, false,},
		{[]string{"thirty-two", "days"}, []int{32, 0, 0}, false,},
		{[]string{"three", "years"}, []int{0, 0, 3}, false,},
		{[]string{"a", "month"}, []int{0, 1, 0}, false,},
		{[]string{"a", "couple", "months"}, []int{0, 2, 0}, false,},
		{[]string{"null"}, []int{0, 0, 0}, true,},
		{[]string{"three", "days", "two", "months", "one", "Year"}, []int{3, 2, 1}, false,},
		{[]string{"three", "days", "and", "two", "months", "and", "one", "Year"}, []int{3, 2, 1}, false,},
	}

	for _, tc := range testCases {
		_, err := ToTimeDelta(tc.input)

		if (err == nil) && tc.shouldErr {
			t.Error("Did not return an error when expected")
			return
		}

		if (err != nil) && !tc.shouldErr {
			t.Error("Returned an error when not expected")
			return
		}


	}
}
