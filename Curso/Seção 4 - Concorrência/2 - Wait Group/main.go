package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup // Abre um grupo de espera

	waitGroup.Add(2) // Ajusta a contagem do grupo de espera para esperar por 2 go routines

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // Reduz o contador de go routines do grupo de espera
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done() // Reduz o contador de go routines do grupo de espera
	}()

	waitGroup.Wait() // Espera a contagem de go routines estabelecida chegar em zera
}

func escrever(texto string) {
	for range 5 {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
