package datelp

import (
        "time"
        "strings"
)

type Result struct {
        Size int
        Date time.Time
}

func Parse(input string) (time.Time, error) {
        classifier := NewClassifier()
        stringReader := strings.NewReader(input)
        iterator := NewWordIterator(stringReader)

        results, err := classifier.Parse(iterator)
        if err != nil {
                return time.Now(), err
        }

        return results.Date, nil
}
