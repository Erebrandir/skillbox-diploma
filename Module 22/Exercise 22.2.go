/*Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)
Что нужно сделать
Заполните упорядоченный массив из 12 элементов и введите число. Необходимо реализовать поиск первого вхождения
заданного числа в массив. Сложность алгоритма должна быть минимальная.
Что оценивается Верность индекса.
При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе числа 2 программа должна вывести индекс 1.*/

package main

import "fmt"

const num = 12

func main() {
	var arr [num]int

	fmt.Println("Введите значения массива: ")
	for i := 0; i < num; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)

	fmt.Println("Введите искомое значение: ")
	var value int
	fmt.Scan(&value)

	index := find(arr, value)

	for i := 0; i < num; i++ {
		if index > 0 && arr[index] == arr[index-1] {
			index = index - 1
		} else {
			index = 0
		}
	}
	fmt.Printf("Value: %v\n", value)
	fmt.Printf("Index: %v\n", index)
}

func find(arr [num]int, value int) (index int) {
	index = -1
	min := 0
	max := num - 1
	for max >= min {
		middle := (max + min) / 2
		if arr[middle] == value {
			index = middle
			break
		} else if arr[middle] > value {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	return
}

//func find(arr [num]int, value int) (index int) {
//	index = -1
//	for i := 0; i < num; i++ {
//		if arr[i] == value {
//			index = i
//			break
//		} else {
//			index = 0
//		}
//	}
//	return
//}
