package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de Partida")

	aplicacao := app.Gerar()

	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

/*
	Execução:
	go run main.go ip => padrão
	go run main.go ip --host www.google.com.br
	go run main.go servidores --host google.com.br
	./linha-de-comando ip => padrão
	./linha-de-comando ip --host www.google.com.br
	./linha-de-comando servidores --host google.com.br
*/
