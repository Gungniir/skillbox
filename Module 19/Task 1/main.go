package main

import "fmt"

func main() {
	fmt.Println(concat([4]int{1, 2, 3, 4}, [5]int{1, 2, 3, 4, 5}))
}

func concat(a [4]int, b [5]int) [9]int {
	var result [9]int

	for i, item := range a {
		result[i] = item
	}

	for i, item := range b {
		result[i+4] = item
	}

	return result
}
