package rotas

import (
	"net/http"
	"webapp/src/middlewares"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(router *http.ServeMux) *http.ServeMux {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaPaginaPrincipal)
	rotas = append(rotas, rotasPublicacoes...)
	rotas = append(rotas, rotaLogout)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			router.HandleFunc(rota.Metodo+" "+rota.URI, middlewares.Autenticar(rota.Funcao))
		} else {
			router.HandleFunc(rota.Metodo+" "+rota.URI, rota.Funcao)
		}
	}

	// Serve o diretório de assets
	fileServer := http.FileServer(http.Dir("./assets/"))                    // Inclui o diretório de assets do frontend
	router.Handle("GET /assets/", http.StripPrefix("/assets/", fileServer)) // Disponibiliza os arquivos para as requisições do navegador

	return router
}
