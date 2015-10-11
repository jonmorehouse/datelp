package datelp

import (
	"time"
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

	return nil, nil
}

func (oc *OffsetTag) ParseAgo() (*time.Time, error) {

	return nil, nil
}

func (oc *OffsetTag) ParseBefore() (*time.Time, error) {

	return nil, nil
}

