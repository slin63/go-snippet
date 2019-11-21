package structs

import (
	"strings"
)

// Formatting
const MaxLines = 10
const Spacer = "-----------------------"

type SnippetMap map[string]Snippet

// Unique identifier: Snippet.Text
func NewSnippet(text string, tags []string) *Snippet {
	return &Snippet{
		Text: text,
		Tags: tags,
	}
}

type Snippet struct {
	Text        string   `json:"text"`
	Tags        []string `json:"tags"`
	AccessCount int      `json:"access_count"`
}

// Return the number of lines in the snippet's text
func (s Snippet) Height() int {
	return strings.Count(s.Text, "\n")
}
