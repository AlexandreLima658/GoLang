package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// função com retorno múltiplo
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := somar(n1, n2)
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 15)
	fmt.Println("A soma é", soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função")
	fmt.Println(resultado)

	// retornar multiplos valores
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	//retornar multiplos valores ignorando o resultadoSubtracao
	resultadoSoma, _ = calculosMatematicos(12, 5)
	fmt.Println(resultadoSoma)

	// retorna multiplos valores ignorando o resultadoSoma
	 _ ,resultadoSubtracao = calculosMatematicos(12, 5)
	fmt.Println(resultadoSubtracao)
}
