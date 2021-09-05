package main

import "fmt"

func main() {
	fmt.Println("Введите десять целых числ, а программа выведет реверсированный массив")

	var buf [10]int

	for i := range buf {
		fmt.Printf("Число №%d: ", i+1)
		if _, err := fmt.Scan(&buf[i]); err != nil {
			panic(err)
		}
	}

	for _, item := range reverse(buf) {
		fmt.Printf("%d ", item)
	}
}

func reverse(a [10]int) (b [10]int) {
	for index, item := range a {
		b[9-index] = item
	}
	return
}
