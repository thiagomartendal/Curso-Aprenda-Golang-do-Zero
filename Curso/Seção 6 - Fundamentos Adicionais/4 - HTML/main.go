package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Idade uint
}

func main() {
	templates = template.Must(template.ParseGlob("*.html")) // Lê os arquivos HTML

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"Usuário 1", 20}

		templates.ExecuteTemplate(w, "index.html", u) // Escreve o arquivo html na rota
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
