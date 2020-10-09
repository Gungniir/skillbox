package main

import (
	"fmt"
	"os"
)

const (
	minimum = 1
	maximum = 10
)

func main() {
	fmt.Printf("Давайте поиграем, загадайте число с %d по %d, а я буду пытаться его угадать :)\n", minimum, maximum)
	min := minimum
	max := maximum

	for {
		if min == max {
			fmt.Printf("Что ж, единсвенное число, удовлетворяющее всем выше условиям – %d :)\n", min)
			return
		}

		guess := (max-min)/2 + min

		fmt.Printf("Я думаю, что вы загадали %d! Я прав?\n", guess)
		fmt.Printf("Введите '<', если загаданное число меньше, или '>', '='  в зависиомсти от случая\n")

		var input string
		_, err := fmt.Scan(&input)
		failOnError(err)

		if input == ">" {
			min = guess + 1
		} else if input == "<" {
			max = guess - 1
		} else if input == "=" {
			fmt.Println("Ура, я отгадал!")
			return
		} else {
			fmt.Println("Ну я так не умею :(")
			return
		}

	}
}

func failOnError(e error) {
	if e == nil {
		return
	}

	fmt.Printf("Произошла ошибка: %s\n", e.Error())
	os.Exit(1)
}
