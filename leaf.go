package datelp

import (
	"errors"
	"regexp"
	"strconv"
)

/*
   LeafClassifiers are responsible for classifying individual components of a
   branch based upon known words and some fine-tuned classify logic
*/
func ClassifyWordAsCommon(word string) bool {
	commonWords := []string{
		"and",
		"a",
		"of",
		"the",
		"in",
	}

	for _, commonWord := range commonWords {
		if word == commonWord {
			return true
		}
	}

	return false
}

func ClassifyWordAsInteger(word string) (int, bool, error) {
	// Classify a word as a potential integer. Its worth mentioning that if
	// this word is a strictly mapped word, such as "twenty" it could be a
	// part of stem. On the other hand, something like "23rd" isn't part of
	// a word chain.

	// TODO: parse words and figure out a way to look for word roots
	// for instance eight + enth | y could be 18|80
	mapping := map[string]int{
		"zero":      0,
		"one":       1,
		"two":       2,
		"three":     3,
		"four":      4,
		"five":      5,
		"six":       6,
		"seven":     7,
		"eight":     8,
		"nine":      9,
		"ten":       10,
		"eleven":    11,
		"twelve":    12,
		"thirteen":  13,
		"fourteen":  14,
		"fifteen":   15,
		"sixteen":   16,
		"seventeen": 17,
		"eighteen":  18,
		"nineteen":  19,

		"twenty":  20,
		"thirty":  30,
		"fourty":  40,
		"fifty":   50,
		"sixty":   60,
		"seventy": 70,
		"eighty":  80,
		"ninety":  90,

		"first":       1,
		"second":      2,
		"third":       3,
		"fourth":      4,
		"fifth":       5,
		"sixth":       6,
		"seventh":     7,
		"eighth":      8,
		"ninth":       9,
		"tenth":       10,
		"eleventh":    11,
		"twelth":      12,
		"thirteenth":  13,
		"fourteenth":  14,
		"fifteenth":   15,
		"sixteenth":   16,
		"seventeenth": 17,
		"eighteenth":  18,
		"nineteenth":  19,
		"twentieth":   20,
		"thirtieth":   30,

		"hundred":  100,
		"thousand": 1000,
	}

	value, exists := mapping[word]
	if exists {
		return value, true, nil
	}

	regex := regexp.MustCompile("-?[0-9]+")
	number := regex.FindString(word)
	integer, err := strconv.ParseInt(number, 10, 32)

	if err != nil {
		return 0, false, errors.New("Not an integer")
	}

	return int(integer), false, nil
}

func ClassifyAsIntegerStem(i Iterator) (int, int, error) {
	/*
	   Classify a chain of integers and return a single integer. A stem is a
	   series of related leaves that comprise a number. This attempts to
	   handle year-like numbers as well as standardized numbers.

	           `two thousand and two`
	           `thirty three`
	           `nineteen ninety five`
	           `three hundred two`
	*/
	count := 0
	values := make([]int, 0)

	for c := 0; c < 5; c++ {
		word, err := i.NextNth(c)
		// iterator is now out of range so it can proceed accordingly
		if err != nil {
			break
		}

		integer, stemPossible, err := ClassifyWordAsInteger(word)
		isCommon := ClassifyWordAsCommon(word)

		// if not an integer and not a common word we want to break
		if err != nil && !isCommon {
			break
		}

		count += 1
		if isCommon {
			continue
		}

		values = append(values, integer)

		// finally if this is not a possible stem then exit
		if !stemPossible {
			break
		}
	}

	if len(values) == 0 {
		return 0, count, errors.New("No integer found")
	}

	if len(values) == 1 {
		return values[0], count, nil
	}

	val := 0

	// this edge case is only to handle the weird date parsing of the teen
	// centuries. For instance dates such as 1993, 1854 etc need to know
	// that the leading teen corresponds to a thousands operation. In this
	// particular case this corresponds to the leading number having an
	// invisible "100" multiplier. This is only true if the second number is < 100
	if values[0] < 20 && values[0] > 10 && values[1] < 100 {
		val = 1e2 * values[0]
		values = values[1:]
	} else if values[1]%10 == 0 {
		val = values[0] * values[1]
		values = values[2:]
	}

	// for the remainder of the elements, we can just handle them as addition operations
	for _, value := range values {
		val += value
	}

	return val, count, nil
}

