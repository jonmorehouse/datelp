package datelp

import (
	"errors"
	"time"
)

func MaxInt(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func ConstantToWeekday(input int) (time.Weekday, error) {
	switch {
	case input == WEEKDAY_SUNDAY:
		return time.Sunday, nil
	case input == WEEKDAY_MONDAY:
		return time.Monday, nil
	case input == WEEKDAY_TUESDAY:
		return time.Tuesday, nil
	case input == WEEKDAY_WEDNESDAY:
		return time.Wednesday, nil
	case input == WEEKDAY_THURSDAY:
		return time.Thursday, nil
	case input == WEEKDAY_FRIDAY:
		return time.Friday, nil
	case input == WEEKDAY_SATURDAY:
		return time.Saturday, nil
	}

	return time.Sunday, errors.New("")
}

func ConstantToMonth(input int) (time.Month, error) {
	switch {
	case input == MONTH_JANUARY:
		return time.January, nil
	case input == MONTH_FEBRUARY:
		return time.February, nil
	case input == MONTH_MARCH:
		return time.March, nil
	case input == MONTH_APRIL:
		return time.April, nil
	case input == MONTH_MAY:
		return time.May, nil
	case input == MONTH_JUNE:
		return time.June, nil
	case input == MONTH_JULY:
		return time.July, nil
	case input == MONTH_AUGUST:
		return time.August, nil
	case input == MONTH_SEPTEMBER:
		return time.September, nil
	case input == MONTH_OCTOBER:
		return time.October, nil
	case input == MONTH_NOVEMBER:
		return time.November, nil
	case input == MONTH_DECEMBER:
		return time.December, nil
	}

	return time.January, errors.New("Invalid month")
}
