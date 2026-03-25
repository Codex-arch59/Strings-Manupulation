package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . input.txt output.txt")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Read file successfull!")
	}

	words := strings.Fields(string(data))
	fmt.Println("file content:", words)

	// commands
	words = processCommands(words)

	// Punctuation function
	words = fixPunctuation(words)

	//qoutes
	words = fixQuotes(words)

	// grammers a to an
	words = fixGrammar(words)

	result := strings.Join(words, " ")

	fmt.Println("Input file:", inputFile)
	fmt.Println("Output file:", outputFile)

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("successful, check result in output file.")
	}
}
