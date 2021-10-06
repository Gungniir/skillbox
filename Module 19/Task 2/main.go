package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 6

func main() {
	var arr [size]int

	rand.Seed(time.Now().Unix())

	for i := range arr {
		arr[i] = rand.Intn(100)
	}

	fmt.Printf("Начальный массив: %v\n", arr)

	arr = sort(arr)

	fmt.Printf("Отсротированный массив: %v\n", arr)
}

func sort(a [size]int) [size]int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	return a
}
