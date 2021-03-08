package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Программа складывает числа очень хитрым образом")

	var a, b int
	fmt.Println("Введите два целых положительных числа через пробел")
	_, err := fmt.Scanf("%d%d", &a, &b)
	failOnError(err)

	sum := a + b
	for a != sum {
		a += 1
	}

	fmt.Println("Сумма:", a)
}

func failOnError(e error) {
	if e == nil {
		return
	}

	fmt.Printf("Произошла ошибка: %s\n", e.Error())
	os.Exit(1)
}
