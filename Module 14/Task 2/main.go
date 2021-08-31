package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cords struct {
	x int
	y int
}

func (c *cords) random() {
	c.x = rand.Intn(1001) - 500
	c.y = rand.Intn(1001) - 500
}

func (c *cords) update() {
	c.x = 2*c.x + 10
	c.y = -3*c.y - 5
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a := new(cords)
	b := new(cords)
	c := new(cords)

	fmt.Println("Геренерируем координаты")

	a.random()
	b.random()
	c.random()

	fmt.Printf("Координаты a(x, y): %d %d\n", a.x, a.y)
	fmt.Printf("Координаты b(x, y): %d %d\n", b.x, b.y)
	fmt.Printf("Координаты c(x, y): %d %d\n", c.x, c.y)

	fmt.Println("Преобразуем координаты")

	a.update()
	b.update()
	c.update()

	fmt.Printf("Координаты a(x, y): %d %d\n", a.x, a.y)
	fmt.Printf("Координаты b(x, y): %d %d\n", b.x, b.y)
	fmt.Printf("Координаты c(x, y): %d %d\n", c.x, c.y)
}
