package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Программа полностью читает файл output.txt")

	file, err := os.Open("output.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Размер файла: %d байт\n", stat.Size())

	//io.ReadAll() - я так понимаю, задание заключается в том, что мы не должны испольовать ReadAll

	buf := make([]byte, stat.Size())

	if _, err := io.ReadFull(file, buf); err != nil {
		panic(err)
	}

	fmt.Print(string(buf))
}
