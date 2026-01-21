package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, codigoStatus int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(codigoStatus)

	if dados != nil {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatal(err)
		}
	}
}

func Erro(w http.ResponseWriter, codigoStatus int, erro error) {
	JSON(w, codigoStatus, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
