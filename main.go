package main

import (
	"fmt"
	"go-snippet/cmd"
	"os"

	"github.com/akamensky/argparse"
)

// rm /Users/sheanlin/.gosnps
// cat /Users/sheanlin/.gosnps | jq .
// go build main.go && ./main -i

// Insert snippets and tag them
// View recent snippets
// View most commonly accessed snippets
// Access snippets by a some short ID
// Look up snippets using fuzzy search with tags
func main() {
	parser := argparse.NewParser("SnippetStore", "Quickly store and lookup snippets")

	insert := parser.Flag("i", "insert a snippet", &argparse.Options{Required: false, Help: "Paste a snippet directly from your clipboard."})
	recent := parser.Flag("r", "recently used snippets", &argparse.Options{Required: false, Help: "View recently used snippets."})
	new := parser.Flag("n", "new snippets", &argparse.Options{Required: false, Help: "View recently added snippets."})

	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
	}

	switch {
	case *insert:
		cmd.Insert()
	case *recent:
		fmt.Println("recent")
	case *new:
		fmt.Println("new")
	default:
		cmd.Select()
	}

	// test, _ := clipboard.ReadAll()
	// fmt.Println(test)
}
