package main

import (
	"fmt"
)

func main() {
	fmt.Println(countConsistentStrings("exdohslrwipnt", []string{
		"xrwlstne", "rs", "ioetdll", "lwi", "r", "pieonois", "r", "xtp", "stia", "gicfuvmnr", "hdntpxse", "sodxws", "v", "hstirooon", "d",
	}))
}

func isConsistent(word string, allowedSet map[rune]bool) bool {
	for _, letter := range word {
		if !allowedSet[letter] {
			return false
		}
	}
	return true
}

func countConsistentStrings(allowed string, words []string) int {
	allowedSet := make(map[rune]bool)
	for _, letter := range allowed {
		allowedSet[letter] = true
	}

	var count int
	for _, word := range words {
		if isConsistent(word, allowedSet) {
			count++
		}
	}

	return count
}
