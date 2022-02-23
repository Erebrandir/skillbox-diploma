/*Задание 1. Подсчёт определителя
Что нужно сделать
Напишите функцию, вычисляющую определитель матрицы размером 3 × 3. */

package main

import "fmt"

const rows = 3
const cols = 3

func deterMat(a [rows][cols]int) int {
	var determinant int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ { //решение методом треугольника
			firstExpression := a[0][0]*a[1][1]*a[2][2] + a[0][1]*a[1][2]*a[2][0] + a[0][2]*a[2][1]*a[1][0]
			secondExpression := -a[0][2]*a[1][1]*a[2][0] - a[0][1]*a[2][2]*a[1][0] - a[0][0]*a[1][2]*a[2][1]
			determinant = firstExpression + secondExpression
		}
	}
	return determinant
}

func main() {
	matrix := [rows][cols]int{
		{9, 2, 5},
		{1, 4, 8},
		{6, 3, 7},
	}

	fmt.Println(deterMat(matrix))
}
