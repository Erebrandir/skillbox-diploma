/*Задание 2. Анонимные функции

Что нужно сделать
Напишите функцию, которая на вход принимает функцию вида A func (int, int) int, а внутри оборачивает и вызывает её при выходе (через defer).
Вызовите эту функцию с тремя разными анонимными функциями A. Тела функций могут быть любыми, но главное, чтобы все три выполняли разное действие.*/

package main

import "fmt"

func funcUseDefer(x int, y int, A func(int, int) int) int {
	defer A(x, y)
	fmt.Println("Something happend...")
	return A(x, y)
}

func main() {
	fmt.Println(funcUseDefer(2, 4, func(x int, y int) int { return x + y }))
	fmt.Println(funcUseDefer(3, 5, func(x int, y int) int { return x * y }))
	fmt.Println(funcUseDefer(4, 6, func(x int, y int) int { return x - y }))
}

//func main() {
//	fmt.Println(calculate(func(a, b int) int { return a + b }, 10, 5))
//	fmt.Println(calculate(func(a, b int) int { return a * b }, 10, 5))
//	fmt.Println(calculate(func(a, b int) int { return a / b }, 10, 5))
//}
//
//func calculate(operation func(int, int) int, a int, b int) (result int) {
//	defer func() {
//		result = operation(5, 4)
//	}()
//	result = 0 // эта строка демонстрирует, что функция в defer меняет значение при выходе из функции, ее нужно удалить
//	return
//}
