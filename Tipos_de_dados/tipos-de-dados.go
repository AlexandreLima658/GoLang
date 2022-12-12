package main

import "fmt"

func main() {
	var numero1 int8 = 10
	var numero2 int16 = 1000
	var numero3 int32 = 100000
	var numero4 int64 = 1000000000000000000
	var numero5 uint8 = 255
	var numero6 uint16 = 65535
	var numero7 uint32 = 4294967295
	var numero8 uint64 = 18446744073709551615

	fmt.Println(numero1, numero2, numero3, numero4, numero5, numero6, numero7, numero8)

	//alias
	//INT32 = RUNE
	var numero9 rune = 123456
	fmt.Println(numero9)

	//UINT8 = BYTE
	var numero10 byte = 123
	fmt.Println(numero10)

	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)
	var numeroReal1 float64 = 123456789.45
	fmt.Println(numeroReal1)

	var str string = "Texto"
	fmt.Println(str)
	srt2 := "Texto"
	fmt.Println(srt2)

}
