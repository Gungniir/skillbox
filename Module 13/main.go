package main

import "fmt"

func main() {
	rev(sum, mul, 5, 2)
}

func sum(a, b int) {
	fmt.Println(a + b)
}

func mul(a, b int) {
	fmt.Println(a * b)
}

func rev(f func(a, b int), g func(a, b int), a, b int) {
	g(a, b)
	f(a, b)
}
