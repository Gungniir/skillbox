package main

import (
	"fmt"
	"os"
)

func main() {
	// Для упрощения будем считать всё в рублях, т.е. про копейки забудем (иначе можно было бы просто всё считать
	// в копейках, чтобы не использовать float64)

	fmt.Println("Данная программа вычисляет сумму налога по прогрессивной шкале в зависимости от полученного заработка: " +
		"13% на доход до 10 000 руб., 20% на доход от 10 000 до 50 000 руб., 30% на доход выше 50 000 руб.")

	var rawIncome int
	fmt.Println("Введите свой доход до вычета налогов в рублях")
	_, err := fmt.Scan(&rawIncome)
	checkError(err)

	fmt.Println("------")

	taxBase := rawIncome
	var tax int
	if taxBase > 50_000 {
		fmt.Printf("Вы заплатите 30%% с %d: %d руб.\n", taxBase-50_000, (taxBase-50_000)*30/100)
		tax += (taxBase - 50_000) * 30 / 100 // 30%
		taxBase = 50_000
	}
	if taxBase > 10_000 {
		fmt.Printf("Вы заплатите 20%% с %d: %d руб.\n", taxBase-10_000, (taxBase-10_000)*20/100)
		tax += (taxBase - 10_000) * 20 / 100 // 20%
		taxBase = 10_000
	}
	fmt.Printf("Вы заплатите 13%% с %d: %d руб.\n", taxBase, taxBase*13/100)
	tax += taxBase * 13 / 100 // 13%

	fmt.Println("------")
	fmt.Printf("Итого налогов: %d\n", tax)
	fmt.Println("------")

	income := rawIncome - tax

	fmt.Printf("Доход с учетом налогов: %d\n", income)
}

func checkError(err error) {
	if err == nil {
		return
	}

	fmt.Println("Произошла ошибка... Попробуйте перезапустить программу")
	os.Exit(1)
}
