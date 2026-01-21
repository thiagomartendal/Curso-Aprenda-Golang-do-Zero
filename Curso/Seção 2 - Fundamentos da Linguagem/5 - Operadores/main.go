package main

import "fmt"

func main() {
	// Aritméticos: +, -, /, *, %
	add := 1 + 2
	sub := 1 - 2
	div := 10 / 4
	mul := 10 * 5
	res := 10 % 2

	fmt.Println(add, sub, div, mul, res)

	var n1 int16 = 10
	var n2 int16 = 15
	soma2 := n1 + n2

	fmt.Println(soma2)

	// Atribuição: =, :=
	var v1 string = "String"
	v2 := "String"

	fmt.Println(v1, v2)

	// Relacionais: >, <, >=, <=, ==, !=
	fmt.Println(1 > 2)
	fmt.Println(2 > 1)
	fmt.Println(1 < 2)
	fmt.Println(2 < 1)
	fmt.Println(1 >= 2)
	fmt.Println(2 >= 1)
	fmt.Println(1 <= 2)
	fmt.Println(2 <= 1)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Lógicos: && (e), || (ou), ! (negação)
	v, f := true, false

	fmt.Println(v && f)
	fmt.Println(v || f)
	fmt.Println(!v, !f)

	// Unários: ++, --, +=, -=, *=, /=, %=
	numero := 10

	numero++
	fmt.Println(numero)

	numero += 15
	fmt.Println(numero)

	numero -= 5
	fmt.Println(numero)

	numero--
	fmt.Println(numero)

	numero *= 3
	fmt.Println(numero)

	numero /= 10
	fmt.Println(numero)

	numero %= 5
	fmt.Println(numero)

	// Não exite operador ternário em go

	var texto string

	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
