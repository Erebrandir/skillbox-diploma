/*Задание 1. Подсчёт чётных и нечётных чисел в массиве
Что нужно сделать
Заполните массив неупорядоченными числами на основе генератора случайных чисел. Введите число. Программа
должна найти это число в массиве и вывести, сколько чисел находится в массиве после введённого. При отсутствии
введённого числа в массиве — вывести 0. Для удобства проверки реализуйте вывод массива на экран.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [n]int
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(10 * n)
	}
	fmt.Println(a)

	fmt.Println("Введите искомое значение: ")
	var value int
	fmt.Scan(&value)

	index := find(a, value)
	fmt.Printf("Value: %v\n", value)
	fmt.Printf("Index: %v\n", index)

	if index != 0 {
		fmt.Printf("Чисел находится в массиве после введённого: %v\n", 9-index)
	}
}

func find(a [n]int, value int) (index int) {
	index = -1
	for i := 0; i < n; i++ {
		if a[i] == value {
			index = i
			break
		} else {
			index = 0
		}
	}
	return
}
