package main

import (
	"fmt"
)

func main() {
	array := createArray(240000)

	target := 567

	fmt.Printf("Iniciando Algoritmo de Busca Binária, número alvo: %d\n", target)

	position, count, err := buscaBinaria(array, target)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Número encontrado na posição: %d em %d iterações\n", position, count)
}

func buscaBinaria(array []int, target int) (int, int, error) { // funcao recebe um array e o numero alvo e retorna o indice, o numero de tentativa, e um possivel erro
	low := 0               // Inicializa variavel low e atribui o valor 0 (indice minimo de um array)
	high := len(array) - 1 // Inicializa variavel high e atribui o valor do tamanho do array - 1 (indice maximo de um array)
	count := 0             // Inicializa e atribui a variavel do contador

	for low <= high { // enquanto a variavel low for menor ou igual a variavel high executa esse bloco em loop
		count++ //incrementa variavel do contador

		mid := (low + high) / 2 // inicializa e atribui a variavel de palpite (mid) calculando a soma do indice minimo com o indice maximo e dividindo por 2
		value := array[mid]     // inicializa e atribui a variavel value com o valor do array no indice do palpite

		fmt.Printf("Tentativa %d: buscando no índice: %d \n", count, mid) // Imprime no console a mensagem de tentativa

		if value == target { // compara valor com numero alvo
			return mid, count, nil // caso o valor naquele indice seja o numero alvo, retorna o indice, numero de tentative, e nao retorna um erro
		}

		if value > target { // verifica se o valor e maior que o numero alvo
			high = mid - 1 // caso o valor seja maior atualiza a variavel high para o indice anterior ao indice do palpite
		} else { // se o valor nao for maior que o numero alvo executa a linha abaixo
			low = mid + 1 // caso o valor seja menor atualiza a variavel low para o proximo indice depois do indice de palpite
		}
	}
	return -1, count, fmt.Errorf("número não encontrado no array") // Caso o numero alvo nao esteja no array, retorna o indice -1, o numero de tentativas, e retorna um erro
}

func createArray(size int) []int {

	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = i + 1
	}

	return array
}
