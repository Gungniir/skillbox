package main

import (
	"fmt"
)

type House struct {
	PeoplePerFloor [floorsCount]int
}

func (h House) isEmpty() bool {
	for _, people := range h.PeoplePerFloor {
		if people > 0 {
			return false
		}
	}
	return true
}

func (h House) printStatus() {
	for n, people := range h.PeoplePerFloor {
		fmt.Printf("Этаж %2d: %2d ожидающих\n", n+1, people)
	}
}

const floorsCount = 24
const liftSize = 2

func main() {
	fmt.Println("Программа показывает алгоритм сбора всех людей лифтом")

	house := House{PeoplePerFloor: [floorsCount]int{
		0, // 1  этаж
		0, // 2  этаж
		0, // 3  этаж
		3, // 4  этаж
		0, // 5  этаж
		0, // 6  этаж
		3, // 7  этаж
		0, // 8  этаж
		0, // 9  этаж
		3, // 10 этаж
		//etc
	}}

	fmt.Println("---Дом---")
	house.printStatus()
	fmt.Printf("---------\n\n")

	// Цикл лифта. Когда q = 0, то он едет на первый этаж, выгружает всех и идет по новой собирать людей
	for q := liftSize; ; q = liftSize {
		if house.isEmpty() {
			break
		}

		fmt.Println("---Лифт поехал----")

		// Цикл сбора людей по этажам
		for i := floorsCount - 1; i > 0; i-- {
			// Создаем ссылку, чтобы меньше писать
			people := &house.PeoplePerFloor[i]

			// Пустой этаж, пропускаем
			if *people == 0 {
				continue
			}

			// Людей на этаже достаточно, чтобы заполнить лифт => заполняемся по полной и едем на первый этаж
			if *people >= q {
				*people -= q

				fmt.Printf("Забираем %d людей с %d этажа. Осталось на этаже %d\n", q, i+1, *people)
				break
			}

			// Людей недостаточно, чтобы заполнить лифт => едем далее
			if *people < q {
				q -= *people

				fmt.Printf("Забираем всех (%d) людей с %d этажа\n", *people, i+1)

				*people = 0

				// Здесь continue можно и не писать, т.к. мы и так перейдем на следующую итерацию
				continue
			}
		}

		fmt.Println("---Лифт приехал---")
		fmt.Println("----Ожидающих-----")
		house.printStatus()
		fmt.Println("------------------")
	}

	fmt.Println("Всё, всех доставили :)")
}
