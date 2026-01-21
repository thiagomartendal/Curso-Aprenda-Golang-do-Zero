package main

import "fmt"

var n int

func init() { // A Função init é executa primeiro, e pode ter uma função init por arquivo, e não por pacote
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}
