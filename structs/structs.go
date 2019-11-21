package structs

import (
	"fmt"
	"strings"
	"time"
)

// Formatting
const MaxLines = 10
const Spacer = "-----------------------"

type SnippetMap map[string]Snippet
type SnippetsByDate []Snippet

// Unique identifier: Snippet.Text
func NewSnippet(text string, tags []string) *Snippet {
	return &Snippet{
		Text:         text,
		Tags:         tags,
		AccessCount:  0,
		LastAccessed: time.Now(),
	}
}

type Snippet struct {
	Text         string    `json:"text"`
	Tags         []string  `json:"tags"`
	AccessCount  int       `json:"access_count"`
	LastAccessed time.Time `json:"last_accessed"`
}

// Return the number of lines in the snippet's text
func (s Snippet) Height() int {
	return strings.Count(s.Text, "\n")
}

func (sm SnippetMap) AddSnippet(s *Snippet) {
	sm[s.Text] = *s
}

// Update fields inside a snippet when it's accessed
func (s SnippetMap) UpdateAccessFields(text string) {
	snippet, ok := s[text]
	if !ok {
		panic(fmt.Sprintf("Snippet with text: %s not found", text))
	}
	snippet.LastAccessed = time.Now()
	snippet.AccessCount += 1
	s[text] = snippet
}

// Custom sorting types
func (s SnippetsByDate) Less(i, j int) bool { return s[i].LastAccessed.After(s[j].LastAccessed) }
func (s SnippetsByDate) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SnippetsByDate) Len() int           { return len(s) }
