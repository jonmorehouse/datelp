package datelp

import (
	"strings"
	"regexp"
	"strconv"
	"errors"
	"fmt"
)

func StripWord(word string) string {
	reg := regexp.MustCompile("[A-Za-z0-9/-]+")

	stripped := reg.FindString(word)
	lower := strings.ToLower(stripped)

	return lower
}

func ToInteger(word string) (uint, error) {
	mapping := map[string]uint{
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
		"twenty": 20,
		"thirty": 30,
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

func ToTimeDelta(words []string) (map[string]int, error) {
	mapping := map[string]*regexp.Regexp{
		"day": regexp.MustCompile("day.*"),
		"week": regexp.MustCompile("week.*"),
		"month": regexp.MustCompile("month.*"),
		"year": regexp.MustCompile("year.*"),
	}

	var value int
	interval := make(map[string]int, 3)

	for _, word := range words {
		val, err := ToInteger(word)
		if err != nil {
			value = int(val)
			continue
		}

		//https://golang.org/src/time/time.go?s=19687:19746#L645
		for key, regex := range mapping {
			str := regex.FindString(word)
			if str != "" {
				interval[key] = value
				return interval, nil
			}
		}
	}

	return nil, errors.New(fmt.Sprintf("Unable to find time delta"))
}


