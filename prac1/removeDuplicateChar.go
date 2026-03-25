package main

import (
	"strings"
)

/*
Write a program that:
1. Creates a `removeDuplicates` function
2. Tests it with these strings:
   - `"innocent"`
   - `"golang"`
   - `"programming"`

Expected output:
```
inoct e
golang  (wait - are there duplicates? check carefully!)
progamin

*/

func removeDuplicate(word string) string {
	word = strings.ToLower(strings.TrimSpace(word))
	seen := make(map[rune]bool)

	result := []rune{}

	for _, ch := range word {
		if !seen[ch] {
			seen[ch] = true
			result = append(result, ch)
		}
	}
	return string(result)
}
