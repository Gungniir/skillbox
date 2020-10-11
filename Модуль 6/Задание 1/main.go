package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Программа выводит все числа с нуля по указанное вами число")

	var to int
	fmt.Println("Введите число")
	_, err := fmt.Scan(&to)
	failOnError(err)

	for i := 0; i <= to; i++ {
		fmt.Println(i)
	}
}

func failOnError(e error) {
	if e == nil {
		return
	}

	fmt.Printf("Произошла ошибка: %s\n", e.Error())
	os.Exit(1)
}
