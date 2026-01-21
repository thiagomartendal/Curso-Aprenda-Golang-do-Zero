package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := range 45 {
		tarefas <- i
	}
	close(tarefas)

	for range 45 {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

// o canal tarefas só recebe dados e o canal resultados só envia
func worker(tarefas <-chan int, resultados chan<- int) {
	for num := range tarefas {
		resultados <- fibonacci(num)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
