package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type usuario struct {
	Nome  string `json:"nome"`  // Especifica os nomes dos campos para a estrutura json
	Idade uint   `json:"idade"` // `json:"-"` - Indica que o campo deve ser ignorado
}

func main() {
	u1JSON := `{"nome": "Usuário 1", "idade": 20}`

	var u1 usuario

	json.Unmarshal([]byte(u1JSON), &u1)

	fmt.Println(u1)

	u2JSON := `{"nome": "Usuário 2", "idade": "15"}`

	u2 := make(map[string]string)

	if err := json.Unmarshal([]byte(u2JSON), &u2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(u2)
}
