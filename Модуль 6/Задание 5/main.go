package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Программа показывает алгоритм распределения яблок по корзинкам")

	var baskets Baskets

	for i := 0; i < len(baskets); i++ {
		var input int
		fmt.Printf("Введите ёмкость корзины √%d\n", i+1)
		_, err := fmt.Scan(&input)
		failOnError(err)

		baskets[i].Max = input
	}

	// q – индекс корзины, к которой мы будет обращаться
	for q := 0; ; q = (q + 1) % len(baskets) {
		if baskets.isFull() {
			fmt.Println("Корзинки успешно заполнены :)")
			break
		}

		if baskets[q].Value == baskets[q].Max {
			// Корзина заполнена
			continue
		}

		baskets[q].Value++

		fmt.Println("---Проход---")
		baskets.printStatus()
		fmt.Println("------------")
	}
}

func failOnError(e error) {
	if e == nil {
		return
	}

	fmt.Printf("Произошла ошибка: %s\n", e.Error())
	os.Exit(1)
}
