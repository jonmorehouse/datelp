package datelp

type InputText interface {
}

type TagClassifier interface {

}

type Classifier struct {
	input *InputText
	tags []*Tag
}

func NewClassifier(input *InputText) *Classifier {
	return &Classifier{
		input: input,
	}
}

func (c *Classifier) BuildTags() (error) {
	// iterate through all of the words in  the input text object
	// for each word check to see if its actually a tag
}





