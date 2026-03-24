package main

import (
	"fmt"
	"strings"
)

/*
strings.Builder in Go
Purpose: strings.Builder is used to efficiently build or concatenate strings without creating many intermediate string objects.
Why it’s useful: In Go, strings are immutable, so using + repeatedly in a loop can be slow and memory-inefficient. strings.Builder avoids this by appending directly to an internal buffer.
Key methods:
WriteString(s string) → appends s to the builder.
String() → returns the final concatenated string.

Additional benefits:

Reduces memory allocations when concatenating many strings.
Cleaner and safer than using + inside loops for large strings.
*/
func main() {

	words := []string{"Go", "is", "fast", "and", "powerful"}

	var sentence strings.Builder

	for i, word := range words {
		if i > 0 {
			sentence.WriteString(" ")
		}
		sentence.WriteString(word)
	}
	result := sentence.String()
	fmt.Println(result)
	fmt.Println(len(result))
	fmt.Println(strings.ToUpper(result))

}
