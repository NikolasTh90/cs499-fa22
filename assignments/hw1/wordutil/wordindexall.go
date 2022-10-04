package wordutil

// package main

import (
	// "fmt"
	"strings"
)

// Finds all occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// a slice that contains the index of each occurence of the word in the input
// string as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordIndexAll(s string) map[string][]int {
	words := make(map[string][]int)
	index := -1
	previous_word := ""
	for _, word := range strings.Fields(s) {
		if strings.Index(s, word) > index {
			index = strings.Index(s, word)
		} else {
			index = index + len(previous_word) + 1
		}
		words[strings.ToLower(word)] = append(words[strings.ToLower(word)], index)
		previous_word = word
	}
	return words
}

// func main() {
// 	str := "Apple orange grape grapefruit apple Orange apple"
// 	fmt.Println(WordIndexAll(str))
// }
