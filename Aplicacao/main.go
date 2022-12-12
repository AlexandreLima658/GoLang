package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	application := app.Generation()
	erro := application.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}

}
