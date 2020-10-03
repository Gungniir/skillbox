package main

import "fmt"

func main() {
	var (
		sum               int = 4_000_000
		porchCount        int = 10
		flatsCountInPorch int = 40
	)

	fmt.Println("Сумма, указанная в квитанции:", sum)
	fmt.Println("Подъездов в доме:", porchCount)
	fmt.Println("Квартир в каждом подъезде:", flatsCountInPorch)
	fmt.Println("-------------")
	fmt.Println("Каждая квартира должна заплатить по", sum/porchCount/flatsCountInPorch, "руб.")
}
