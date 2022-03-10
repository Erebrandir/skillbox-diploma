package main

import (
	"awesomeProject/Module28/pkg/storage"
	"awesomeProject/Module28/pkg/student"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	storage := storage.NewStorage()
	stdin := bufio.NewReader(os.Stdin)
	var count = 1

	fmt.Println("Введите имя студента, возраст и курс:")

	for {
		line, errEOF := stdin.ReadString('\n')
		if errEOF == io.EOF {
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

		std := student.NewStudent(studentName, studentAge, studentGrade)

		_, err := storage.Get(studentName)
		if err != nil {
			storage.Put(std)
		} else {
			fmt.Print("Студент с таким именем уже есть в хранилище!\n")
		}
	}

	for _, value := range storage {
		fmt.Printf("%v. %s\n", count, value.Info())
		count++
	}
}
