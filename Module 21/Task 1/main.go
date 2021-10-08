package main

import "fmt"

func main() {
	var (
		x int16   = 254
		y uint8   = 165
		z float32 = 1.2
	)

	fmt.Println(func(x int16, y uint8, z float32) float32 { return 2*float32(x) + float32(y)*float32(y) - 3/z }(x, y, z))
}
