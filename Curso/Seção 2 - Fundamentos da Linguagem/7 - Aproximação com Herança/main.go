package main

import "fmt"

type Pessoa struct {
	nome  string
	idade uint8
}

type Estudante struct {
	Pessoa // Uma forma aproximada de modelar herança de orientação a objetos
	curso  string
}

func main() {
	p := Pessoa{"Nome A", 20}
	fmt.Println(p)

	e := Estudante{p, "Ciência da Computação"}
	fmt.Println(e)

	fmt.Println(e.nome)
	fmt.Println(e.idade)
	fmt.Println(e.curso)
}
