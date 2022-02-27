/*Цель задания
Научиться работать с композитными типами данных: структурами и картами

Что нужно сделать
Напишите программу, которая считывает ввод с stdin, создаёт структуру student и записывает указатель на структуру в
хранилище map[studentName] *Student.
Программа должна получать строки в бесконечном цикле, создать структуру Student через функцию newStudent, далее
сохранить указатель на эту структуру в map, а после получения EOF (ctrl + z) вывести на экран имена всех студентов
из хранилища. Также необходимо реализовать методы put, get.
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

func (s Student) info() string {
	return s.name + " " + strconv.Itoa(s.age) + " " + strconv.Itoa(s.grade)
}

func main() {
	studentMap := make(map[string]*Student, 0)
	fmt.Println("Введите имя студента, возраст и курс:")

	var count = 1
	var reader = bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("----------------------")
			fmt.Println("Студенты из хранилища:")
			break
		}
		lineFields := strings.Fields(line)

		if len(lineFields) < 3 {
			fmt.Print("Необходимо ввести имя, возраст и курс! Попробуйте снова.\n")
			continue
		}

		studentName := lineFields[0]
		studentAge, errAge := strconv.Atoi(lineFields[1])
		studentGrade, errGrade := strconv.Atoi(lineFields[2])

		if errAge != nil || errGrade != nil {
			fmt.Print("Ошибка возраста студента и курса! Попробуйте снова.\n")
			continue
		}

		student := Student{
			name:  studentName,
			age:   studentAge,
			grade: studentGrade,
		}

		if _, err := get(studentMap, student.name); err != nil {
			put(studentMap, &student)
		} else {
			fmt.Print("Студент с таким именем уже есть в хранилище! Попробуйте снова.\n")
		}
	}

	for _, value := range studentMap {
		fmt.Printf("%v. %s\n", count, value.info())
		count++
	}
}

func put(list map[string]*Student, value *Student) (int, error) {
	fmt.Println(list[value.name])
	if list[value.name] == nil {
		return -1, errors.New("Ошибка добавления в хранилище!")
	} else {
		return 0, nil
	}
}

func get(list map[string]*Student, name string) (*Student, error) {
	if list[name] == nil {
		return nil, errors.New("Студент не найден в хранилище!")
	} else {
		return list[name], nil
	}
}
