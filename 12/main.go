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

func main() {

	vulquimar := Cliente{
		Nome:  "Vulks",
		Idade: 37,
		Ativo: true,
	}

	vulquimar.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", vulquimar.Nome, vulquimar.Idade, vulquimar.Ativo)

}
