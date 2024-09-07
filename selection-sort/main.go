package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := createArray(10)

	fmt.Printf("Array não ordenado: %d \n", array)

	arrayOrdenado := selectionSort(array);

	fmt.Printf("Array ordenado: %d \n", arrayOrdenado)
}

func selectionSort(array []int) []int {
	arrayOrdenado := make([]int, len(array))

	for i := range array {
		menor := menorValor(array)
		arrayOrdenado[i] = array[menor]
		array = remove(array, menor)
	}

	return arrayOrdenado
}

func menorValor(array []int) int {
	indiceMenor := 0 //Inicia a variavel IndiceMenor como 0

	for i, valor := range array { // loop for para percorrer todo o array
		if valor < array[indiceMenor] { //compara se o valor atual do loop é menor que o nosso menor valor
			indiceMenor = i // Caso seja, altera a variavel indiceMenor para o valor do indice atual
		}
	}

	return indiceMenor	// retorna o indice do menor valor
}

func remove(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func createArray(size int) []int {

	array := make([]int, size)

	for i := range array {
		array[i] = rand.Intn(100)
	}

	return array
}
