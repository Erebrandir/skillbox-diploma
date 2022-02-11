/*Задание 2. Поиск символов в нескольких строках
Что нужно сделать
Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune, а
возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово в
предложении i (строку надо разбить на слова и взять последнее).

Советы и рекомендации
Не забудьте проверить, что вы получили больше чем 0 аргументов.

Пример вывода результата в первом элементе массива
'H' position 0
'E' position 1
'L' position 9*/

package main

import (
	"fmt"
	"strings"
)

func parseTest(sentences [4]string, chars [5]rune) (a []rune, j int) {
	for i := 0; i < len(sentences); i++ {
		phraseForPrint := strings.Split(sentences[i], " ")
		fmt.Println("\nPhrase is:", phraseForPrint)

		words := strings.Replace(sentences[i], " ", " ", -1)
		words = strings.ToLower(words)
		a := []rune(words)

		m := make(map[rune]int)
		for j := 0; j < len(a); j++ {
			for l := 0; l < len(chars); l++ {
				if a[j] == chars[l] {
					m[chars[l]] = j
					//fmt.Printf("%c position %v\n", chars[l], j)
				}
			}
		}
		for key, value := range m {
			fmt.Printf("%c position %v\n", key, value)
		}
	}
	return
}

func main() {
	sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [5]rune{'h', 'e', 'l', 'п', 'м'}

	parseTest(sentences, chars)
}
