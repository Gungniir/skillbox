package main

import "fmt"

func main() {
	fmt.Println("Введите десять целых числ, а программа подсчитает кол-во чётных и нечётных")

	var buf [10]int

	for i := range buf {
		fmt.Printf("Число №%d: ", i+1)
		if _, err := fmt.Scan(&buf[i]); err != nil {
			panic(err)
		}
	}

	var (
		numberOfOdd  int
		numberOfEven int
	)

	for _, i := range buf {
		if i%2 == 0 {
			numberOfEven++
		} else {
			numberOfOdd++
		}
	}

	fmt.Printf("Вы ввели %d чётных и %d нечётных чисел\n", numberOfEven, numberOfOdd)
}
