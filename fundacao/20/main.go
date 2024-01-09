package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m := map[string]int{"Vulquimar": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Vulquimar": 290.50, "João": 389, "Maria": 525.90}
	m3 := map[string]MyNumber{"Vulquimar": 2900, "João": 3890, "Maria": 5250}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))

}
