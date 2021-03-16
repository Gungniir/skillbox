package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Сумма(руб.): ")
	var s float64
	_, err := fmt.Scan(&s)

	if err != nil {
		panic(err)
	}

	s *= 100          // Получаем сумму в копейках
	s = math.Trunc(s) // Отбрасываем лишние данные

	fmt.Print("Ежемесячный процент(%): ")
	var p float64
	_, err = fmt.Scan(&p)

	if err != nil {
		panic(err)
	}

	p = 1 + p/100 // Сохраняем процент в нужно нам форме

	fmt.Print("Кол-во лет: ")
	var n int
	_, err = fmt.Scan(&n)

	n *= 12 // Получаем кол-во месяцев

	if err != nil {
		panic(err)
	}

	fee := 0.

	for i := 0; i < n; i++ {
		next := math.Trunc(s * p)

		fee += s*p - next
		s = next
	}

	fmt.Println("---------------")
	fmt.Printf("Вы получите: %.2f руб.\n", s/100)
	fmt.Printf("Банк заработает на округлении: %.2f руб.\n", fee/100)
}
