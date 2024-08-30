package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := createArray(15);

	fmt.Printf("Array não ordenado: %d \n", array)

	arrayOrdenado := make([]int, len(array))

	for i := range array {
		menor := selectionSort(array);
		arrayOrdenado[i] = array[menor];
		array = remove(array, menor);
	}
	fmt.Printf("Array ordenado: %d \n", arrayOrdenado)
}

func selectionSort(array []int) int {

	menor := array[0];
	indiceMenor := 0;

	for i := 1; i < len(array); i++ {
		if array[i] < menor {
			menor = array[i];
			indiceMenor = i;
		}
	}

	return indiceMenor;
}

func remove(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func createArray(width int) []int {

	array := make([]int, width)

	for i := 0; i < width; i++ {
		array[i] = rand.Intn(100)
	}

	return array
}
