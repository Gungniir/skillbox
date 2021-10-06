package main

import (
	"fmt"
	"math/rand"
)

const (
	rows1 = 3
	cols1 = 5
	rows2 = 5
	cols2 = 4
)

func main() {
	var a [rows1][cols1]int
	var b [rows2][cols2]int

	for i := range a {
		for j := range a[i] {
			a[i][j] = rand.Intn(20) - 10
		}
	}

	for i := range b {
		for j := range b[i] {
			b[i][j] = rand.Intn(20) - 10
		}
	}

	fmt.Println("Матрица a:")
	for _, row := range a {
		for _, element := range row {
			fmt.Printf(" %3d", element)
		}
		fmt.Println()
	}

	fmt.Println("Матрица b:")
	for _, row := range b {
		for _, element := range row {
			fmt.Printf(" %3d", element)
		}
		fmt.Println()
	}

	fmt.Println("Матрица a*b:")
	for _, row := range prod(a, b) {
		for _, element := range row {
			fmt.Printf(" %4d", element)
		}
		fmt.Println()
	}
}

func prod(a [rows1][cols1]int, b [rows2][cols2]int) [rows1][cols2]int {
	var result [rows1][cols2]int

	for rowIndex, row := range a {
		for columnIndex := 0; columnIndex < len(b[0]); columnIndex++ {
			for elementIndex, element := range row {
				result[rowIndex][columnIndex] += element * b[elementIndex][columnIndex]
			}
		}
	}

	return result
}
