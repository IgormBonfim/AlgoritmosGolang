package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := createArray(5)

	fmt.Printf("Array não ordenado: %d \n", array)

	arrayOrdenado := quicksort(array)

	fmt.Printf("Array ordenado: %d \n", arrayOrdenado)
}

func quicksort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	pivo := array[len(array)/2]
	var menores []int
	var maiores []int
	var iguais []int

	for _, value := range array {
		if value < pivo {
			menores = append(menores, value)
		} else if value > pivo {
			maiores = append(maiores, value)
		} else {
			iguais = append(iguais, value)
		}
	}
	fmt.Printf("Ordenando, menores: %d, pivo: %d, maiores: %d \n", menores, pivo, maiores)

	return append(quicksort(menores), append(iguais, quicksort(maiores)...)...)
}

func createArray(width int) []int {

	array := make([]int, width)

	for i := 0; i < width; i++ {
		array[i] = rand.Intn(100)
	}

	return array
}
