package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите номер билета, а программа подскажет вам, как к нему относиться")

	var ticket int
	_, err := fmt.Scan(&ticket)
	failOnError(err)

	if ticket > 9999 || ticket < 1000 {
		fmt.Println("Неверный формат билета")
		return
	}

	numbers := split(ticket)

	if numbers[0] == numbers[3] && numbers[1] == numbers[2] {
		fmt.Println("У вас зеркальный билет! (И, конечно, счастливый)")
	} else if numbers[0]+numbers[1] == numbers[2]+numbers[3] {
		fmt.Println("У вас счастливый билет!")
	} else {
		fmt.Println("У вас обычный билет :(")
	}

}

func failOnError(e error) {
	if e == nil {
		return
	}

	fmt.Printf("Произошла ошибка: %s\n", e.Error())
	os.Exit(1)
}

func split(a int) [4]int {
	return [4]int{
		a / 1000,
		a / 100 % 10,
		a / 10 % 10,
		a % 10,
	}
}
