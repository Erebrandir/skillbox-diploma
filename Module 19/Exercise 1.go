/*Задание 1. Слияние отсортированных массивов
Что нужно сделать Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
Советы и рекомендации. Обратите внимание на размеры массивов.
Что оценивается. Правильность размеров. Правильный порядок элементов в конечном массиве.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arrayFirst := [5]int{1, 3, 5, 7, 9}
	arraySecond := [4]int{2, 4, 6, 8}

	var arrayUnion [9]int

	array := arrayUnion[:0]
	array = append(array, arrayFirst[:]...)
	array = append(array, arraySecond[:]...)

	sort.Ints(array)
	fmt.Println(array)
}
