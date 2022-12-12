package enderecos

import (
	"strings"
)

//TypeAddress é um tipo de dado que define um endereço
func TypeAddress(address string) string {
	tiposValidos := []string{"rua", "avenida", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(address)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
