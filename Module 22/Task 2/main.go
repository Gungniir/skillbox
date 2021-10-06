package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 20

func main() {
	var arr [size]int

	rand.Seed(time.Now().UnixNano())

	fmt.Print("Массив: ")
	for i := range arr {
		arr[i] = rand.Intn(200) - 100
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	var n int

	fmt.Println("Введите число")
	fmt.Print("n: ")
	_, err := fmt.Scan(&n)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Элементов после %d: ", n)
	answer := 0
	for i, item := range arr {
		if item == n {
			answer = len(arr) - i - 1
			break
		}
	}
	fmt.Println(answer)

}
