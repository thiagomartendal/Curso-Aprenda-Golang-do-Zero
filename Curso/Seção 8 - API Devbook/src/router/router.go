package router

import (
	"api/src/router/rotas"
	"net/http"
)

func Gerar() *http.ServeMux {
	r := http.NewServeMux()
	return rotas.Configurar(r)
}
