package main

import "fmt"

type Usuario struct {
	nome  string
	idade int
}

func (u Usuario) salvar() {
	fmt.Printf("Salvando o usuário %s\n", u.nome)
}

func (u Usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *Usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := Usuario{"Usuário 1", 20}

	fmt.Println(usuario1)

	usuario1.salvar()

	fmt.Println(usuario1.maiorDeIdade())

	usuario1.fazerAniversario()

	fmt.Println(usuario1)
}
