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
	storages := *storage.NewUniversity()
	stdin := bufio.NewReader(os.Stdin)

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
		storages.Get()

		err := storages.Put(std)
		if err != nil {
			fmt.Print("Студент с таким именем уже есть в хранилище!\n")
		} else {
			continue
		}
	}
	count := 1
	for _, value := range storages.Get() {
		fmt.Printf("%v. %s\n", count, value.Info())
		count++
	}
}
