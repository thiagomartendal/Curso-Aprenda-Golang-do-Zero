package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada")

	media := (n1 + n2) / 2

	return (media >= 6)
}

func main() {
	defer funcao1()
	funcao2()

	fmt.Println(alunoAprovado(7, 8))
}
