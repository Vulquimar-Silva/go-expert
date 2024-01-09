package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func main() {

	vulquimar := Cliente{
		Nome:  "Vulks",
		Idade: 37,
		Ativo: true,
	}

	vulquimar.Ativo = false
	vulquimar.Desativar()

}
