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
const alreadyExists = `This snippet already exists.
Would you like to replace it? (y/n)`

func Insert() {
	// Grab existing snippets
	var snippets structs.SnippetMap = ReadJSONBlob()

	tags := []string{}
	text, err := clipboard.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	text = strings.TrimSpace(text)

	if _, contains := snippets[text]; contains {
		if !PromptChoice(alreadyExists) {
			os.Exit(0)
		}
	}

	// Display the successfully extracted snippet text
	fmt.Println("Received:")
	fmt.Println(displayString(text))

	// Prompt the user for tags
	setTags(&tags)
	newSnippet := *structs.NewSnippet(text, tags)

	WriteJSONBlob(newSnippet, snippets)
}

func setTags(tags *[]string) {
	snr := bufio.NewScanner(os.Stdin)
	fmt.Println(tagInstructions)
	for fmt.Print("> "); snr.Scan(); fmt.Print("> ") {
		tag := strings.TrimSpace(snr.Text())
		if len(tag) == 0 {
			break
		}
		*tags = append(*tags, tag)
	}
}

func displayString(s string) string {
	if strings.Count(s, "\n") > structs.MaxLines {
		return trim(s)
	}
	return fmt.Sprintf("%s\n%s\n%s", structs.Spacer, s, structs.Spacer)
}

func trim(s string) string {
	indices := []int{}
	for i := 0; i <= len(s) && len(indices) <= structs.MaxLines; i++ {
		if s[i] == '\n' {
			indices = append(indices, i)
		}
	}
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s\n",
		structs.Spacer,
		s[:indices[len(indices)-1]],
		". . .",
		structs.Spacer,
	)
}
