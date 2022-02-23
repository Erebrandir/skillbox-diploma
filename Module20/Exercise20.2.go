/*Задание 2. Умножение матриц
Что нужно сделать
Напишите функцию, умножающую две матрицы размерами 3 × 5 и 5 × 4.*/

package main

import "fmt"

func multiplicationMatrix(a [rowsFirst][colsFirst]int, b [rowsSecond][colsSecond]int) [rowsFirst][colsSecond]int {
	var multiplicationnmat [rowsFirst][colsSecond]int
	for i := 0; i < rowsFirst; i++ {
		for j := 0; j < colsSecond; j++ {
			for k := 0; k < colsFirst; k++ {
				multiplicationnmat[i][j] = multiplicationnmat[i][j] + a[i][k]*b[k][j]
			}
		}
	}
	return multiplicationnmat
}

const rowsFirst = 3
const colsFirst = 5
const rowsSecond = 5
const colsSecond = 4

func main() {
	var matrix [rowsFirst][colsSecond]int
	matrixFirst := [rowsFirst][colsFirst]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}

	matrixSecond := [rowsSecond][colsSecond]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}

	matrix = multiplicationMatrix(matrixFirst, matrixSecond)
	fmt.Println("Умножение матриц будет равно: ")
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
