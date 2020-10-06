package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Программа подсчитывает ваши баллы по трем предметам: (русский, " +
		"математика профильная, информатика) и говорит, смогли ли вы поступить :)")

	var russianScore int
	fmt.Println("Сколько вы получили баллов по русскому?")
	_, err := fmt.Scan(&russianScore)
	checkError(err)

	var mathScore int
	fmt.Println("Сколько вы получили баллов по профильной математике?")
	_, err = fmt.Scan(&mathScore)
	checkError(err)

	var infScore int
	fmt.Println("Сколько вы получили баллов по информатике?")
	_, err = fmt.Scan(&infScore)
	checkError(err)

	totalScore := russianScore + mathScore + infScore

	if totalScore >= 275 {
		fmt.Println("Отлично! Вы проходите!")
	} else {
		fmt.Printf("Эх, не чуть-чуть не хватило, нужно было набрать еще %d...\n", 275-totalScore)
	}
}

func checkError(err error) {
	if err == nil {
		return
	}

	fmt.Println("Произошла ошибка... Попробуйте перезапустить программу")
	os.Exit(1)
}
