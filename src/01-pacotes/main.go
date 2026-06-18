package main

import (
	"fmt"
	"pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	auxiliar.Escrever()
	if err := checkmail.ValidateFormat("rodrigopacheco100@gmail.com"); err != nil {
		fmt.Println(err)
	}
}
