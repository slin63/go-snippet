package cmd

import (
	"go-snippet/structs"
	"sort"
)

func sortByAccessDate(snippets *[]structs.Snippet) {
	sort.Sort(structs.SnippetsByDate(*snippets))
}
