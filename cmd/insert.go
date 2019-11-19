package cmd

import (
	"bufio"
	"fmt"
	"go-snippet/structs"
	"log"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

const tagInstructions = `Enter tags for this snippet.
Submit a blank tag to finish.`

func Insert() {
	tags := []string{}
	text, err := clipboard.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Display the successfully extracted snippet text
	snippet := structs.Snippet{Text: strings.TrimSpace(text)}
	fmt.Println("Received:")
	fmt.Println(displayString(snippet))

	// Prompt the user for tags
	setTags(&tags)
	fmt.Println(tags)
}

func setTags(tags *[]string) {
	snr := bufio.NewScanner(os.Stdin)
	fmt.Println(tagInstructions)
	for fmt.Print("> "); snr.Scan(); fmt.Print("> ") {
		tag := snr.Text()
		if len(tag) == 0 {
			break
		}
		*tags = append(*tags, tag)
	}
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
		"%s\n%s\n%s\n%s\n",
		structs.Spacer,
		s.Text[:indices[len(indices)-1]],
		". . .",
		structs.Spacer,
	)
}
