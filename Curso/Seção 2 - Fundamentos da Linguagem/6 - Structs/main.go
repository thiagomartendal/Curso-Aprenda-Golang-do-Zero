package main

import "fmt"

type Usuario struct {
	nome     string
	idade    uint8
	endereco Endereco
}

type Endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u Usuario

	u.nome = "Nome A"
	u.idade = 20
	fmt.Println(u)

	endereco := Endereco{"Rua B", 122}

	u2 := Usuario{"Nome B", 14, endereco}
	fmt.Println(u2)

	u3 := Usuario{idade: 15}
	fmt.Println(u3)
}
