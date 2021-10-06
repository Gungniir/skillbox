package main

import "log"

func main() {
	work(10, 20, func(a int, b int) int {
		return a + b
	})
	work(10, 20, func(a int, b int) int {
		return a - b
	})
	work(10, 20, func(a int, b int) int {
		return a * b
	})
}

func work(a, b int, A func(int, int) int) {
	defer func() {
		res := A(a, b)
		log.Printf("Res: %d", res)
	}()
}
