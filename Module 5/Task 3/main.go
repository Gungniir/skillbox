package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c int

	fmt.Println("Введите 3 числа через пробел")
	_, err := fmt.Scanf("%d%d%d", &a, &b, &c)
	failOnError(err)

	if a == b || a == c || b == c {
		fmt.Println("Два или более числа совпадают")
	} else {
		fmt.Println("Совпадений нет")
	}
}

func failOnError(e error) {
	if e == nil {
		return
	}

	fmt.Printf("Произошла ошибка: %s\n", e.Error())
	os.Exit(1)
}
