package desenvolvedor

type Desenvolvedor struct {
	Nome      string
	Conexoes  []Desenvolvedor
	Linguagem string
}

func New(nome string, conexoes []Desenvolvedor, linguagem string) Desenvolvedor {
	dev := Desenvolvedor{nome, conexoes, linguagem}
	return dev
}

func (desenvolvedor *Desenvolvedor) DesenvolveEmGolang() bool {
	return desenvolvedor.Linguagem == "Golang"
}
