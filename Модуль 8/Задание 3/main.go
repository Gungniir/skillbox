package main

import (
	"encoding/json"
	"fmt"
)

const lemonadeCost = 5

func main() {
	fmt.Printf("Привет! Эта программа помощи продавцам лимонада по %d рублей! :)\nВведите чеки в очереди :)\n", lemonadeCost)

	var billsRaw []byte
	_, err := fmt.Scan(&billsRaw)

	if err != nil {
		panic("Не удалось считать массив купюр")
	}

	var bills []int
	err = json.Unmarshal(billsRaw, &bills)

	if err != nil {
		panic("Неверный формат входных данных: используйте [5, 20, 35]")
	}

	fmt.Println(lemonadeChange(bills))
}

func lemonadeChange(bills []int) bool {
	cash := 0

	for _, bill := range bills {
		if bill-lemonadeCost > cash {
			return false
		}
		cash -= bill - lemonadeCost // Сдача
		cash += lemonadeCost        // Цена
	}

	return true
}
