// package wordutil
package main

import (
	"fmt"
	"strings"
)

// Finds first occurrence of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// the index of the first occurence of the word in the input string as the
// corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordIndex(s string) map[string]int {
	word_index := make(map[string]int)
	s = strings.ToLower(s)
	words := strings.Split(s, " ")
	for _, word := range words {
		word_index[word] = strings.Index(s, word)
	}
	return word_index
}

func main() {
	fmt.Println(WordIndex("A string to do a test"))
}
