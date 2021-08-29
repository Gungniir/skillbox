package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Программа записывает ваши сообщения в файл")

	scanner := bufio.NewScanner(os.Stdin)
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "выход" {
			fmt.Println("Завершение работы...")
			break
		}

		now := time.Now()

		_, err := fmt.Fprintf(file, "%d-%02d-%02d %02d:%02d:%02d %s\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), msg)

		if err != nil {
			panic(err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Произошла ошибка:", err)
	}
}
