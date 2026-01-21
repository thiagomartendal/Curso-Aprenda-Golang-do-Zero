package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Os campos de uma struct que será convertida em json devem ser exportados, por isso, devem iniciar em letra maiúscula
type usuario struct {
	Nome  string `json:"nome"` // Especifica os nomes dos campos para a estrutura json
	Idade uint   `json:"idade"`
}

func main() {
	u1 := usuario{"Usuário 1", 20}
	fmt.Println(u1)

	u1JSON, err := json.Marshal(u1) // Converte um objeto para JSON
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes.NewBuffer(u1JSON))

	u2 := map[string]string{
		"nome":  "Usuário 2",
		"idade": "15",
	}
	fmt.Println(u2)

	u2JSON, err := json.Marshal(u2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes.NewBuffer(u2JSON))
}
