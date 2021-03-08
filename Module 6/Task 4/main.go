package main

import (
	"fmt"
)

func main() {
	fmt.Println("Тренируемся работе с циклами")

	a, b, c := 0, 0, 0

	for {
		if a < 200 {
			a++
		}

		if b < 100 {
			b++
		}

		c++

		if c == 1000 {
			break
		}
	}

	fmt.Println(a, b, c)
}
