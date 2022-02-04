package main

import "fmt"

func reverseArray(input []int) []int {
	var outputArray []int

	for i := len(input) - 1; i >= 0; i-- {
		outputArray = append(outputArray, input[i])
	}

	return outputArray //gdhfghfgjfjgjh
}
func main() {
	arrayOfNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arrayOfNumbers)
	fmt.Println(reverseArray(arrayOfNumbers))
}
