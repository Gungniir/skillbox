package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Привет! Эта программа напомнит вам, как называются оставшиеся рабочие дни недели :)")

	var day string
	_, err := fmt.Scan(&day)

	if err != nil {
		panic("Не удалось считать название месяца")
	}

	day = strings.ToLower(day)
	switch day {
	case "пн":
		fmt.Println("Понедельник")
		fallthrough
	case "вт":
		fmt.Println("Вторник")
		fallthrough
	case "ср":
		fmt.Println("Среда")
		fallthrough
	case "чт":
		fmt.Println("Четверг")
		fallthrough
	case "пт":
		fmt.Println("Пятница")
	default:
		fmt.Println("Такого дня нет! :)")
	}
}
