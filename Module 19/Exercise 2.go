/*Задание 2. Сортировка пузырьком
Что нужно сделать. Отсортируйте массив длиной шесть пузырьком.
Советы и рекомендации. Принцип сортировки пузырьком можно посмотреть на «Википедии», там есть наглядная демонстрация, или на YouTube.
Что оценивается. Правильность результата сортировки и отсутствие ошибок.
*/

package main

import "fmt"

const size = 6

func main() {
	sample := [size]int{3, 4, 5, 2, 1, 6}
	fmt.Println(bubbleSort(sample))
}

func bubbleSort(arr [size]int) [size]int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
