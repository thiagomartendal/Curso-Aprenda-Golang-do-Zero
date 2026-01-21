package main

import "fmt"

func main() {
	msg := func(texto string) string {
		return fmt.Sprintf("Msg: %s", texto)
	}("Olá Mundo!")

	fmt.Println(msg)
}
