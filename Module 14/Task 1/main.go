package main

import "fmt"

func main() {
	fmt.Println("Введите число, а программа подскажет вам, четное ли оно")

	var a int
	if _, err := fmt.Scan(&a); err != nil {
		panic(err)
	}

	fmt.Printf("Четное ли число %d: %t\n", a, isEven(a))
}

func isEven(a int) bool {
	return a%2 == 0
}
