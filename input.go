package datelp

import (
	"io"
	"bufio"
	"fmt"
	"errors"
)

type Input interface {
	Parse(io.Reader) error

	Fetch(int, int, int) ([]string, error)
	Move(int)
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
		t.data = append(t.data, word)

		t.size++
	}

	return nil
}

func (t *TextInput) Reset() {
	t.index = 0
}

func (t *TextInput) Move(increment int) {
	newIndex := uint(int(t.index) + increment)

	if newIndex < 0 {
		newIndex = 0
	}

	t.index = newIndex
}

func (t *TextInput) Fetch(index uint, leftOffset uint, rightOffset uint) ([]string, error) {
	start := int(index - leftOffset)
	end := index + rightOffset + 1

	if start < 0 || end > t.size {
		return nil, errors.New(fmt.Sprintf("Attempted to fetch nonexistent range: [%d:%d]", start, end))
	}

	return t.data[start:end], nil
}

