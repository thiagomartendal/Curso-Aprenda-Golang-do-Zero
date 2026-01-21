package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil { // Tenta recuperar o fluxo de execução do erro ocorrido
		fmt.Println("Recuperação de execução.")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é 6.") // Sinaliza um erro em tempo de execução
}

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós execução.")
}
