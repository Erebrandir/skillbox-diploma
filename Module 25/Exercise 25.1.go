/*Задача
Цель задания
Написать программу для нахождения подстроки в кириллической подстроке. Программа должна запускаться с помощью команды:
go run main.go --str "строка для поиска" --substr "поиска"
Для реализации такой работы с флагами воспользуйтесь пакетом flags, а для поиска подстроки в строке вам понадобятся руны.

Что нужно сделать
Спроектировать алгоритм поиска подстроки. Определить строку и подстроку, используя флаги.
Написать алгоритм реализацию для работы со строками UTF-8 (для этого необходимо воспользоваться рунами).*/

//Решение здорового человека
//package main

//import (
//	"fmt"
//	"strings"
//)

//func main() {
//	present := strings.Contains("abc", "ab")
//	fmt.Println(present)

//	present = strings.Contains("abc", "xyz")
//	fmt.Println(present)
//}

//Решение - не понятно зачем этот геморрой =)
package main

import (
	"flag"
	"fmt"
)

func containString(str string, substr string) bool {
	substrRunes := []rune(substr)
	strRunes := []rune(str)

	j := 0
	for _, r := range strRunes {
		if r == substrRunes[j] {
			j++
			if j == len(substrRunes) {
				return true
			}
		} else {
			j = 0
		}
	}
	return false
}

func main() {

	//Данные для проверки
	//Input привет мир!, вет
	//Output true

	//Input Программирование - это просто, вание
	//Output true

	//Input Программирование - это просто, корабль
	//Output false

	var str = flag.String("str", "Программирование - это просто", "Input str")
	var substr = flag.String("substr", "вание", "Input substr")
	flag.Parse()

	fmt.Println(containString(*str, *substr))
}
