package datelp

/*
Tags correspond to known paths that the application looks for a date within.

The OffsetTag applies an offset to some sort of date. For instance this would
take something like "next,this,last,ago" and checks for a date.

AdverbTag looks for adverb descriptors such as `tomorrow`, `yesterday` `today`

DayTag looks for noun descriptors such as `sunday`, `june first`. Note, last
`sunday` would be classified as an adverb and thus would not be applicable here.

RegexDayTag

TimestampTag

*/
type Tag interface {
	// has a parse
}

type OffsetTag {
	index int
	classified bool
}

func (o *Offset) Classify(input * InputText) {

}





