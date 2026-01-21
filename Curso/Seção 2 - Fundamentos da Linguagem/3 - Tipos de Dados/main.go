package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos inteiros: int, int8, int16, int32, int 64
	// Tipos inteiros não sinalizados: uint, uint8, uint16, uint32, uint 64
	// Tipos reais: float32, float64
	// Tipo string: string
	// Tipo booleano: bool
	// Tipo erro: error

	var numero1 int16 = 100
	var numero2 uint32 = 10000
	var numero3 rune = 123456 // Alias - int32 = rune
	var numero4 byte = 123    // Alias - uint8 = byte

	var real1 float32 = 123.45
	var real2 float64 = 123000000000000.45
	real3 := 12345.67

	var string1 string = "Texto"
	string2 := "Texto2"

	char := 'B' // Atribui o valor ascii de B

	var booleano1 bool // Valor zero do tipo booleano é false
	var booleano2 bool = true

	var erro1 error // Valor zero do tipo erro é nil (nulo)
	var erro2 error = errors.New("Erro interno")

	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(real1)
	fmt.Println(real2)
	fmt.Println(real3)
	fmt.Println(string1)
	fmt.Println(string2)
	fmt.Println(char)
	fmt.Println(booleano1)
	fmt.Println(booleano2)
	fmt.Println(erro1)
	fmt.Println(erro2)
}
