package cmd

import (
	"fmt"
	"go-snippet/structs"
	"log"
	"strings"

	"github.com/atotto/clipboard"
)

func Insert() {
	text, err := clipboard.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	snippet := structs.Snippet{Text: strings.TrimSpace(text)}
	fmt.Println(snippet.Height())
	fmt.Println(displayString(snippet))
}

func displayString(s structs.Snippet) string {
	if s.Height() > structs.MaxLines {
		return trim(s)
	}
	return fmt.Sprintf("%s\n%s\n%s", structs.Spacer, s.Text, structs.Spacer)
}

func trim(s structs.Snippet) string {
	indices := []int{}
	for i := 0; i <= len(s.Text) && len(indices) <= structs.MaxLines; i++ {
		if s.Text[i] == '\n' {
			indices = append(indices, i)
		}
	}
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s%d%s",
		structs.Spacer,
		s.Text[:indices[len(indices)-1]],
		". . .",
		structs.Spacer,
		"(originally ",
		s.Height(),
		" lines)",
	)
}
