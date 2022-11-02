package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	soma := somar(10, 15)
	fmt.Println("A soma é", soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	f("Texto da funcao 1")

	
}
