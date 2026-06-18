package main

import "fmt"

type Usuario struct {
	Nome  string
	Idade int
}

type Estudante struct {
	Usuario
	Curso     string
	Matricula string
}

func main() {
	estudante := Estudante{
		Usuario: Usuario{
			Nome:  "Rodrigo",
			Idade: 35,
		},
		Curso:     "Ciência da Computação",
		Matricula: "123456",
	}

	fmt.Println(estudante)
}
