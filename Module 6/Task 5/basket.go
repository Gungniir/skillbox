package main

import "fmt"

const basketsCount = 3

type Basket struct {
	Max   int
	Value int
}

type Baskets [basketsCount]Basket

func (b *Baskets) isFull() bool {
	for _, i := range b {
		if i.Max != i.Value {
			return false
		}
	}

	return true
}

func (b *Baskets) printStatus() {
	for n, i := range b {
		fmt.Printf("Корзина №%2d: %2d/%2d\n", n+1, i.Value, i.Max)
	}
}
