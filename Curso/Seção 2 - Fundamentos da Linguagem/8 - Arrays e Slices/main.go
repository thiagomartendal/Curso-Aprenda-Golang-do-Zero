package main

import "fmt"

func main() {
	var array1 [5]string
	array1[0] = "A"
	array1[1] = "B"
	array1[2] = "C"
	array1[3] = "D"
	array1[4] = "E"
	fmt.Println(array1)

	array2 := [5]string{"A", "B", "C", "D", "E"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice1 := []int{10, 20, 30, 40, 50}
	fmt.Println(slice1)

	slice1 = append(slice1, 60)
	fmt.Println(slice1)

	slice2 := array2[1:3] // Pega os valores dos índices 1 e 2
	fmt.Println(slice2)

	array2[1] = "F" // O slice2 aponta para array2, por isso também é alterado ao alterar o array2
	fmt.Println(slice2)

	// Arrays internos

	slice3 := make([]float32, 10, 11) // Cria um array de 11 posições e um slice de 10 posições

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho
	fmt.Println(cap(slice3)) // Capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6) // Adiciona um 12° elemento e dobra sua nova capacidade

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)

	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 10)

	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
