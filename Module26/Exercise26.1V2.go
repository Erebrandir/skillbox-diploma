package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	switch len(os.Args) {
	case 2:
		firstFile := os.Args[1]
		printFileContent(&firstFile)
	case 3:
		firstFile := os.Args[1]
		secondFile := os.Args[2]
		content := joinFilesContent(&firstFile, &secondFile)
		fmt.Printf("<%s> + <%s> содержимое: <%s>", firstFile, secondFile, content)
	case 4:
		firstFile := os.Args[1]
		secondFile := os.Args[2]
		outFile := os.Args[3]
		content := joinFilesContent(&firstFile, &secondFile)
		err := ioutil.WriteFile(outFile, []byte(content), 0644)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Содержимое файла <%s> заполнено сконкатенированным содержанием файлов <%s> и <%s>", outFile, firstFile, secondFile)
	default:
		fmt.Println("Программу необходимо запустить с аргументами командной строки!")
		return
	}
}

func joinFilesContent(fileNameFirst, fileNameSecond *string) string {
	fileFirstContent, err := ioutil.ReadFile(*fileNameFirst)
	if err != nil {
		panic(err)
	}

	fileSecondContent, err := ioutil.ReadFile(*fileNameSecond)
	if err != nil {
		panic(err)
	}

	return strings.Join([]string{string(fileFirstContent), string(fileSecondContent)}, "; ")
}

func printFileContent(fileName *string) {
	content, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("<%s> содержимое: <%s>", *fileName, string(content))
}
