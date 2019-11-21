package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-snippet/structs"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

const Store = ".gosnps"

func ReadJSONBlob() structs.SnippetMap {
	store := getStore()
	var dat structs.SnippetMap

	// If the file doesn't exist, return an empty map and
	// create a new file with an empty JSON object.
	if !fileExists(store) {
		file, err := os.Create(store)
		file.WriteString("{}")
		if err != nil {
			panic(err)
		}
		file.Close()
		return structs.SnippetMap{}
	}

	// Otherwise, business as usual.
	blob, err := ioutil.ReadFile(store)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(blob, &dat); err != nil {
		panic(err)
	}
	return dat
}

func WriteJSONBlob(newSnippet structs.Snippet, snippets structs.SnippetMap) structs.SnippetMap {
	// Store Snippets with their text as their primary key
	snippets[newSnippet.Text] = newSnippet

	store := getStore()
	result, err := json.Marshal(snippets)
	if err != nil {
		log.Fatal(err)
	}
	os.Create(store)
	ioutil.WriteFile(store, result, 0644)
	fmt.Println("Wrote to:", store)
	return snippets
}

func PromptChoice(prompt string) bool {
	snr := bufio.NewScanner(os.Stdin)
	fmt.Println(prompt)
	for fmt.Print("> "); snr.Scan(); fmt.Print("> ") {
		if choice := strings.ToLower(strings.TrimSpace(snr.Text())); choice != "y" {
			return false
		}
		break
	}
	return true
}

func getStore() string {
	return (homeDir() + "/" + Store)
}

func homeDir() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	return user.HomeDir
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
