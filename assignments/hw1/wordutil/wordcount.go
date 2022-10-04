package wordutil

// package main

import (
	"strings"
)

// Counts occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// its count as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordCount(s string) map[string]int {
	word_count := make(map[string]int)
	s = strings.ToLower(s)
	words := strings.Fields(s)
	for _, word := range words {
		word_count[word]++
	}
	return word_count
}

// func main() {
// 	fmt.Println(WordCount("A string to do a test"))
// }
