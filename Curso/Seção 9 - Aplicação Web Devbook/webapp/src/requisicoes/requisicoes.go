package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(metodo, url, dados)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.Ler(r)
	req.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	res, err := client.Do(req) // Faz a requisição
	if err != nil {
		return nil, err
	}

	return res, nil
}
