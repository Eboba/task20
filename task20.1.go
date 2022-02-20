package main

import (
	"fmt"
)

func determinant(A [3][3]int) int {
	// a11·a22·a33 + a12·a23·a31 + a13·a21·a32 - a13·a22·a31 - a11·a23·a32 - a12·a21·a33
	return A[0][0]*A[1][1]*A[2][2] + A[0][1]*A[1][2]*A[2][0] + A[0][2]*A[1][0]*A[2][1] - A[0][2]*A[1][1]*A[2][0] - A[0][0]*A[1][2]*A[2][1] - A[0][1]*A[1][0]*A[2][2]
}
func main() {
	fmt.Println("20.5 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 1. Подсчёт определителя")
	fmt.Println("------------")

	var matrix [3][3]int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Введите значение для матрицы строка: %v столбец %v:", i+1, j+1)
			fmt.Println("")
			fmt.Scan(&matrix[i][j])
		}
	}

	fmt.Println("Для матрицы:")
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("определитель =", determinant(matrix))
}
