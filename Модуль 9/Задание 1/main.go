package main

import (
	"fmt"
	"math"
)

func main() {
	overflows := make([]int, 2)
	var lastA uint8
	var lastB uint16

	for i := 0; i <= math.MaxUint32; i++ {
		a, b := uint8(i), uint16(i)

		if a < lastA {
			overflows[0]++
		}

		if b < lastB {
			overflows[1]++
		}

		lastA, lastB = a, b
	}

	fmt.Printf("Переполнений uint8: %d\nПереполнений uint16: %d", overflows[0], overflows[1])
}
