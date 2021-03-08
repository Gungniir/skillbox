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
	cash := make([]int, 2) // 0 - пятёрки, 1 - десятки

	for _, bill := range bills {
		switch bill {
		case lemonadeCost:
			cash[0]++
		case 10:
			if cash[0] == 0 {
				return false
			}
			cash[0]--
			cash[1]++
		case 20:
			if cash[1] > 0 && cash[0] > 0 {
				cash[1]--
				cash[0]--
				break
			}
			if cash[0] > 2 {
				cash[0] -= 3
				break
			}
			return false
		}
	}

	return true
}
