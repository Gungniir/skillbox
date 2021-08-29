package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	fmt.Println("Программа записывает ваши сообщения в файл, а затем читает файл")

	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]byte, 0)

	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "выход" {
			fmt.Println("Завершение работы...")
			break
		}

		now := time.Now()

		line := []byte(fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d %s\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), msg))

		lines = append(lines, line...)
	}
	err := ioutil.WriteFile("output.txt", lines, fs.ModeAppend)

	if err != nil {
		panic(err)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Произошла ошибка:", err)
	}

	fmt.Println("--Читаем файл--")

	data, err := ioutil.ReadFile("output.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
