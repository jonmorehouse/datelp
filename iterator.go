package datelp

import (
        "io"
        "bufio"
        "errors"
)

type Iterator interface {
        Current() string
        Next() (string, error)
        NextNth(int) (string, error)

        Prev() (string, error)
        PrevNth(int) (string, error)

        Move() error
        MoveN(int) error
        End() bool
}

type WordIterator struct {
        words []string
        index int
}

func NewWordIterator(input io.Reader) Iterator {
        scanner := bufio.NewScanner(input)
        scanner.Split(bufio.ScanWords)

        words := make([]string, 0)
        for scanner.Scan() {
                word := scanner.Text()
                words = append(words, word)
        }

        return &WordIterator{
                words: words,
                index: 0,
        }
}

func (i WordIterator) End() bool {
        return i.index + 1 >= len(i.words) 
}


func (i *WordIterator) Move() error {
        if i.index + 1 >= len(i.words) {
                return errors.New("out of range")
        }

        i.index += 1
        return nil
}

func (i *WordIterator) MoveN(n int) error {
        if i.index + n >= len(i.words) || i.index + n < 0 {
                return errors.New("out of range")
        }

        i.index += n
        return nil
}

func (i WordIterator) Current() string {
        return i.words[i.index]
}

func (i WordIterator) Next() (string, error) {
        if i.index + 1 >= len(i.words) {
                return "", errors.New("out of range")
        }

        return i.words[i.index+1], nil
}

func (i WordIterator) Prev() (string, error) {
        if i.index - 1 < 0 {
                return "", errors.New("out of range")
        }

        return i.words[i.index - 1], nil
}

func (i WordIterator) PrevNth(n int) (string, error) {
        if i.index - n < 0 {
                return "", errors.New("out of range")
        }

        return i.words[i.index - n], nil
}

func (i WordIterator) NextNth(n int) (string, error) {
        if i.index + n >= len(i.words) {
                return "", errors.New("out of range")
        }

        return i.words[i.index + n], nil
}

