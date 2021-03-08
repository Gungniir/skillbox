package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Я умный-умный помощник, встроенный в кассу. Помогаю учесть такие тонкие моменты начисления " +
		"скидок и надбавок в зависимости от дня недели, суммы заказа и количества посетителей в чеке!")

	var weekday string
	fmt.Println("Введите день недели целым словом или двубуквенным сокращением, например: \"Понедельник\" или \"ПН\"")
	_, err := fmt.Scan(&weekday)
	checkError(err)
	weekday = strings.ToUpper(weekday)

	var guestsCount int
	fmt.Println("Сколько гостей в одном чеке?")
	_, err = fmt.Scan(&guestsCount)
	checkError(err)

	// Чтобы проводить все операции без погрешности, храним всё в целом типе. Т.е. в копейках
	var cost int
	fmt.Println("Введите сумму чека в копейках")
	_, err = fmt.Scan(&cost)
	checkError(err)

	fmt.Println("--------")
	// Здесь будем хранить все надбавки (-) или скидки (+)
	discount := 0

	// Надбавка для больших компаний
	if guestsCount > 5 {
		fmt.Printf("Компания большая, начислим надбавку в 10%%: %d копеек.\n", cost*10/100)
		discount -= cost * 10 / 100 // 10% надбавка => знак отрицательный
	}

	// Скидка по понедельникам
	if weekday == "ПН" || weekday == "ПОНЕДЕЛЬНИК" {
		fmt.Printf("Понедельник день тяжелый, дарим скидку в 10%%: %d копеек.\n", cost*10/100)
		discount += cost * 10 / 100 // 10% скидка => знак положительный
	}

	// Скидка по пятницам при чеке больше 10 000
	if (weekday == "ПТ" || weekday == "ПЯТНИЦА") && cost > 10_000 {
		fmt.Printf("Богатый чек, да еще и в пятницу! Дарим скидку в 5%%: %d копеек.\n", cost*5/100)
		discount += cost * 5 / 100 // 5% скидка
	}

	total := cost - discount

	fmt.Println("--------")
	fmt.Printf("Итого к оплате: %d.%d руб.\n", total/100, total%100)
}

func checkError(err error) {
	if err == nil {
		return
	}

	fmt.Println("Произошла ошибка... Попробуйте перезапустить программу")
	os.Exit(1)
}
