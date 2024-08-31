package main

import (
	"Algoritmos/Busca-em-largura/desenvolvedor"
	"fmt"
)

func main() {

	igor := createDevsNetwork()

	fmt.Printf("Buscando um Dev Golang\n")

	resultado, err := buscaEmLargura(igor)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("O Dev %s programa em Golang \n", resultado.Nome)

}

func buscaEmLargura(dev *desenvolvedor.Desenvolvedor) (*desenvolvedor.Desenvolvedor, error) {

	queue := dev.Conexoes
	verificados := make(map[string]bool)

	for len(queue) > 0 {

		fmt.Printf("Verificando se %s programa em Golang \n", queue[0].Nome)

		if verificados[queue[0].Nome] {
			fmt.Printf("%s já foi verificado \n", queue[0].Nome)
			queue = queue[1:]
			continue
		}

		if queue[0].DesenvolveEmGolang() {
			return &queue[0], nil
		}

		verificados[queue[0].Nome] = true

		queue = append(queue, queue[0].Conexoes...)
		queue = queue[1:]
	}

	return nil, fmt.Errorf("nenhum Dev Golang encontrado")
}

func createDevsNetwork() *desenvolvedor.Desenvolvedor {
	diogo := desenvolvedor.New("Diogo", make([]desenvolvedor.Desenvolvedor, 0), "Golang")
	fernando := desenvolvedor.New("Fernando", make([]desenvolvedor.Desenvolvedor, 0), "PHP")
	emanuel := desenvolvedor.New("Emanuel", []desenvolvedor.Desenvolvedor{fernando}, "JavaScript")
	andre := desenvolvedor.New("Andre", []desenvolvedor.Desenvolvedor{fernando, emanuel}, "JavaScript")
	ricardo := desenvolvedor.New("Ricardo", []desenvolvedor.Desenvolvedor{fernando, andre, emanuel, diogo}, "C#")
	ana := desenvolvedor.New("Ana", []desenvolvedor.Desenvolvedor{fernando, ricardo, andre, emanuel}, "Java")
	mauricio := desenvolvedor.New("Mauricio", []desenvolvedor.Desenvolvedor{fernando, ana, ricardo, andre, emanuel}, "C#")

	diego := desenvolvedor.New("Diego", []desenvolvedor.Desenvolvedor{emanuel, fernando}, "JavaScript")
	pablo := desenvolvedor.New("Pablo", []desenvolvedor.Desenvolvedor{andre, ricardo, fernando}, "Java")
	matheus := desenvolvedor.New("Matheus", []desenvolvedor.Desenvolvedor{pablo, ricardo, ana, mauricio}, "TypeScript")

	igor := desenvolvedor.New("Igor Bonfim", []desenvolvedor.Desenvolvedor{pablo, matheus, diego}, "C#")

	return &igor
}
