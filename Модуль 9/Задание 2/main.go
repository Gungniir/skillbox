package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int16

	_, err := fmt.Scanf("%d %d", &a, &b)

	if err != nil {
		panic(err)
	}

	result := int32(a) * int32(b)

	if result >= 0 {
		switch {
		case result <= math.MaxUint8:
			fmt.Println("uint8")
		case result <= math.MaxUint16:
			fmt.Println("uint16")
		default:
			fmt.Println("uint32")
		}

		return
	}

	result = -result

	switch {
	case result <= math.MaxInt8:
		fmt.Println("int8")
	case result <= math.MaxInt16:
		fmt.Println("int16")
	default:
		fmt.Println("int32")
	}
}
