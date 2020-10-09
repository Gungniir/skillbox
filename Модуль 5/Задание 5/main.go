package main

import (
	"fmt"
	"os"
)

const ratesCount = 3

func main() {
	fmt.Printf("Введите %d процентные ставки, а программа выведет вам лучшие два предложения\n", ratesCount)

	// max[1] - максимальная процентная ставка, max[0] - максимальная процентная ставка после max[1],
	// т.е. max[0] >= max[1]
	var max [2]float64

	for i := 0; i < ratesCount; i++ {
		var rate float64
		_, err := fmt.Scan(&rate)
		failOnError(err)

		if rate > max[1] {
			max[1] = rate
		} else if rate > max[0] {
			max[0] = rate
		}
	}

	fmt.Printf("Лучшие процентные ставки: %.2f%%; %.2f%%\n", max[0], max[1])
}

func failOnError(e error) {
	if e == nil {
		return
	}

	fmt.Printf("Произошла ошибка: %s\n", e.Error())
	os.Exit(1)
}
