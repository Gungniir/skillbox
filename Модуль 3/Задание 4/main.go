package main

import (
	"fmt"
	"math"
)

//goland:noinspection GoUnhandledErrorResult
func main() {
	var startHeight float64
	fmt.Println("Введите высоту саженца")
	fmt.Scan(&startHeight)

	var increasePerDay float64
	fmt.Println("Введите скорость роста бамбука")
	fmt.Scan(&increasePerDay)

	var decreasePerDay float64
	fmt.Println("Введите скорость поедания бамбука гусеницами")
	fmt.Scan(&decreasePerDay)

	var daysCount float64
	fmt.Println("Введите количество дней, за которое вы хотите рассчитать высоту бамбука")
	fmt.Scan(&daysCount)
	nightsCount := int(daysCount) // Отбрасываем дробную часть?

	var neededHeight float64
	fmt.Println("Введите высоту бамбука, при которой его можно срубить и продать")
	fmt.Scan(&neededHeight)

	totalHeight := startHeight + daysCount*increasePerDay - float64(nightsCount)*decreasePerDay
	fmt.Printf("Через %.2f дней с момента посадки высота бамбука станет равной %.2f см.\n", daysCount, totalHeight)

	neededDays := int(math.Ceil(neededHeight / (increasePerDay - decreasePerDay)))
	fmt.Printf("Бамбук достигнет высоты, при которой его можно срубить и продать через %d дней\n", neededDays)
}
