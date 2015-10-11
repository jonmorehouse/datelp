package datelp

import (
	"io"
	"bufio"
	"fmt"
	"errors"
)

type Input interface {
	Parse(io.Reader) error

	FetchRange(uint, uint, uint) ([]string, error)
	Fetch(uint) (string, error)
	Move(int) (uint, error)
	Reset()
}

type TextInput struct {
	data []string
	index uint
	size uint
}

func NewTextInput() *TextInput {
	return &TextInput{
		data: make([]string, 0),
		index: 0,
		size: 0,
	}
}

func (t *TextInput) Parse(input io.Reader) (error) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		word = StripWord(word)

		t.data = append(t.data, word)

		t.size++
	}

	return nil
}

func (t *TextInput) Reset() {
	t.index = 0
}

func (t *TextInput) Move(increment int) (uint, error) {
	t.index = uint(int(t.index) + increment)

	if t.index < 0 || t.index == t.size {
		t.index = 0
		return t.index, errors.New(fmt.Sprintf("Reached end of dataset"))
	}

	return t.index, nil
}

func (t *TextInput) FetchRange(index uint, leftOffset uint, rightOffset uint) ([]string, error) {
	start := int(index - leftOffset)
	end := index + rightOffset + 1

	if start < 0 || end > t.size {
		return nil, errors.New(fmt.Sprintf("Attempted to fetch nonexistent range: [%d:%d]", start, end))
	}

	return t.data[start:end], nil
}

func (t *TextInput) Fetch(index uint) (string, error) {
	if index < 0 || index > t.size {
		return "", errors.New(fmt.Sprintf("Out of range"))
	}

	return t.data[index], nil
}
