package main

import "fmt"

var number1 = 5
var number2 = 6
var number3 = 7

func main() {
	fmt.Println("Программа показывает, как работает область видимости в Go")

	fmt.Println("Введите число")

	var input int
	if _, err := fmt.Scan(&input); err != nil {
		panic(err)
	}

	fmt.Println(addNumber1(addNumber2(addNumber3(input))))
}

func addNumber1(a int) int {
	return a + number1 + number2
}

func addNumber2(a int) int {
	return a + number2 + number3
}

func addNumber3(a int) int {
	return a + number1 + number3
}
