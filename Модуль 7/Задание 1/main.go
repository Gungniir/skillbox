package main

import (
	"fmt"
	"math"
)

func main() {
	count := 0

	// Реализация без вложенных циклов
	//for i := 100000; i <= 999999; i++ {
	//	if i / 1000 == i % 1000 {
	//		count++
	//	}
	//}

	// Реализация с двумя вложенными циклами
	for i := 100000; i <= 999999; i++ {
		var start, end int

		// Получение первых трех цифр
		for j := 0; j < 3; j++ {
			start = start*10 + (i / int(math.Pow10(5-j)) % 10)
		}

		// Получение последних трех цифр
		for j := 0; j < 3; j++ {
			end = end*10 + (i / int(math.Pow10(2-j)) % 10)
		}

		if start == end {
			count++
		}
	}

	fmt.Println(count)
}
