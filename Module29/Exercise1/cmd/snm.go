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

import "fmt"

func main() {
	for {
		fmt.Println("Введите число или кодовое слово - стоп: ")
		var number int
		fmt.Scan(&number)

		if string(number) != "стоп" {
			fc := squareOfnNumber(number)
			sc := multiplicationOfNumber(<-fc)

			fmt.Println(<-sc)
			fmt.Println()
		} else {
			fmt.Println("Вы ввели стоп слово, работа программы остановлена.")
			break
		}
	}
}

func squareOfnNumber(num int) chan int {
	firstChan := make(chan int)
	go func() {
		res := num * num
		firstChan <- res
		close(firstChan)
	}()
	return firstChan
}

func multiplicationOfNumber(num int) chan int {
	secondChan := make(chan int)
	fmt.Println(num)
	go func() {
		numNew := 2 * num
		secondChan <- numNew
		close(secondChan)
	}()
	return secondChan
}
