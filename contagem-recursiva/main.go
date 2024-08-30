package main

import (
	"fmt"
)

func main() {
	array := createArray(5)

	fmt.Printf("Contando número de itens no array: %d \n", array)
	resultado := count(array)
	fmt.Printf("Número de itens no array: %d \n", resultado)
	
}

func count(array []int) int {
	if len(array) < 1 {
		return 0
	}
	
	arrayRestante := array[1:]

	return 1 + count(arrayRestante)
}

func createArray(size int) []int {

	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = i + 1
	}

	return array
}
