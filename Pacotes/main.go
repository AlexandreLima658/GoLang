package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("Alexandre@gmail.com")
	fmt.Println(erro)

}
