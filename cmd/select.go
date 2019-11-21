package cmd

import (
	"fmt"
	"go-snippet/structs"

	"github.com/atotto/clipboard"
)

const selectPrompt = `Select a snippet to copy to your clipboard.`

// Offer the user some snippets to select, with most recent snippets at the top.
func Select() {
	// Grab existing snippets
	var snippets structs.SnippetMap = ReadSnippetsFromFile()
	var sortedSnippets *[]structs.Snippet = getSnippetValues(snippets)

	// Format and display options
	displayString := displayEnumeratedSnippets(sortedSnippets)
	fmt.Println(displayString)

	// Prompt user for selection
	var selected int = PromptEnumerable(selectPrompt)

	// Grab selected snippet, update snippet access information
	selectedSnippet := (*sortedSnippets)[selected-1]
	fmt.Println(selectedSnippet.Text)
	snippets.UpdateAccessFields(selectedSnippet.Text)
	fmt.Println(selectedSnippet)

	// Copy to clipboard
	clipboard.WriteAll(selectedSnippet.Text)
	fmt.Printf("Copied choice (%d) to clipboard.\n", selected)

	// Update snippets
	WriteSnippetMapToFile(snippets)
}

func displayEnumeratedSnippets(snippets *[]structs.Snippet) string {
	var displayString = ""
	for i, snippet := range *snippets {
		displayString += fmt.Sprintf("%d.\n%s\n%s\n%s\n\n", i+1, structs.Spacer, snippet.Text, structs.Spacer)
	}
	return displayString
}

func getSnippetValues(snippets structs.SnippetMap) *[]structs.Snippet {
	snippetValues := []structs.Snippet{}
	for _, value := range snippets {
		snippetValues = append(snippetValues, value)
	}
	sortByAccessDate(&snippetValues)

	return &snippetValues
}
