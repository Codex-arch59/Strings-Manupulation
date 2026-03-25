package main

func frequentChar(s string) (rune, int) {
	mostFreq := make(map[rune]int)

	for _, ch := range s {
		if ch == ' ' {
			continue
		}
		mostFreq[ch]++
	}

	var topChar rune
	topCount := 0

	for char, count := range mostFreq {
		if count > topCount {
			topCount = count
			topChar = char
		}
	}

	return topChar, topCount
}
