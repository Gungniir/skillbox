package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "a10 10 20b 20 30c30 30 dd"

	for _, s := range strings.Split(str, " ") {
		_, err := strconv.Atoi(s)

		if err == nil {
			fmt.Println(s)
		}
	}
}
