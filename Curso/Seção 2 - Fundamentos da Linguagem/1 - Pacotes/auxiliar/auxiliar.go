package auxiliar

import "fmt"

// Funções com nome iniciando em letra maiúscula definem um nome exportado
// e podem ser importadas em outros arquivos
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
