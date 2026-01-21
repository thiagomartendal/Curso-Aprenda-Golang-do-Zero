package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	add := n1 + n2
	sub := n1 - n2

	return add, sub
}

func main() {
	soma := somar(10, 20)

	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	txt := f("Função F")

	fmt.Println(txt)

	add, sub := calculos(10, 15)
	adicao, _ := calculos(10, 15)
	_, subtracao := calculos(10, 15)

	fmt.Println(add, sub)
	fmt.Println(adicao)
	fmt.Println(subtracao)
}
