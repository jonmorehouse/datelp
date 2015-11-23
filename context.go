package datelp

import (
	"errors"
	"math"
	"time"
)

const (
	DIRECTION_CURRENT = 0
	DIRECTION_LEFT    = -1
	DIRECTION_RIGHT   = 1
)

const (
	INTERVAL_DAY = iota << 1
	INTERVAL_WEEKDAY
	INTERVAL_WEEK
	INTERVAL_MONTH
	INTERVAL_YEAR
	INTERVAL_CENTURY
)

const (
	SYNONYM_TODAY = iota << 1
	SYNONYM_TOMORROW
	SYNONYM_YESTERDAY
)

const (
	WEEKDAY_SUNDAY = iota << 1
	WEEKDAY_MONDAY
	WEEKDAY_TUESDAY
	WEEKDAY_WEDNESDAY
	WEEKDAY_THURSDAY
	WEEKDAY_FRIDAY
	WEEKDAY_SATURDAY
)

const (
	MONTH_JANUARY = iota << 1
	MONTH_FEBRUARY
	MONTH_MARCH
	MONTH_APRIL
	MONTH_MAY
	MONTH_JUNE
	MONTH_JULY
	MONTH_AUGUST
	MONTH_SEPTEMBER
	MONTH_OCTOBER
	MONTH_NOVEMBER
	MONTH_DECEMBER
)

type OffsetContext struct {
	interval  int // day, month, week year
	direction int // previous,after
	count     int // how large of an offset (in terms of quantity) eg: 2 weeks
	value     int // offset based upon a value instead of an interval. eg: next tuesday instead of next week
	truncate  int // useful for cases when the value is only supposed to be percieved as accurate to a certain interval
	size      int //number of successful elements that the offset found
}

// TODO update this to support passing in a start date
func (oc OffsetContext) Compile(origin time.Time) (time.Time, error) {
	// by default we assume that if any value is specified, then this is a
	// value Classification and treat it as such
	if oc.value >= 0 && (oc.interval == INTERVAL_WEEKDAY || oc.interval == INTERVAL_MONTH) {
		return oc.valueOffset(origin)
	}

	// if not value is created, then its a simple matter of building the
	// date from the correct parameters
	return oc.offset(origin)
}

func (oc OffsetContext) valueOffset(origin time.Time) (time.Time, error) {
	/*
	   A value offset is an offset that is based upon a specific value. For
	   instance a particular week day or month.

	   This compiler should be able to derive dates from classifications along the lines of:

	           Next monday
	           Next June
	           Last July
	           This Tuesday
	           Next Thursday

	   Its worth noting that something like `this month` or `this week` is not
	   considered a valueOffset, rather an interval offset.

	   1. get current value
	   2. get previous value and get next value
	   3. calculate start value
	   4. apply offset
	*/
	if oc.interval != INTERVAL_WEEKDAY && oc.interval != INTERVAL_MONTH {
		return time.Now(), errors.New("Unable to parse as value offset")
	}

	// figure out a delta which will correspond to the closest version of
	// this value. For instance, if its tuesday and we say next Friday this
	// should resolve to 3 days to the right
	index := 0
	start := 0
	threshold := 0
	value := oc.value

	if oc.interval == INTERVAL_WEEKDAY {
		tvalue, _ := ConstantToWeekday(value)
		value = int(tvalue)
		index = 2
		start = int(origin.Weekday())
		threshold = 2
	} else {
		tvalue, _ := ConstantToMonth(value)
		value = int(tvalue)
		index = 1
		start = int(origin.Month())
		threshold = 3
	}

	// calculate the delta by figuring out what value to
	// increment/decrement to get to the closest same value.
	delta := value - start
	if int(math.Abs(float64(delta))) > threshold {
		delta = value - start
	}

	// now we add or substract to the overall delta, by calculating how
	// much of an offset we need. For instance if this is next wednesday
	// and its currently monday then delta will +2 at this point. We will
	// then add +7 to the delta because we want _next_
	if oc.interval == INTERVAL_WEEKDAY {
		delta = delta + 7*oc.count*int(oc.direction)
	} else {
		delta = delta + 12*oc.count*int(oc.direction)
	}

	// number of years, months, days
	offsetArgs := [3]int{0, 0, 0}
	offsetArgs[index] = delta

	// the compiled date is the date that we compile once we've figured out
	// what delta needs to be applied
	compiledDate := origin.AddDate(offsetArgs[0], offsetArgs[1], offsetArgs[2])
	return compiledDate, nil
}

func (oc OffsetContext) offset(origin time.Time) (time.Time, error) {
	var days, months, years int

	switch oc.interval {
	case INTERVAL_WEEK:
		days = oc.count * 7
	case INTERVAL_DAY:
		days = oc.count
	case INTERVAL_MONTH:
		months = oc.count
	case INTERVAL_YEAR:
		years = oc.count
	}

	var direction int
	switch oc.direction {
	case DIRECTION_LEFT:
		direction = -1
	case DIRECTION_RIGHT:
		direction = 1
	case DIRECTION_CURRENT:
		direction = 0
	}

	d := origin.AddDate(direction*years, direction*months, direction*days)
	return d, nil
}

type DateContext struct {
	size     int // number of successful elements that belong to this context
	synonym  int // eg: TODAY/YESTERDAY/TOMORROW
	weekday  int // week day in particular
	month    int // eg MONTH_JUNE (constant)
	monthday int // 0-31 day
	year     int // year such as 2015
}

func (dc DateContext) Compile() (time.Time, error) {
	if dc.synonym >= 0 {
		return dc.compileSynonym()
	}

	return dc.compile()
}

func (dc DateContext) compileSynonym() (time.Time, error) {
	// TODO: figure out a way to make this more aligned with the work that
	// we're doing in the OffsetContext. It could be argued that this isn't
	// really a context since the synonyms only apply to a single day and
	// are extremely specialized.
	start := time.Now()
	dayOffset := 0

	switch {
	case dc.synonym == SYNONYM_YESTERDAY:
		dayOffset = -1
	case dc.synonym == SYNONYM_TODAY:
		dayOffset = 0
	case dc.synonym == SYNONYM_TOMORROW:
		dayOffset = 1
	default:
		return time.Now(), errors.New("Invalid day synonym")
	}

	return start.AddDate(0, 0, dayOffset), nil
}

func (dc DateContext) compileRelative() (time.Time, error) {
	// TODO: support value based dates, such as 3rd week of december or 3rd
	// Monday of the month.

	return time.Now(), nil
}

func (dc DateContext) isValid() bool {
	if dc.size == 0 {
		return false
	}

	// the only thing returned was a weekday, and this should use the offset parser instead!
	if dc.size == 1 && dc.weekday >= 0 {
		return false
	}

	return true
}

func (dc DateContext) compile() (time.Time, error) {
	var year, day int
	var month time.Month

	year = dc.year
	if year < 1 {
		year = time.Now().Year()
	}

	month, err := ConstantToMonth(dc.month)
	if err != nil {
		month = time.Now().Month()
	}

	day = dc.monthday
	if day < 1 {
		day = 1
	}

	compiledDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return compiledDate, nil
}
