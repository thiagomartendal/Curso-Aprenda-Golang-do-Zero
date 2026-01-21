package main

import (
	"Modulo/auxiliar"
	"fmt"

	"github.com/badoux/checkmail"
)

/*
	Criando um módulo: go mod init nome_do_modulo
	Compilação de projeto em módulo: go build
	Compilação para a raiz: go install - Compila e exporta o executável para a pasta raiz onde o Go foi instalado
	Importar pacote externo: go get local_do_pacote
	Remoção de dependências não usadas no arquivo mod: go mod tidy
*/

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	// Pacotes Externos
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
