package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Página Inicial"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
