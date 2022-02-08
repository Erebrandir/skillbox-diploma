/*Задание 1. Расчёт по формуле
Что нужно сделать
Напишите функцию, производящую следующие вычисления.
S = 2 × x + y ^ 2 − 3/z, где x — int16, y — uint8, a z — float32.
Тип S должен быть во float32.*/

package main

import (
	"fmt"
	"math"
)

func calculatingTheEquation(x int16, y uint8, z float32) float32 {
	s := 2*float32(x) + float32(math.Pow(float64(y), 2)) - 3/z
	return s
}

func main() {
	fmt.Println(calculatingTheEquation(-5, 5, 11.5))
}
