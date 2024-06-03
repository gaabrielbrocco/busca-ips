package main

import (
	"busca-ip/app"
	"log"
	"os"
)

func main() {

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
