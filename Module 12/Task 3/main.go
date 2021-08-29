package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Программа проверяет уровни доступа")

	file, err := os.OpenFile("testfile.txt", os.O_CREATE, 0444)

	if err != nil {
		panic(err)
	}

	fmt.Println("Успешно создан новый файл testfile.txt")

	if err = file.Close(); err != nil {
		panic(err)
	}

	fmt.Println("Открываем новый файл с целью записи")

	file, err = os.OpenFile("testfile.txt", os.O_WRONLY, 0444)

	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			fmt.Println("Не удалось открыть файл для записи. Всё отлично: тест пройден")
			return
		}

		panic(err)
	}
}
