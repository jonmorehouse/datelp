package datelp

import (
	"time"
	"errors"
)

/*
Tags correspond to known paths that the application looks for a date within.

The OffsetClassifier applies an offset to some sort of date. For instance this would
take something like "next,this,last,ago" and checks for a date.

AdverbTag looks for adverb descriptors such as `tomorrow`, `yesterday` `today`

DayTag looks for noun descriptors such as `sunday`, `june first`. Note, last
`sunday` would be classified as an adverb and thus would not be applicable here.

RegexDayTag

TimestampTag

*/
type Tag interface {
	Parse() (*time.Time, error)
}

type OffsetTag struct {
	offsetType string // ago, this, next, last, before
	words []string
}

func NewOffsetTag(words []string, offsetType string) *OffsetTag {
	return &OffsetTag{
		words: words,
		offsetType: offsetType,
	}
}

func (oc *OffsetTag) Parse() (*time.Time, error) {
	var parser func() (*time.Time, error)

	// switch based upon which parser this particular tag needs
	switch {
	case "ago" == oc.offsetType:
		parser = oc.ParseAgo
	case "before" == oc.offsetType:
		parser = oc.ParseBefore
	case "this" == oc.offsetType:
		parser = oc.ParseThis
	case "next" == oc.offsetType:
		parser = oc.ParseNext
	case "last" == oc.offsetType:
		parser = oc.ParseLast
	}

	return parser()
}

func (oc *OffsetTag) ParseAgo() (*time.Time, error) {
	timeDelta, err := ToTimeDelta(oc.words)
	if err != nil {
		return nil, err
	}

	current := time.Now()
	args := make([]int, 3, 3)

	for key, value := range timeDelta {
		switch {
		case "day" == key: 
			args[0] -= value
		case "week" == key:
			args[0] -= value*7
		case "month" == key:
			args[1] -= value
		case "year" == key:
			args[2] -= value

		}
	}

	offsetTime := current.AddDate(args[2], args[1], args[0])
	return &offsetTime, nil
}

func (oc *OffsetTag) ParseBefore() (*time.Time, error) {

	return nil, nil
}

func (oc *OffsetTag) ParseThis() (*time.Time, error) {
	offsetIndex := 0
	for index, word := range oc.words {
		if word != oc.offsetType {
			continue
		}

		offsetIndex = index
		break
	}

	interval, err := ParseInterval(oc.words[offsetIndex:])
	if err != nil {
		return nil, err
	}

	current := time.Now()
	var duration time.Duration

	switch {
	case "day" == interval:
		duration = time.Hour * 24
	case "week" == interval:
		duration = time.Hour*24*7
	case "month" == interval: 
		duration = time.Hour*24*30
	case "year" == interval:
		duration = time.Hour*24*365
	}

	if duration == 0 {
		return nil, errors.New("Unable to parse requested tag")
	}


	offsetTime := current.Truncate(duration)
	return &offsetTime, nil
}

func (oc *OffsetTag) ParseNext() (*time.Time, error) {
	offsetIndex := 0
	for index, word := range oc.words {
		if word != oc.offsetType {
			continue
		}

		offsetIndex = index
		break
	}

	intervalHandler := func(interval string) (*time.Time, error) {
		var args [3]int
		switch {
		case "day" == interval:
			args[0] = 1
		case "week" == interval:
			args[0] = 7
		case "month" == interval: 
			args[1] = 1
		case "year" == interval:
			args[2] = 1
		}

		offset := time.Now().AddDate(args[2], args[1], args[0])
		return &offset, nil
	}

	monthHandler := func(month time.Month) (*time.Time, error) {
		current := time.Now()
		months := 0

		if current.Month() == month {
			months = 12

		} else if current.Month() < month {
			// the month in question is only a few months in advance
			// for instance next july ... when the current month is june!
			months = int(month) - int(current.Month()) + 12
		} else {
			months = int(current.Month()) - int(month) + 12
		}

		offset := current.AddDate(0, months, 0)
		//truncatedOffset := offset.Truncate(24*time.Hour*30)

		return &offset, nil
	}

	weekdayHandler := func(weekday time.Weekday) (*time.Time, error) {

		return nil, nil
	}

	// lets try to find a weekday, a month or an interval
	interval, err := ParseInterval(oc.words[offsetIndex:])
	if err == nil {
		return intervalHandler(interval)
	}

	month, err := ParseMonth(oc.words[offsetIndex:])
	if err == nil {
		return monthHandler(month)
	}

	weekday, err := ParseWeekday(oc.words[offsetIndex:])
	if err == nil {
		return weekdayHandler(weekday)
	}

	return nil, errors.New("Unable to parse offset")
}

func (oc *OffsetTag) ParseLast() (*time.Time, error) {

	return nil, nil
}

