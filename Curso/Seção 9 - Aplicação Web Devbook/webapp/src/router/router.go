package router

import (
	"net/http"
	"webapp/src/router/rotas"
)

func Gerar() *http.ServeMux {
	r := http.NewServeMux()
	return rotas.Configurar(r)
}
