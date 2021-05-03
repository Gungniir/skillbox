package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	str := "Go is an Open source programming Language that makes it ♥ 123 ♠ Easy to build simple, reliable, and efficient Software. Абябь. Ъьъ"

	fmt.Println(UpperWordCont(str))
}

func UpperWordCont(str string) int {
	words := strings.Fields(str)
	count := 0

	for _, word := range words {
		rn, _ := utf8.DecodeRuneInString(word) // First rune

		if unicode.IsUpper(rn) {
			count += 1
		}
	}

	return count
}
