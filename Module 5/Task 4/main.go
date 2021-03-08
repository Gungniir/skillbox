package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Данная программа подскажет вам, сможете ли вы оплатить без сдачи")

	var cost int
	fmt.Println("Введите сумму в рублях")
	_, err := fmt.Scan(&cost)
	failOnError(err)

	var coins [3]int
	fmt.Println("Введите через пробел номинал ваших трех монет")
	_, err = fmt.Scanf("%d%d%d", &coins[0], &coins[1], &coins[2])
	failOnError(err)

	if coins[0]+coins[1]+coins[2] < cost {
		fmt.Println("У вас не хватит денег")
	} else if coins[0] == cost || coins[1] == cost || coins[2] == cost ||
		coins[0]+coins[1] == cost || coins[0]+coins[2] == cost || coins[1]+coins[2] == cost ||
		coins[0]+coins[1]+coins[2] == cost {
		fmt.Println("Вы можете оплатить без сдачи")
	} else {
		fmt.Println("Вы не можете оплатить без сдачи")
	}
}

func failOnError(e error) {
	if e == nil {
		return
	}

	fmt.Printf("Произошла ошибка: %s\n", e.Error())
	os.Exit(1)
}
