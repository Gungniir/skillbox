package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Данная программа подскажет, если среди чисел хотя бы одно положительное")

	var a, b, c int

	fmt.Println("Введите три целых числа через пробел")
	_, err := fmt.Scanf("%d%d%d", &a, &b, &c)
	failOnError(err)

	if a > 0 || b > 0 || c > 0 {
		fmt.Println("Среди введенных чисел есть положительное число")
	} else {
		fmt.Println("Среди введенных чисел нет положительных")
	}

}

func failOnError(e error) {
	if e == nil {
		return
	}

	fmt.Printf("Произошла ошибка: %s\n", e.Error())
	os.Exit(1)
}
