package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Привет! Эта программа напомнит вам, к какому времени года относятся месяцы! Введите название месяца в именительном падеже :)")

	var month string
	_, err := fmt.Scan(&month)

	if err != nil {
		panic("Не удалось считать название месяца")
	}

	switch strings.ToLower(month) {
	case "декабрь", "январь", "февраль":
		fmt.Println("Да это же зима!")
	case "март", "апрель", "май":
		fmt.Println("Да это же весна!")
	case "июнь", "июль", "август":
		fmt.Println("Да это же лето!")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("Да это же осень!")
	default:
		fmt.Println("Такого месяца нет...")
	}
}
