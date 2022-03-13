/*
Что нужно сделать
Реализуйте паттерн-конвейер:
Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
Произведение: следующая горутина умножает квадрат числа на 2.
При вводе «стоп» выполнение программы останавливается.

Советы и рекомендации
Воспользуйтесь небуферизированными каналами и waitgroup.
*/

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for {
		c := make(chan int, 2)
		fmt.Println("\nВведите число или кодовое слово - стоп: ")
		var input string
		fmt.Scan(&input)

		if input != "стоп" {
			number, err := strconv.Atoi(input)
			if err != nil {
				fmt.Print("Ошибка обработки введенных данных.\n")
				break
			} else {
				wg.Add(2)
				go func() {
					defer wg.Done()
					squareOfNumber(number, c)
				}()

				go func() {
					defer wg.Done()
					defer close(c)
					multiplicationOfNumber(c)
				}()
			}
		} else {
			fmt.Println("Вы ввели стоп слово, работа программы остановлена.")
			break
		}

		wg.Wait()
		for val := range c {
			fmt.Println(val)
		}
	}
}

func squareOfNumber(num int, chanel chan int) {
	res := num * num
	chanel <- res
}

func multiplicationOfNumber(chanel chan int) {
	num := <-chanel
	numMultip := num * 2
	chanel <- num
	chanel <- numMultip
}
