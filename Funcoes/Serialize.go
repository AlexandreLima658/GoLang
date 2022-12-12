package main

import (
	"encoding/json"
	"fmt"
)

type Banco struct {
	Code string
	Name string
}

func Serialize(banco Banco) string {
	json, _ := json.Marshal(banco)
	return string(json)
}

func main() {
	banco := Banco{"001", "Banco do Brasil S.A."}
	fmt.Println(Serialize(banco))
}
