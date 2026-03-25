package main

import "strings"

func fixQuotes(words []string) []string {
	for i := 0; i < len(words); i++ {

		if words[i] == "'" {
			for j := i + 1; j < len(words); j++ {
				if words[j] == "'" {
					inside := strings.Join(words[i+1:j], " ")
					words[i] = "'" + inside + "'"
					words = append(words[:i+1], words[j+1:]...)
					break
				}
			}
		}
	}
	return words
}
