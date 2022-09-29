package main

import 	("strings"; "fmt")

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	s = strings.ToLower(s)
	words := strings.Fields(s)
    for _, word := range words {
		counts[word]++
    }
	return counts
}

func main() {
	fmt.Println(WordCount("A sentence for a testing"))
}