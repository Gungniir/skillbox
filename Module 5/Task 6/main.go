package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Данная программа решает за вас квадратные уравнения!")

	fmt.Println("Введите через пробел три коэффициента под ряд")

	var a, b, c float64
	_, err := fmt.Scanf("%f%f%f", &a, &b, &c)
	failOnError(err)

	d := b*b - 4*a*c
	if d > 0 {
		d = math.Sqrt(d)

		fmt.Println("x1 =", (-b+d)/(2*a))
		fmt.Println("x2 =", (-b-d)/(2*a))
	} else if d == 0 {
		fmt.Println("x =", -b/(2*a))
	} else {
		fmt.Println("Корней нет")
	}
}

func failOnError(e error) {
	if e == nil {
		return
	}

	fmt.Printf("Произошла ошибка: %s\n", e.Error())
	os.Exit(1)
}
