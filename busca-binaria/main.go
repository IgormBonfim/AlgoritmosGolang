package main

import (
	"fmt"
)

func main() {
	array := createArray(240000)
	
	position, err := buscaBinaria(array, 1);

	if err != nil {
		fmt.Println(err.Error());
	}
    fmt.Println("Número encontrado na posição: \n", position);
}

func buscaBinaria(list []int, target int) (int, error) {
	low := 0;
	high := len(list) - 1;
	
	for low <= high {
		half:= (low + high)/2;
		temp := list[half];

		fmt.Println("Buscando no indice: \n", half)

		if temp == target {
			return half, nil;
		}

		if temp > target {
			high = half - 1;
		} else {
			low = half + 1
		}
	}
	return -1, NewError("Número não encontrado no array")
}

func createArray(width int) []int {

	array := make([]int, width);

	for i := 0; i < width; i++ {
		array[i] = i+1;
	}

	return array
}