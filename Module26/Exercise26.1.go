/*Цель задания
Написать программу аналог cat.

Что нужно сделать
Программа должна получать на вход имена двух файлов, необходимо конкатенировать их содержимое, используя strings.Join.*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func unpackingFile(fileName string) string {
	fileCopy, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fileCopy.Close()

	resultBytes, err := ioutil.ReadAll(fileCopy)
	if err != nil {
		panic(err)
	}
	file := string(resultBytes)
	return file
}

func main() {
	firstFileName := "first.txt"
	secondFileName := "second.txt"

	firstFile := unpackingFile(firstFileName)
	fmt.Println("Содержимое первого файла: ", firstFile)

	secondFile := unpackingFile(secondFileName)
	fmt.Println("Содержимое первого файла: ", secondFile)

	sliceFiles := []string{firstFile, secondFile}
	content := strings.Join(sliceFiles, "; ")

	outFile := "third.txt"
	err := ioutil.WriteFile(outFile, []byte(content), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nСодержимое файла <%s> заполнено сконкатенированным содержанием файлов <%s> и <%s>", outFile, firstFile, secondFile)
	fmt.Printf("\n<%s> + <%s>  содержимое: <%s>", firstFile, secondFile, content)
}
