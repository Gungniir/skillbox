package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Программа выводит ёлочку! Все же помнят, что скоро новым год? Ах... Он уже прошёл...")
	fmt.Print("Высота ёлочки=")

	var x int
	_, err := fmt.Scan(&x)

	if err != nil {
		fmt.Println("Не удалось считать высоту")
		os.Exit(1)
	}

	if x == 0 {
		fmt.Println("Вам правда нужна ёлочка с нулевой высотой? Что ж... Держите:")
		return
	}

	fillWidth := (x-1)*2 + 1

	for i := 0; i < x; i++ {
		currentWidth := i*2 + 1
		starts := strings.Repeat("*", currentWidth)
		spaces := strings.Repeat(" ", (fillWidth-currentWidth)/2)
		fmt.Print(spaces, starts, "\n")
	}
}
