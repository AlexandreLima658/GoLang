package main

import "fmt"

func main() {
	var varival1 string = "Variavel 1"
	var varival2 string = "Variavel 2"

	fmt.Println(varival1)
	fmt.Println(varival2)

	variavel3 := "Variavel 3"
	variavel4 := "Variavel 4"

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	var (
		variavel5 string = "Variavel 5"
		variavel6 string = "Variavel 6"
	)

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	variavel7, variavel8 := "Variavel 7", "Variavel 8"

	fmt.Println(variavel7)
	fmt.Println(variavel8)

	variavel7, variavel8 = variavel8, variavel7

	fmt.Println(variavel7, variavel8)

}
