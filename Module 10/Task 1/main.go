package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Введите x, точность, а программа посчитает e^x")

	fmt.Print("x: ")
	var x float64
	_, err := fmt.Scan(&x)

	if err != nil {
		panic(err)
	}

	fmt.Print("Точность(по 50): ")
	var epsilon float64
	_, err = fmt.Scan(&epsilon)

	epsilon = math.Pow(10, -epsilon)

	if err != nil {
		panic(err)
	}

	answer := 1.
	fact := 1.

	for i := 1.; ; i++ {
		fact *= i
		next := answer + (math.Pow(x, i) / fact)

		if math.Abs(answer-next) < epsilon {
			fmt.Printf("Ответ: %.50f\n", next)
			break
		}

		answer = next
	}
}
