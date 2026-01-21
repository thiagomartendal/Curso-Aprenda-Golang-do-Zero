package main

import "fmt"

func main() {
	var v1 int = 10
	var v2 int = v1
	fmt.Println(v1, v2)

	v1++
	fmt.Println(v1, v2)

	var v3 int = 100
	var p *int = &v3

	fmt.Println(v3, *p, p) // *p => Desfaz a referência

	v3 = 150
	fmt.Println(v3, *p, p)
}
