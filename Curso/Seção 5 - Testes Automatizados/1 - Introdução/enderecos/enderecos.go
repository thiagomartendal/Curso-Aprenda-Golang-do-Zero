package enderecos

import "strings"

// Verifica se um endereço tem um tipo válido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "estrada", "avenida", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.ToTitle(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
