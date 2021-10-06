package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 12

func main() {
	var arr [size]int

	rand.Seed(time.Now().UnixNano())

	fmt.Print("Массив: ")
	last := 0
	for i := range arr {
		last += rand.Intn(10)
		arr[i] = last
		fmt.Printf("%d ", last)
	}
	fmt.Println()

	var n int
	fmt.Println("Введите число")
	fmt.Print("n: ")
	_, err := fmt.Scan(&n)

	if err != nil {
		panic(err)
	}

	var (
		minIndex = 0
		maxIndex = len(arr) - 1
	)

	// Бинарный поиск
	for minIndex != maxIndex {
		currentIndex := (minIndex + maxIndex) / 2

		if arr[currentIndex] > n {
			maxIndex = currentIndex
		} else if arr[currentIndex] < n {
			minIndex = currentIndex + (minIndex+maxIndex)%2
		} else {
			minIndex = currentIndex
			maxIndex = currentIndex
		}
	}

	// Поиск первого вхождения
	for minIndex > 0 && arr[minIndex-1] == arr[minIndex] {
		minIndex--
	}

	if arr[minIndex] == n {
		fmt.Printf("Число найдено, его индекс: %d\n", minIndex)
	} else {
		fmt.Println("Такого числа в массиве нет")
	}

}
