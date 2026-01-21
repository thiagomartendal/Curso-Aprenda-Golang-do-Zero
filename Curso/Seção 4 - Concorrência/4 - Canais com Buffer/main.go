package main

import "fmt"

func main() {
	canal := make(chan string, 2) // Canal com buffer de tamanho 2 (Só bloqueia quando a capacidade máxima é atingida)

	canal <- "Olá Mundo!"
	canal <- "Programando em Go"

	msg1 := <-canal
	msg2 := <-canal

	fmt.Println(msg1)
	fmt.Println(msg2)
}
