package rotas

import (
	"api/src/middlewares"
	"net/http"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *http.ServeMux) *http.ServeMux {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...) // três pontos indica a união entre dois slices

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(
				rota.Metodo+" "+rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			)
		} else {
			r.HandleFunc(rota.Metodo+" "+rota.URI, middlewares.Logger(rota.Funcao))
		}
	}

	return r
}
