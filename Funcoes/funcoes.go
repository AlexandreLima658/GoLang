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

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

}
