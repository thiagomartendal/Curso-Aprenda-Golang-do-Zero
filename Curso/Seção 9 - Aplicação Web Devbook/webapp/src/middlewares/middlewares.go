package middlewares

import (
	"net/http"
	"webapp/src/cookies"
)

// Função desnecessária criada apenas para imprimir as rotas
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proximaFuncao(w, r)
	}
}

func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := cookies.Ler(r)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		proximaFuncao(w, r)
	}
}