func ClassifyAsWeekday(i Iterator) (int, error) {
	mappings := map[int][]string{
		WEEKDAY_SUNDAY:    []string{"s", "sunday", "sun", "sundays"},
		WEEKDAY_MONDAY:    []string{"m", "monday", "mon", "mondays"},
		WEEKDAY_TUESDAY:   []string{"t", "tuesday", "tues", "tuesdays"},
		WEEKDAY_WEDNESDAY: []string{"t", "wednesday", "wed", "wednesdays"},
		WEEKDAY_THURSDAY:  []string{"th", "thursday", "thurs", "thu", "thursdays"},
		WEEKDAY_FRIDAY:    []string{"f", "friday", "fri", "fridays"},
		WEEKDAY_SATURDAY:  []string{"s", "saturday", "sat", "saturdays"},
	}

	for weekday, identifiers := range mappings {
		for _, word := range identifiers {
			if word == i.Current() {
				return weekday, nil
			}
		}
	}

	return 0, errors.New("Not a weekday")
}

func ClassifyAsDateday(i Iterator) (int, int, error) {
	integer, count, err := ClassifyAsIntegerStem(i)
	if err != nil {
		return 0, 0, err
	}

	if integer > 31 || integer < 1 {
		return 0, 0, errors.New("Integer out of daydate range")
	}

	return integer, count, nil
}

func ClassifyAsMonth(i Iterator) (int, error) {
	mappings := map[int][]string{
		MONTH_JANUARY:   []string{"jan", "january"},
		MONTH_FEBRUARY:  []string{"feb", "february"},
		MONTH_MARCH:     []string{"mar", "march"},
		MONTH_APRIL:     []string{"apr", "april"},
		MONTH_MAY:       []string{"may"},
		MONTH_JUNE:      []string{"june"},
		MONTH_JULY:      []string{"july"},
		MONTH_AUGUST:    []string{"aug", "august"},
		MONTH_SEPTEMBER: []string{"sep", "sept", "september"},
		MONTH_OCTOBER:   []string{"oct", "october"},
		MONTH_NOVEMBER:  []string{"nov", "november"},
		MONTH_DECEMBER:  []string{"dec", "december"},
	}

	for month, identifiers := range mappings {
		for _, identifier := range identifiers {
			if i.Current() == identifier {
				return month, nil
			}
		}
	}

	integer, _, err := ClassifyWordAsInteger(i.Current())
	if err != nil {
		return 0, errors.New("Unable to parse as a month")
	}

	if integer < 0 || integer > 12 {
		return 0, errors.New("Unable to parse as a month")
	}

	return integer, nil
}

func ClassifyAsYear(i Iterator) (int, int, error) {
	year, count, err := ClassifyAsIntegerStem(i)
	if err != nil {
		return 0, 0, err
	}

	if year < 1e3 {
		return 0, 0, errors.New("Unable to classify as year")
	}

	return year, count, nil
}

func ClassifyAsCommon(i Iterator) (string, error) {
	// for something like 3rd day of the month this would return for the "of the"
	// 2 weeks in the future
	if ClassifyWordAsCommon(i.Current()) {
		return i.Current(), nil
	}

	return "", errors.New("Not a common word")
}

func ClassifyAsDirection(i Iterator) (int, error) {
	mapping := map[int][]string{
		DIRECTION_CURRENT: []string{"this"},
		DIRECTION_LEFT:    []string{"before", "ago", "last"},
		DIRECTION_RIGHT:   []string{"next", "future", "from", "after"},
	}

	for interval, knownWords := range mapping {
		for _, word := range knownWords {
			if i.Current() == word {
				return interval, nil
			}
		}
	}

	return 0, errors.New("Not a direction")
}

func ClassifyAsInterval(i Iterator) (int, error) {
	mapping := map[int]*regexp.Regexp{
		INTERVAL_DAY:     regexp.MustCompile("^days?"),
		INTERVAL_WEEK:    regexp.MustCompile("^week*"),
		INTERVAL_MONTH:   regexp.MustCompile("^month*"),
		INTERVAL_YEAR:    regexp.MustCompile("^year*"),
		INTERVAL_CENTURY: regexp.MustCompile("^centur[y|ies]"),
	}

	for interval, re := range mapping {
		if re.MatchString(i.Current()) {
			return interval, nil
		}
	}

	return 0, errors.New("Not an interval")
}

func ClassifyAsDaySynonym(i Iterator) (int, error) {
	mapping := map[int]string{
		SYNONYM_YESTERDAY: "yesterday",
		SYNONYM_TODAY:     "today",
		SYNONYM_TOMORROW:  "tomorrow",
	}

	for synonym, keyword := range mapping {
		if i.Current() == keyword {
			return synonym, nil
		}
	}

	return 0, errors.New("Not a day synonym")
}
