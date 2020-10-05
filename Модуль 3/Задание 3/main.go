package main

import (
	"fmt"
)

func main() {
	a := 42
	b := 153

	a, b = b, a

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
