package main

import "fmt"

type Usuario struct {
	Nome  string
	Idade int

	Endereco Endereco
}

type Endereco struct {
	logradouro string
	numero     uint16
}

func main() {
	usuario := Usuario{
		Nome:  "Rodrigo",
		Idade: 35,
		Endereco: Endereco{
			logradouro: "Rua A",
			numero:     123,
		},
	}

	fmt.Println(usuario)
}
