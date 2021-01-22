package main

import (
	"fmt"
	"os"
)

const pattern = " *"

func main() {
	fmt.Println("Программа выводит доску формата X на X. Укажите желаемый X :)")
	fmt.Print("X=")

	var x int
	_, err := fmt.Scan(&x)

	if err != nil {
		fmt.Println("Не удалось считать число")
		os.Exit(1)
	}

	if x < 1 {
		fmt.Println("X должен быть положительным")
		os.Exit(1)
	}

	// Способ 1. Экономия на цп

	// Создаём строку, включающую в себя две строки одновременно
	//row := strings.Repeat(pattern, x/2+1)
	//
	//for i := 0; i < x; i++ {
	//	fmt.Println(row[i%2 : x+i%2])
	//}

	// Способ 2. Экономия памяти
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			fmt.Print(string(pattern[(i%2+j%2)%2]))
		}
		fmt.Print("\n")
	}

}
