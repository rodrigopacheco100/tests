package main

import "fmt"

func main() {
	var texto string = "Texto"
	var (
		nome   string  = "Rodrigo Pacheco"
		idade  int     = 30
		altura float64 = 1.77
		ativo  bool    = true
	)
	texto2 := "Texto2"
	texto3, texto4 := "Texto3", "Texto4"
	const pi float64 = 3.14159

	fmt.Println(texto, texto2, texto3, texto4, nome, idade, altura, ativo, pi)

	// Inverter valores
	texto2, texto3 = texto3, texto2
	fmt.Println(texto2, texto3)
}
