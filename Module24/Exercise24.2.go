/*Задание 2. Анонимные функции
Что нужно сделать
Напишите анонимную функцию, которая на вход получает массив типа integer, сортирует его пузырьком и переворачивает
(либо сразу сортирует в обратном порядке, как посчитаете нужным).*/

package main

import "fmt"

const size = 10

func main() {
	array := [size]int{3, 4, 1, 2, 5, 7, -1, 0, 9, -2}

	fmt.Printf("Исходный массив: %v\n", array)

	reverseBubbleSort := func(items [size]int) [size]int {
		for i := 0; i < size; i++ {
			for j := size - 1; j >= i+1; j-- {
				if items[j] > items[j-1] {
					items[j], items[j-1] = items[j-1], items[j]
				}
			}
		}
		return items
	}
	fmt.Printf("Отсортированный массив: %v", reverseBubbleSort(array))
}
