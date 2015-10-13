package datelp

import (
	"time"
	"testing"
	"fmt"
)

func TestOffsetAgoParser(t *testing.T) {
	testCases := []struct{
		input []string
		duration time.Duration
		shouldErr bool
	}{ 
		{[]string{"two", "weeks", "ago"}, time.Hour*24*14, false,},
		{[]string{"a", "day", "ago"}, time.Hour*24, false,},
		{[]string{"this", "is", "bad"}, time.Hour, true,},
	}

	for _, tc := range testCases {
		tag := NewOffsetTag(tc.input, "ago")
		offset, err := tag.Parse()

		if err == nil && tc.shouldErr {
			t.Error("Did not err when expected")
			return
		}

		if err != nil && !tc.shouldErr {
			t.Error("Erred when not expected")
			return
		}

		// can't test the offset correctly if it should've erred
		if tc.shouldErr {
			continue
		}

		expectedTime := time.Now().Add(-tc.duration).Truncate(time.Hour)
		if expectedTime != (*offset).Truncate(time.Hour) {
			t.Error("Did not convert the correct time")
			return
		}
	}
}

func TestOffsetThisParser(t *testing.T) {
	testCases := []struct{
		input []string
		duration time.Duration
		shouldErr bool
	} {
		{[]string{"this", "year"}, time.Hour*24*365, false,},
		{[]string{"this", "month"}, time.Hour*24*30, false,},
		{[]string{"this", "week"}, time.Hour*24*7, false,},
		{[]string{"day"}, time.Hour*24, false,},
	}

	for _, tc := range testCases {
		tag := NewOffsetTag(tc.input, "this")
		offset, err := tag.Parse()
		
		// validate that error cases are handled correctly
		if err == nil && tc.shouldErr {
			t.Error("Did not error when an error was expected")
			return
		}

		if err != nil && !tc.shouldErr {
			t.Error("Erred when not expected")
			return
		}

		// pass if this case should error out!
		if tc.shouldErr {
			continue
		}

		// ensure that the offset is accurate at least to the hour
		if (*offset).Truncate(time.Hour) != time.Now().Truncate(tc.duration).Truncate(time.Hour) {
			t.Error("Did not create the correct offset")
			return
		}
	}
}


func TestOffsetNextParser(t *testing.T) {
	testCases := []struct{
		words []string
	} {
		{[]string{"next", "june"},},
		{[]string{"next", "october"},},
		{[]string{"next", "year"},},
		{[]string{"next", "week"},},
		{[]string{"next", "day"},},
		{[]string{"next", "month"},},
	}

	for _, tc := range testCases {
		tag := NewOffsetTag(tc.words, "next")
		offset, _ := tag.Parse()

		fmt.Println(offset, tc.words)



	}
}
