// TESTE DE UNIDADE

package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTypeAddress(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{

		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Estrada de ferro", "Tipo Inválido"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoObtido := TypeAddress(cenario.enderecoInserido)
		if retornoObtido != cenario.retornoEsperado {
			t.Errorf("O tipo retornado '%s' é diferente do esperado '%s'", retornoObtido, cenario.retornoEsperado)
		}
	}
}
