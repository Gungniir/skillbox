package main

import (
	"fmt"
)

func main() {
	fmt.Println("Программа максимальное количество билетов, которое нужно купить, чтобы точно получить счастливый билет")

	last := 100001
	max := 1

	for i := 100002; i <= 999999; i++ {
		if i%1000 == i/1000 {
			if max < i-last {
				max = i - last
			}

			last = i
		}
	}

	fmt.Println(max)
}
