/*Задание 1. Чётные и нечётные
Что нужно сделать
Напишите функцию, которая принимает массив чисел, а возвращает два массива: один из чётных чисел, второй из нечётных.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func bubbleSort(input [n]int) [n]int {
	for i := n; i > 0; i-- {
		for j := 1; j < i; j++ {
			if input[j-1] > input[j] {
				input[j], input[j-1] = input[j-1], input[j]
			}
		}
	}
	return input
}

func arrayEvenOdd(arr [n]int) (evenArr []int, notEvenArr []int) {
	for _, j := range arr {
		if j%2 == 0 {
			evenArr = append(evenArr, j)
		} else {
			notEvenArr = append(notEvenArr, j)
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [n]int
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(10 * n)
	}

	fmt.Println("Unsorted array -", a)
	b := bubbleSort(a)
	fmt.Println("Sorted array -", b)
	fmt.Println("\nEven array and Odd array:")
	fmt.Println(arrayEvenOdd(b))
}
