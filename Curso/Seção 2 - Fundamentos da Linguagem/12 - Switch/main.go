package main

import "fmt"

func diaDaSemana(numero int) (dia string) {
	switch numero {
	case 1:
		dia = "Domingo"
	case 2:
		dia = "Segunda-Feira"
	case 3:
		dia = "Terça-Feira"
	case 4:
		dia = "Quarta-Feira"
	case 5:
		dia = "Quinta-Feira"
	case 6:
		dia = "Sexta-Feira"
	case 7:
		dia = "Sábado"
	default:
		dia = "Indefinido"
	}
	return
}

func diaDaSemana2(numero int) (dia string) {
	switch {
	case numero == 1:
		dia = "Domingo"
		// fallthrough // Indica que deve seguir para o próximo caso
	case numero == 2:
		dia = "Segunda-Feira"
	case numero == 3:
		dia = "Terça-Feira"
	case numero == 4:
		dia = "Quarta-Feira"
	case numero == 5:
		dia = "Quinta-Feira"
	case numero == 6:
		dia = "Sexta-Feira"
	case numero == 7:
		dia = "Sábado"
	default:
		dia = "Indefinido"
	}
	return
}

func main() {
	dia := diaDaSemana(6)

	fmt.Println(dia)

	dia = diaDaSemana(8)

	fmt.Println(dia)

	dia = diaDaSemana2(1)

	fmt.Println(dia)
}
