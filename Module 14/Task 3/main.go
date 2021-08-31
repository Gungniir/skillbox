package main

import "fmt"

func main() {
	fmt.Println("Введите число")

	var input int
	if _, err := fmt.Scan(&input); err != nil {
		panic(err)
	}

	fmt.Println(sum(mul(input)))
}

func mul(a int) (out int) {
	out = a * 2
	return
}

func sum(a int) (out int) {
	out = a + 2
	return
}
