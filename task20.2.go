package main

import (
	"fmt"
)

const columnOne = 3
const lineOne = 5
const columnTwo = 5
const lineTwo = 4

func matrixMultiplication(a [columnOne][lineOne]int, b [columnTwo][lineTwo]int) (c [columnOne][lineTwo]int) {

	var sum int
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			sum = 0
			for k := 0; k < lineOne; k++ {
				sum += a[i][k] * b[k][j]
			}
			c[i][j] = sum
		}
	}
	return
}

func main() {
	fmt.Println("20.5 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 2. Умножение матриц")
	fmt.Println("------------")

	matrix := [columnOne][lineOne]int{
		{4, 4, 1, 2, 3},
		{1, 5, 7, 8, 9},
		{1, 5, -7, 5, 6},
	}

	matrixTwo := [columnTwo][lineTwo]int{
		{1, 3, 6, 2},
		{2, 6, 1, 4},
		{1, 3, 4, 5},
		{1, 5, 1, 2},
		{4, 5, 3, 2},
	}

	fmt.Println(matrixMultiplication(matrix, matrixTwo))
}
