package datelp

import (
	"strings"
	"regexp"
	"strconv"
	"errors"
	"fmt"
	"time"
)

func StripWord(word string) string {
	reg := regexp.MustCompile("[A-Za-z0-9/-]+")

	stripped := reg.FindString(word)
	lower := strings.ToLower(stripped)

	return lower
}

func WordToInteger(word string) (uint, error) {
	mapping := map[string]uint{
		"zero": 0,
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"ten": 10,
		"eleven": 11,
		"twelve": 12,
		"thirteen": 13,
		"fourteen": 14,
		"fifteen": 15,
		"sixteen": 16,
		"seventeen": 17,
		"eighteen": 18,
		"nineteen": 19,
		"twenty": 20,
		"thirty": 30,

		"first": 1,	
		"second": 2,
		"third": 3,
		"fourth": 4,
		"fifth": 5,
		"sixth": 6,
		"seventh": 7,
		"eighth": 8,
		"ninth": 9,
		"tenth": 10,
		"eleventh": 11,
		"twelth": 12,
		"thirteenth": 13,
		"fourteenth": 14,
		"fifteenth": 15,
		"sixteenth": 16,
		"seventeenth": 17,
		"eighteenth": 18,
		"nineteenth": 19,
		"twentieth": 20,
		"thirtieth": 30,
	}

	value, exists := mapping[word]
	if exists {
		return value, nil
	}

	// pull out just the number, in case of something like 1st, 2nd 3rd etc
	regex := regexp.MustCompile("[0-9]+")
	number := regex.FindString(word)
	integer, err := strconv.ParseUint(number, 10, 32)

	if err != nil {
		return 0, err
	}
	
	return uint(integer), nil
}

func WordsToInteger(words []string) (int, error) {
	value := 0
	found := false

	for _, word := range words {
		subWords := strings.Split(word, "-")
		for _, subWord := range subWords {
			subWordValue, err := WordToInteger(subWord)

			if err != nil {
				continue
			}

			found = true
			value += int(subWordValue)
		}
	}

	if !found {
		return 0, errors.New("Unable to parse as integer")
	}

	return value, nil
}

func ToTimeDelta(words []string) (map[string]int, error) {
	mapping := map[string]*regexp.Regexp{
		"day": regexp.MustCompile("day.*"),
		"week": regexp.MustCompile("week.*"),
		"month": regexp.MustCompile("month.*"),
		"year": regexp.MustCompile("year.*"),
	}

	foundAnyInterval := false
	delta := make(map[string]int)
	leftBound := 0

	// find each interval, and then take the left side of the words array
	// and look for a multiplier number
	for index, word := range words {
		for interval, regex := range mapping {
			if !regex.MatchString(word) {
				continue
			}

			foundAnyInterval = true

			// fetch the integers to the left of this element. This
			// assumes that the preceding integer is always the
			// multiplier. For instance month 2 is not equal to 2
			// months
			value, err := WordsToInteger(words[leftBound:index])
			if err != nil {
				value = 1
			}

			delta[interval] = value
			leftBound = index
		}
	}

	if !foundAnyInterval {
		return nil, errors.New(fmt.Sprintf("Unable to find time delta"))
	}

	return delta, nil
}

func ParseWeekday(words []string) (time.Weekday, error) {
	weekdays := map[string]time.Weekday{
		"mon": time.Monday,
		"monday": time.Monday,
		"tues": time.Tuesday,
		"tuesday": time.Tuesday,
		"wed": time.Wednesday,
		"wednesday": time.Wednesday,
		"thurs": time.Thursday,
		"thursday": time.Thursday,
		"fri": time.Friday,
		"friday": time.Friday,
		"sat": time.Saturday,
		"saturday": time.Saturday,
		"sunday": time.Sunday,
		"sun": time.Sunday,
	}

	for _, word := range words {
		weekday, exists := weekdays[word]
		if !exists {
			continue
		}

		return weekday, nil
	}


	return time.Sunday, errors.New("Unable to parse weekday")
}

func ParseMonth(words []string) (time.Month, error) {
	mapping := map[string]time.Month {
		"january": time.January,
		"jan": time.January,
		"february": time.February,
		"feb": time.February,
		"march": time.March,
		"mar": time.March,
		"april": time.April,
		"apr": time.April,
		"may": time.May,
		"june": time.June,
		"july": time.July,
		"august": time.August,
		"aug": time.August,
		"september": time.September,
		"sep": time.September,
		"october": time.October,
		"oct": time.October,
		"november": time.November,
		"nov": time.November,
		"december": time.December,
		"dec": time.December,
	}

	for _, word := range words {
		month, exists := mapping[word]
		if !exists {
			continue
		}

		// month was found!
		return month, nil
	}

	return time.January, errors.New("Unable to parse month")
}

func ParseInterval(words []string) (string, error) {
	mapping := map[string]*regexp.Regexp{
		"day": regexp.MustCompile("day.*"),
		"week": regexp.MustCompile("week.*"),
		"month": regexp.MustCompile("month.*"),
		"year": regexp.MustCompile("year.*"),
	}

	for _, word := range words {
		for interval, regex := range mapping {
			if regex.MatchString(word) {
				return interval, nil
			}
		}
	}

	return "", errors.New("Unable to parse interval")
}


