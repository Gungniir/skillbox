package main

import "fmt"

func main() {
	var (
		workDuration    int = 480
		orderDuration   int = 2
		packingDuration int = 4
	)

	fmt.Println("Введите длительность смены в минутах:", workDuration)
	fmt.Println("Сколько минут клиент делает заказ?", orderDuration)
	fmt.Println("Сколько минут кассир собирает заказ?", packingDuration)
	fmt.Println("--------------")
	fmt.Println("За смену длиной", workDuration, "минут кассир успеет обслуживать",
		workDuration/(orderDuration+packingDuration), "клиентов")
}
