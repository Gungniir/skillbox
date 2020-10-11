package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Программа считает сумму скидки")

	var cost float64
	fmt.Println("Введите сумму чека")
	_, err := fmt.Scan(&cost)
	failOnError(err)

	var discountRate float64
	fmt.Println("Введите скидку в процентах")
	_, err = fmt.Scan(&discountRate)
	failOnError(err)

	if discountRate > 30 {
		fmt.Println("Скидка не может превышать 30%, устанавливаем максимальную скидку в 30%")
		discountRate = 30
	}

	discount := cost * discountRate / 100

	if discount > 2000 {
		fmt.Println("Сумма скидки не может превышать 2000 руб, устанавливаем максимальную сумму скидку в 2000 руб")
		discount = 2000
	}

	fmt.Printf("Итого сумма скидки: %.2f руб\n", discount)
}

func failOnError(e error) {
	if e == nil {
		return
	}

	fmt.Printf("Произошла ошибка: %s\n", e.Error())
	os.Exit(1)
}
