package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Данная программа запрашивает у вас 3 числа и говорит, есть ли среди них число больше 5")

	flag := false

	for i := 1; i < 4; i++ {
		var input int
		fmt.Printf("Введите число √%d\n", i)
		_, err := fmt.Scan(&input)
		checkError(err)

		if input > 5 {
			flag = true
		}
	}

	if flag {
		fmt.Println("Среди введенных чисел есть хотя бы одно число больше 5")
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
