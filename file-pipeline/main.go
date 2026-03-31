package main

import (
	"fmt"
	"strings"
)

func Capitalize(s string) string {
	skip := map[string]bool{
		"a":   true,
		"an":  true,
		"the": true,
		"and": true,
		"but": true,
		"or":  true,
		"nor": true,
		"on":  true,
		"at":  true,
		"to":  true,
		"by":  true,
		"in":  true,
		"of":  true,
		"up":  true,
		"as":  true,
		"is":  true,
	}

	words := strings.Fields(s)
	for i, word := range words {
		if _, ok := skip[word]; !ok || i == 0 {
			words[i] = strings.Title(word)
		}
	}

	return strings.Join(words, " ")

}

func main() {
	s := "      Is  It Okay       sentinel    of online  "

	fmt.Println(Capitalize(s))
}
