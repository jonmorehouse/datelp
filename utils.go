package datelp

import (
	"strings"
	"regexp"
)

func StripWord(word string) string {
	reg := regexp.MustCompile("[A-Za-z0-9/-]+")

	stripped := reg.FindString(word)
	lower := strings.ToLower(stripped)

	return lower
}
