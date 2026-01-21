package enderecos

import "testing"

// Teste Unitário
// Execução: go test - Deve ser executado dentro da pasta onde estão os testes ou o caminho da pasta deve ser informado no comando
// Teste de cobertura do código: go test --cover - Indica a porcentagem de código da função testada que foi coberta pelo teste
// go test -v - Oferece detalhes dos testes feitos
// go test --coverprofile arquivo.txt - Gera um arquivo indicando a cobertura feita pelos testes
// go tool cover --func=arquivo.txt - Exibe no terminal as avaliações escritas no arquivo de cobertura
// go tool cover --html=arquivo.txt - Gera um arquivo html que indica quais linhas do código foram cobertas pelo teste

type CenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// t.Parallel() // Indica que o teste é feito de forma paralela

	cenariosDeTeste := []CenarioDeTeste{
		{"Rua ABC", "RUA"},
		{"Avenida 1", "AVENIDA"},
		{"Estrada A", "ESTRADA"},
		{"Rodovia R", "RODOVIA"},
		{"Praça", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s.", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}
