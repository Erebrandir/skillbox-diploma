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

	fmt.Println("Пожалуйста, введите имя, возраст и курс студентов с переносом строки:")

	for {
		fmt.Print("->: ")
		line, errEOF := stdin.ReadString('\n')

		if errEOF == io.EOF {
			fmt.Println()
			fmt.Println("Конец ввода! Список студентов:")
			break
		}

		lineFields := strings.Fields(line)

		if len(lineFields) < 3 {
			fmt.Print("Необходимо ввести имя, возраст и грейд! Пожалуйста, попробуйте снова...\n")
			continue
		}

		studentName := lineFields[0]
		studentAge, errAge := strconv.Atoi(lineFields[1])
		studentGrade, errGrade := strconv.Atoi(lineFields[2])

		if errAge != nil || errGrade != nil {
			fmt.Print("Ошибка при обработке возраста студента или его грейда! Пожалуйста, попробуйте снова...\n")
			continue
		}

		std := student.NewStudent(studentName, studentAge, studentGrade)

		_, err := storage.Get(studentName)
		if err != nil {
			storage.Put(std)
		} else {
			fmt.Print("Студент с таким именем уже есть в хранилище! Попробуйте снова...\n")
		}
	}

	for _, value := range storage {
		fmt.Printf("-> %s\n", value.Info())
	}
}
