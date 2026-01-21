package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!") // O comando go inicia uma go routine
	escrever("Programando em Go")
}

func escrever(texto string) {
	for { // Laço infinito
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
