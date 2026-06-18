package main

import "errors"

func main() {
	println(soma(1, 2))
	println(getName())
	printHello(getName())
	println(splitAndName(getName()))
	println(multipleDeclarations())
	anonymousFunc()
	println(applyFunc(func(x int) int {
		return x * 2
	}, 2))
	err, age := isAdult(18)
	if err != nil {
		println(err)
	}
	println(age)
}

// função simples com retorno
func soma(a, b int) int {
	return a + b
}

// função que não recebe parâmetros
func getName() string {
	return "Rodrigo"
}

// função sem retorno
func printHello(name string) {
	println("Hello", name)
}

// função com múltiplos retornos
func splitAndName(name string) (string, string, int, int) {
	return name, getName(), 0, 0
}

// função com múltiplas declarações
func multipleDeclarations() (int, int, string, string) {
	var a int
	var b int
	var c string
	var d string
	return a, b, c, d
}

// função anônima
func anonymousFunc() {
	func(name string) {
		println("Hello", name)
	}("Rodrigo")
}

// função que recebe função como parâmetro
func applyFunc(f func(int) int, x int) int {
	return f(x)
}

// função que pode condicionalmente retornar um erro
func isAdult(idade int) (error, int) {
	if idade < 18 {
		return errors.New("menor de idade"), 0
	}
	return nil, idade
}

// função que pode receber um número variável de parâmetros
func variadicFunc(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

// método
type Person struct {
	Name string
	Age  int
}

// função que retorna outra função
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// método (função associada a um tipo)
func (p Person) sayHello() {
	println("Hello", p.Name)
}
