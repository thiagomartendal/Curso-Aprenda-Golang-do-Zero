package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	msg := <-canal // Recebe o valor do canal
	fmt.Println(msg)
}

func escrever(texto string, canal chan string) {
	for range 5 {
		canal <- texto // Atribui um valor para o canal
		time.Sleep(time.Second)
	}
}
