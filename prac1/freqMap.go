package main

import "fmt"

/*
Write a program that:
1. Stores `"innocent"` in a variable
2. Builds a frequency map of all characters
3. Prints each character and its count

Expected output:
```
i → 1
n → 3
o → 1
c → 1
e → 1
t → 1
*/

func wordFrequencyCOunter(name string) map[rune]int {
	freq := make(map[rune]int)

	for _, ch := range name {
		freq[ch]++
	}

	for char, count := range freq {
		fmt.Printf("%c --> %d\n", char, count)
	}
	return freq
}
