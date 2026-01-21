package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Alguém",
		"sobrenome": "de Algo",
	}

	usuario["idade"] = "Alguma"

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	// Maps aninhados
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Alguém",
			"ultimo":   "de Algo",
		},
		"curso": {
			"nome": "Algum",
			"onde": "em um lugar",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"], usuario2["nome"]["ultimo"])
	fmt.Println(usuario2["curso"])

	delete(usuario2, "curso")

	fmt.Println(usuario2)
}
