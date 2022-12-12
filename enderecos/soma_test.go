package enderecos

import "testing"

func TestSoma(t *testing.T) {
	testes := []struct {
		valores []int
		res     int
	}{
		{valores: []int{1, 2}, res: 3},
		{valores: []int{1, 2, 3}, res: 6},
		{valores: []int{1, 2, 3, 4}, res: 10},
		{valores: []int{1, 2, 3, 4, 5}, res: 15},
	}

	for _, teste := range testes {
		total := Soma(teste.valores...)
		if total != teste.res {
			t.Errorf("O total %d é diferente do esperado %d", total, teste.res)
		}
	}

}
