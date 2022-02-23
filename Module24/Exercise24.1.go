/*Задание 1. Сортировка вставками
Что нужно сделать
Напишите функцию, сортирующую массив длины 10 вставками.*/

package main

import "fmt"

func main() {
	arr := []int{3, 4, 1, 2, 5, 7, -1, 0, 9, -2}
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for ; j >= 1 && arr[j-1] > x; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = x
	}
}
