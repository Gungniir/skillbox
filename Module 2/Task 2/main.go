package main

import "fmt"

func main() {
	var (
		itemCost     int = 6400
		deliveryCost int = 350
		discount     int = 700
	)

	fmt.Printf("Стоимость товара: %d\n", itemCost)
	fmt.Printf("Стоимость доставки: %d\n", deliveryCost)
	fmt.Printf("Размер скидки: %d\n", discount)
	fmt.Println("-----------")
	fmt.Printf("Итого: %d", itemCost+deliveryCost-discount)
}
