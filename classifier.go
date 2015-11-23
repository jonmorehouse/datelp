package datelp

import (
	"errors"
	"time"
)

type Classifier struct {
	offset *OffsetContext
	date   *DateContext
	start  time.Time
}

func NewClassifier() *Classifier {
	return &Classifier{}
}

func (c *Classifier) Parse(i Iterator) (*Result, error) {
	err := c.buildContexts(i)
	if err != nil || (c.offset.size == 0 && c.date.size == 0) {
		// neither context could be built, exit and emit an error
		return nil, errors.New("Unable to build any context. No date parseable")
	}

	result := &Result{Size: MaxInt(c.offset.size, c.date.size), Date: time.Now()}

	if c.date.isValid() {
		date, err := c.date.Compile()
		if err != nil {
			return nil, err
		}
		result.Date = date
	}

	date, err := c.offset.Compile(result.Date)
	if err != nil {
		return result, nil
	}
	result.Date = date

	return result, nil
}

func (c *Classifier) buildContexts(i Iterator) error {
	errs := 0
	successes := 0

	c.offset = &OffsetContext{
		interval: 0,
		count:    1,
		size:     0,
	}
	c.date = &DateContext{
		size:    0,
		weekday: -1,
		synonym: -1,
	}

	// its worth mentioning that this element loops through the element as
	// both a date and offset context. The goal here is to be able to
	// identify each component in isolation without interfering with one
	// another. When both an offset and date context are found, then we use
	// the date as the "starting" point for the offset.
	for {
		offsetCount, offsetErr := c.parseOffset(i)
		dateCount, dateErr := c.parseDate(i)

		if offsetErr != nil && dateErr != nil {
			errs += 1
		} else {
			successes += MaxInt(offsetCount, dateCount)
		}

		// if 4 errs in a row have happened or we are at the end of the
		// iterator, break because it can be assumed that nothing of
		// importance was actually found
		if err := i.MoveN(MaxInt(offsetCount, dateCount)); errs == 4 || err != nil {
			break
		}
	}

	if successes == 0 {
		return errors.New("Nothing found")
	}

	return nil
}

func (c *Classifier) parseOffset(i Iterator) (int, error) {
	if _, err := ClassifyAsCommon(i); err == nil {
		return 1, nil
	}

	if value, err := ClassifyAsDirection(i); err == nil {
		c.offset.direction = value
		c.offset.size += 1
		return 1, nil
	}

	if value, count, err := ClassifyAsIntegerStem(i); err == nil {
		c.offset.count = value
		c.offset.size += count
		return count, nil
	}

	if value, err := ClassifyAsInterval(i); err == nil {
		c.offset.interval = value
		c.offset.size += 1
		return 1, nil
	}

	if value, err := ClassifyAsWeekday(i); err == nil {
		c.offset.value = value
		c.offset.interval = INTERVAL_WEEKDAY
		c.offset.size += 1
		return 1, nil
	}

	if value, err := ClassifyAsMonth(i); err == nil {
		c.offset.value = value
		c.offset.interval = INTERVAL_MONTH
		c.offset.size += 1
		return 1, nil
	}

	return 1, errors.New("Unable to parse any valid pattern or value from iterator")
}

func (c *Classifier) parseDate(i Iterator) (int, error) {
	// this looks for arbitrary components of a date and attempts to parse
	// them together into a dateContext which can be compiled and used as the starting point
	if _, err := ClassifyAsCommon(i); err == nil {
		return 1, nil
	}

	if value, err := ClassifyAsDaySynonym(i); err == nil {
		c.date.synonym = value
		c.date.size += 1
		return 1, nil
	}

	if value, err := ClassifyAsWeekday(i); err == nil {
		c.date.weekday = value
		c.date.size += 1
		return 1, nil
	}

	if value, err := ClassifyAsMonth(i); err == nil {
		// here's an example of where this sort of thing breaks June
		// 1st 2015... the second time around this will pick up the 1
		// as a month :( Need to figure out a way to "partially"
		// classify things
		if c.date.month == 0 {
			c.date.month = value
			c.date.size += 1
			return 1, nil
		}
	}

	if value, count, err := ClassifyAsDateday(i); err == nil {
		c.date.monthday = value
		c.date.size += 1
		return count, nil
	}

	if value, count, err := ClassifyAsYear(i); err == nil {
		c.date.year = value
		c.date.size += 1
		return count, nil
	}

	return 1, errors.New("Unable to parse any valid pattern or value from iterator")
}
