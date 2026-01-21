package main

import "fmt"

func soma(numeros ...int) (total int) {
	for _, n := range numeros {
		total += n
	}

	return
}

func escrever(texto string, numeros ...int) {
	for _, n := range numeros {
		fmt.Println(texto, n)
	}
}

func main() {
	s := soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)

	fmt.Println(s)

	escrever("Olá Mundo", 1, 2, 3, 4)
}
