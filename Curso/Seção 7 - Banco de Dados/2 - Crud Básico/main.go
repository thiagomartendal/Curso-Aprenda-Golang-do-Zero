package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("POST /usuarios", servidor.CriarUsuario)
	router.HandleFunc("GET /usuarios", servidor.BuscarUsuarios)
	router.HandleFunc("GET /usuarios/{id}", servidor.BuscarUsuario)
	router.HandleFunc("PUT /usuarios/{id}", servidor.AtualizarUsuario)
	router.HandleFunc("DELETE /usuarios/{id}", servidor.DeletarUsuario)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
