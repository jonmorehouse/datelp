package datelp

import (
        "testing"
        "strings"
)

func newWordIterator(input string) Iterator {
        wordReader := strings.NewReader(input)

        return NewWordIterator(wordReader)
}

func TestWordIteratorNext(t *testing.T) {
        i := newWordIterator("hello world")

        word, err := i.Next()
        if word != "world" || err != nil {
                t.Error("iterator returned wrong word")
                return
        }

        i.Move()
        word, err = i.Next()
        if word != "" || err == nil {
                t.Error("iterator should have returned error")
                return
        }
}

func TestWordIteratorPrev(t *testing.T) {
        i := newWordIterator("hello world")

        word, err := i.Prev()
        if err == nil || word != "" {
                t.Error("iterator should have returned error")
                return
        }

        i.Move()
        word, err = i.Prev()
        if err != nil || word != "hello" {
                t.Error("iterator should have succeeded")
                return
        }
}
