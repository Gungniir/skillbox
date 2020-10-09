package main

import (
	"fmt"
	"os"
)

func main() {
	var x, y int

	fmt.Println("Введите координаты точки через пробел x и y. Например, 10 15")
	_, err := fmt.Scanf("%d%d", &x, &y)
	failOnError(err)

	if x == 0 || y == 0 {
		fmt.Println("Данная точка не принадлежит ни одной из четвертей плоскости")
	} else if x > 0 && y > 0 {
		fmt.Println("Данная точка принадлежит первой четверти плоскости")
	} else if x < 0 && y > 0 {
		fmt.Println("Данная точка принадлежит второй четверти плоскости")
	} else if x < 0 && y < 0 {
		fmt.Println("Данная точка принадлежит третьей четверти плоскости")
	} else if x > 0 && y < 0 {
		fmt.Println("Данная точка принадлежит четвертой четверти плоскости")
	}
}

func failOnError(e error) {
	if e == nil {
		return
	}

	fmt.Printf("Произошла ошибка: %s\n", e.Error())
	os.Exit(1)
}
