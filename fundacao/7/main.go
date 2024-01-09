package main

import "fmt"

func main() {

	salarios := map[string]int{"Vulquimar": 9000, "Nayara": 9500, "Maria": 8000}
	fmt.Println(salarios["Nayara"])
	delete(salarios, "Vulquimar")
	salarios["Vulks"] = 18000
	fmt.Println(salarios["Vulks"])

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Vulquimar"] = 9000

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
