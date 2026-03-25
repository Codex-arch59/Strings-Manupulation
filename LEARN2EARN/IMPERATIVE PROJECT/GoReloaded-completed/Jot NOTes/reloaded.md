```go 
package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Q1, write a function that converts bin to decimal.
//Q2, write a function that converts hex word to decimal
//Q3, write a function that caps the first letter of a word.
//Q4, write a function that caps the nth word in a slice of striing.
//Q5, write a function that returns true or false if a word contains punctuations.
//Q6, write a function that fix punctuation error in a slice of word or string.
//Q7, write a function that returns an "an" if the first character in the word is a vowel.
//Q8, write a function that fix spaces between trailling and leading qoutes in a slice of string or a string

func binTodecimal(word string) (int64, error) {
	num, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func hexTodecimal(word string) (int64, error) {
	num, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func capsFirstCharacter(word string) string {
	word = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	return word
}

func capsNWord(word []string, n int) []string {
	for i := len(word) - n; i < len(word); i++ {
		if n > len(word) {
			n = len(word)
		}
		word[i] = strings.ToUpper(word[i])
	}
	return word
}

func containsPunctuation(word string) bool {
	if strings.ContainsAny(word, ",.?;:!") {
		return true
	}
	return false
}

func fixPuntuation(word []string) string { // check here!
	for i := len(word) - 1; i > 0; i-- {
		if word[i] == "," || word[i] == "." || word[i] == ":" || word[i] == "?" || word[i] == "!" {
			word[i-1] = word[i-1] + word[i]
			word = append(word[:i], word[i+1:]...)
		}
	}
	return strings.Join(word, " ")
}

func isVowel(word string) string {
	if strings.ContainsAny(word[:1], "aeiouAEIOUhH") {
		return "an"
	}
	return "a"
}

func fixQoutes(qoutes string) string {
	if strings.HasPrefix(qoutes, "'") && strings.HasSuffix(qoutes, "'") {
		inside := strings.TrimSpace(qoutes[1 : len(qoutes)-1])
		return "'" + inside + "'"
	}
	return qoutes
}

func main() {
	sentence := []string{"hello", "innocent", ",", "you", "are", "a", "programmer?"}
	word := "' hello innocent, welcome to GO! '"

	//Q1
	fmt.Println(binTodecimal("1010"))
	//Q2
	fmt.Println(hexTodecimal("FF"))
	//Q3
	fmt.Println(capsFirstCharacter("heLLO"))
	//Q4
	fmt.Println(capsNWord(sentence, 3))
	//Q5
	fmt.Println(containsPunctuation(word))
	//Q6
	fmt.Println(fixPuntuation(sentence))
	//Q7
	fmt.Println(isVowel("apple"))
	fmt.Println(isVowel("book"))
	//Q8
	fmt.Println(fixQoutes(word))

}
```