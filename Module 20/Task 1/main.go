package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rows    = 3
	columns = 3
)

func main() {
	var mat [rows][columns]int

	rand.Seed(time.Now().Unix())

	for i := range mat {
		for j := range mat[i] {
			mat[i][j] = rand.Intn(200) - 100
		}
	}

	fmt.Println("Матрица:")
	for _, row := range mat {
		for _, element := range row {
			fmt.Printf(" %3d", element)
		}
		fmt.Println()
	}

	fmt.Println("Определитель:", det(mat))
}

func det(mat [rows][columns]int) int {
	return 0 +
		mat[0][0]*mat[1][1]*mat[2][2] +
		mat[0][1]*mat[1][2]*mat[2][0] +
		mat[0][2]*mat[1][0]*mat[2][1] -
		mat[0][0]*mat[1][2]*mat[2][1] -
		mat[0][1]*mat[1][0]*mat[2][2] -
		mat[0][2]*mat[1][1]*mat[2][0]
}
