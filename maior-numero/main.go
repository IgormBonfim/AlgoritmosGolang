package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := createArray(5)

	fmt.Printf("Determinando o menor valor no array: %d \n", array)
	resultado := max(array)
	fmt.Printf("Maior valor: %d \n", resultado)
	
}

func max(array []int) int {
	if len(array) < 1 {
		return 0
	}
	
	arrayRestante := array[1:]
	
	maiorRestante := max(arrayRestante)
	
	fmt.Printf("Comparando: %d com : %d \n", array[0], maiorRestante)
	

	if array[0] > maiorRestante {
		return array[0]
	} else {
		return maiorRestante
	}
}

func createArray(size int) []int {

	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = rand.Intn(100)
	}

	return array
}
