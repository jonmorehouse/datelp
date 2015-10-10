package datelp


// classifiers are responsible for one thing ... building tags and passing in the appropriate slice of words to tags
type Classifier interface {
	//func (c *Classifier) classify(Input)
}

// this takes in an input and will return a set of tags that correspond to things it found classifiable
type OffsetClassifier struct {

	// hash table for looking up tags based upon index
	tags map [int]*Tag
}

func (oc *OffsetClassifier) Classify(input Input) {
	// want to iterate through the input and do things like "fetch previous"
	// the idea is that this will look for words and if its a valid word will create a tag (and will pass in some extra words into it)
}


