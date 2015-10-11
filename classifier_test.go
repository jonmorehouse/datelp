package datelp

//import (
	//"fmt"
	//"strings"
	//"testing"
//)

//func NewInput(str string) Input {
	//reader := strings.NewReader(str)
	//input := NewTextInput()
	//input.Parse(reader)

	//return input
//}

//func TestOffsetClassifier(t *testing.T) {
	//// its import to note that the classification system only passes in
	//// "possible" matches where each possible match is a tag. Thus adding
	//// in the "this is not a date" is a valid option and _should_ be a tag
	//inputStr := "This week. This month. this.. year. This is not a date. 4 weeks ago. 2 days before june 5. next wednesday. next month. next June. last july. this"
	//input := NewInput(inputStr)
	//classifier := NewOffsetClassifier()
	//classifier.Classify(input)

	//testCases := []struct{
		//index uint
		//offsetType string
		//tagWords []string
	//}{
		//{0, "this", []string{"this", "week"},},
		//{2, "this", []string{"this", "month"},},
		//{4, "this", []string{"this", "year"},},
		//{6, "this", []string{"this", "is"},},
		//{13, "ago", []string{"4", "weeks", "ago"},},
		//{16, "before", []string{"2", "days", "before", "june", "5"},},
		//{19, "next", []string{"next", "wednesday"},},
		//{21, "next", []string{"next", "month"},},
		//{23, "next", []string{"next", "june"},},
		//{25, "last", []string{"last", "july"},},
	//}

	//for _, tc := range testCases {
		//tag, err := classifier.GetTag(tc.index)
		//if err != nil {
			//t.Error(fmt.Sprintf("Missed a classification: %s @ %d", tc.offsetType, tc.index))
			//return
		//}

		//// cast this back to a specialized offsetTag
		//offsetTag, ok := tag.(*OffsetTag)
		//if !ok || offsetTag == nil {
			//t.Error("Something is wrong with the offsetTag")
			//return
		//}

		//if offsetTag.offsetType != tc.offsetType {
			//t.Error(fmt.Sprintf("Incorrect OffsetType: %s. Expected: %s", offsetTag.offsetType, tc.offsetType))
			//return
		//}
	//}

	//// check to ensure that the tailing `this` doesn't get classified
	//if len(classifier.GetTags()) != len(testCases) {
		//t.Error("Incorrectly classified a tailing statement")
		//return
	//}
//}


