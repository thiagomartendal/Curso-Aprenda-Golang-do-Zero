package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		// time.Sleep(time.Second)
		i++
	}

	for j := 0; j < 10; j += 2 {
		fmt.Println(j)
	}

	nomes := [3]string{"Nome A", "Nome B", "Nome C"}

	for i, v := range nomes {
		fmt.Println(i, v)
	}

	for _, v := range nomes {
		fmt.Println(v)
	}

	for i, v := range "Palavra" {
		fmt.Println(i, string(v))
	}

	usuario := map[string]string{
		"nome":      "Alguém",
		"sobrenome": "de Algo",
	}

	for k, v := range usuario {
		fmt.Println(k, v)
	}

	// for {
	// 	fmt.Println("Infinito")
	// }
}
