package mystr

import "strings"

func Cat(words []string) string {
	joinedWords := words[0]

	for _, word := range words[1:] {
		joinedWords += " "
		joinedWords += word
	}

	return joinedWords
}

func Join(words []string) string {
	return strings.Join(words, " ")
}
