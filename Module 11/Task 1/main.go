package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software. Абябь. Ъьъ"

	fmt.Println(UpperWordCont(str))
}

func UpperWordCont(fullString string) int {
	count := 0

	for _, s := range strings.Split(fullString, " ") {
		if len(s) > 0 && s[0] >= 'A' && s[0] <= 'Z' { // Латиница (один символ - один байт)
			count++
		} else if len(s) > 1 && s[:2] >= "А" && s[:2] <= "Я" { // Кириллица (один символ - два байта)
			count++
		}
	}

	return count
}
