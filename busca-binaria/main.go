	package main

	import (
		"fmt"
	)

	func main() {
		array := createArray(1024)
		
		position, count, err := buscaBinaria(array, 1024);

		if err != nil {
			fmt.Println(err.Error());
		}
		fmt.Printf("Número encontrado na posição: %d em %d iterações\n", position, count)
	}

	func buscaBinaria(list []int, target int) (int, int, error) {
		low := 0;
		high := len(list) - 1;
		count := 0;
		
		for low <= high {
			mid:= (low + high)/2;
			temp := list[mid];
			
			fmt.Printf("Buscando no indice: %d \n", mid)
			
			if temp == target {
				return mid, count, nil;
			}
			count++;
			
			if temp > target {
				high = mid - 1;
				} else {
					low = mid + 1
				}
		}
		return -1, count, NewError("Número não encontrado no array")
	}

	func createArray(width int) []int {

		array := make([]int, width);

		for i := 0; i < width; i++ {
			array[i] = i+1;
		}

		return array
	}