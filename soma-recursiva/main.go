package main

import (
	"fmt"
)

func main() {
	array := createArray(5)

	fmt.Printf("Somando array: %d \n", array)
	resultado := soma(array)
	fmt.Printf("Resultado da soma: %d \n", resultado)

}

func soma(array []int) int {
	if len(array) == 1 {
		return array[0]
	}

	proximoArray := array[1:]

	fmt.Printf("Somando: %d com %d \n", array[0], proximoArray)

	return array[0] + soma(proximoArray)
}

func createArray(size int) []int {

	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = i + 1
	}

	return array
}
