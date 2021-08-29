package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int

	fmt.Println("Введите натуральное число скобочек, и программа выведет все возможные комбинации")

	if _, err := fmt.Scan(&a); err != nil {
		panic(err)
	}

	fmt.Println(generate(a, ""))
}

func generate(len int, now string) []string {
	if len == 0 {
		need := strings.Count(now, "(") - strings.Count(now, ")")
		for i := 0; i < need; i++ {
			now += ")"
		}
		return []string{now}
	}

	out := make([]string, 0)

	if strings.Count(now, "(") > strings.Count(now, ")") {
		out = append(out, generate(len, now+")")...)
	}
	out = append(out, generate(len-1, now+"(")...)
	return out
}
