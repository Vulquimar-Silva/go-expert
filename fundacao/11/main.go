package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	vulquimar := Cliente{
		Nome:  "Vulks",
		Idade: 37,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", vulquimar.Nome, vulquimar.Idade, vulquimar.Ativo)

}
