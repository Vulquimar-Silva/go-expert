package main

func main() {

	// Memória -> Endereço -> Valor
	// Variavel -> ponteiro que tem um endereço na memória -> Valor

	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = 30
	println(a)

}
