package datelp

import (
	"testing"
	"errors"
	"strings"
)

func TestTextInputParse(t *testing.T) {
	reader := strings.NewReader("this is a block of text")
	input := &TextInput{}
	input.Parse(reader)

	if input.size != 6 {
		t.Error("Unable to parse the text correctly")
	}
}

func TestTextInputFetch(t *testing.T) {
	reader := strings.NewReader("this is a block of text")
	input := NewTextInput()
	input.Parse(reader)

	testCases := []struct {
		l, i, r uint
		expected []string
		err error
	}{
		{1, 0, 0, []string{"is"}, nil,},
		{0, 1, 0, nil, errors.New(""),},
		{4, 1, 1, []string{"block", "of", "text"}, nil,},
		{5, 3, 0, []string{"a", "block", "of", "text"}, nil,},
	}

	for _, tc := range testCases {
		words, err := input.Fetch(tc.l, tc.i, tc.r)
		
		if tc.err == nil && err != nil {
			t.Error("Returned an error when it should not have")
		} else if tc.err != nil && err == nil {
			t.Error("Did not return error when it should have")
		}

		for i := range tc.expected {
			if tc.expected[i] != words[i] {
				t.Error("Fetch did not return the correct words")
			}
		}
	}
}
