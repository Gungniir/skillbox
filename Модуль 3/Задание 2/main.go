package main

import (
	"fmt"
)

var stops = [...]string{
	"Улица Программистов",
	"Проспект Алгоритмов",
	"Метро Скиллбокс",
	"Платформа Golang",
}

//goland:noinspection GoUnhandledErrorResult
func main() {
	fmt.Println("Программа учета доходов маршрута \"Скиллбоксовский\".\n" +
		"На маршруте всего 4 остановки. Цена проезда 20 рублей.")

	fmt.Println("-------------")

	fmt.Println("Выезжаем из депо...")

	passengersCount := 0
	income := 0

	for i := 0; i < 4; i++ {
		fmt.Println("------Едем------")
		fmt.Printf("Прибываем на остановку %q. В салоне пассажиров: %d.\n", stops[i], passengersCount)

		var input int // Что будет хранится будет понятно из вопроса перед сканом

		fmt.Println("Сколько пассажиров вышло на остановке? ")
		fmt.Scan(&input)
		passengersCount -= input

		fmt.Println("Сколько пассажиров зашло на остановке? ")
		fmt.Scan(&input)
		passengersCount += input
		income += input * 20

		fmt.Printf("Отправляемся с отсновки %q. В салоне пассажиров: %d.\n", stops[i], passengersCount)
	}

	fmt.Println("-------------")

	fmt.Println("Это была конечная остановка, отправляемся в депо...")

	if passengersCount > 0 {
		fmt.Println("Так, стоп, у нас остался лишний груз в автобусе. Выпинываем всех веником!")

		// Условно принимаем, что у нас средний человек :)
		fmt.Printf("Что ж, наш автобус стал весить на %d кг меньше, едем в депо :)\n", passengersCount*60)
	}

	fmt.Println("-------------")
	salary := income / 4
	fuelCost := income / 5
	tax := income / 5
	repairing := income / 5
	total := income - (salary + fuelCost + tax + repairing)

	fmt.Printf("Всего заработали: %d руб.\n", income)
	fmt.Printf("Зарплата водителя: %d руб.\n", salary)
	fmt.Printf("Расзоды на топливо: %d руб.\n", fuelCost)
	fmt.Printf("Налоги: %d руб.\n", tax)
	fmt.Printf("Расходы на ремонт машины: %d руб.\n", repairing)
	fmt.Printf("Итого доход: %d руб.\n", total)
}
