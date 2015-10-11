package datelp

import (
	"fmt"
	"errors"
)

// classifiers are responsible for one thing ... building tags and passing in the appropriate slice of words to tags
type Classifier interface {
	Classify (Input)
	GetTags() (map[uint]Tag)
	GetTag(uint) (Tag, error)
}

type OffsetClassificationRule struct {
	offsetType string
	left uint
	right uint
}

// this takes in an input and will return a set of tags that correspond to things it found classifiable
type OffsetClassifier struct {
	rules []OffsetClassificationRule

	// hash table for looking up tags based upon index
	tags map [uint]Tag
}

func NewOffsetClassifier() Classifier {
	rules := make([]OffsetClassificationRule, 5)
	rules[0] = OffsetClassificationRule{"ago", 2, 0}
	rules[1] = OffsetClassificationRule{"before", 2, 3}
	rules[2] = OffsetClassificationRule{"this", 0, 1}
	rules[3] = OffsetClassificationRule{"next", 0, 1}
	rules[4] = OffsetClassificationRule{"last", 0, 1}
	rules[5] = OffsetClassificationRule{"after", 2, 2}

	return &OffsetClassifier{
		tags: make(map[uint]Tag),
		rules: rules,
	}
}

func (oc *OffsetClassifier) Classify(input Input) {
	handler := func(index uint) int {
		word, err := input.Fetch(index)
		if err != nil {
			return 1
		}

		// look up which particular rule, if any that needs to be
		// created for this particular word that we are iterating
		// through
		for _, rule := range oc.rules {
			if rule.offsetType != word {
				continue
			}

			// if this is a match, go ahead and fetch the words
			// that are needed to create a tag for it
			words, err := input.FetchRange(index, rule.left, rule.right)
			if err != nil {
				return 1
			}

			oc.tags[index] = NewOffsetTag(words, rule.offsetType)
			return 1
		}

		return 1
	}
		
	increment := 0
	for {
		index, err := input.Move(increment)
		if err != nil {
			break
		}

		increment = handler(index)
	}
}

func (oc *OffsetClassifier) GetTags() map[uint]Tag {
	return oc.tags
}

func (oc *OffsetClassifier) GetTag(index uint) (Tag, error) {
	tag, exists := oc.tags[index]
	if !exists {
		return tag, errors.New(fmt.Sprintf("Requested tag doesn't exist: %d", index))
	}

	return tag, nil
}
