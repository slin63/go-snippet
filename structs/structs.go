package structs

import "strings"

// Limits on how many lines we'll display
const MaxLines = 10
const Spacer = "-----------------------"

type Snippet struct {
	Text        string
	Tags        []string
	AccessCount int
	Id          string
	// length      int

}

// Return the number of lines in the snippet's text
func (s Snippet) Height() int {
	return strings.Count(s.Text, "\n")
}
