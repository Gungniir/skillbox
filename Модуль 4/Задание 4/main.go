package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Данная программа запрашивает у вас 3 числа и говорит, сколько среди них число больше 5")

	passed := 0

	for i := 1; i < 4; i++ {
		var input int
		fmt.Printf("Введите число √%d\n", i)
		_, err := fmt.Scan(&input)
		checkError(err)

		if input > 5 {
			passed++
		}
	}

	if passed > 0 {
		fmt.Printf("Среди введенных чисел есть числа больше 5, а именно их: %d\n", passed)
	} else {
		fmt.Println("Среди введенных чисел нет числа больше 5")
	}
}

func checkError(err error) {
	if err == nil {
		return
	}

	fmt.Println("Произошла ошибка... Попробуйте перезапустить программу")
	os.Exit(1)
}
