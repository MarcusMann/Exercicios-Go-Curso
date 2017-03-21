package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Marcus Mann" // forma literal de se criar uma váriavel
	x := 10 // isso aqui é um número inteiro
	
	// verifica o tipo da variável
	fmt.Println("É do tipo: ", reflect.TypeOf(nome))
	fmt.Println("É do tipo: ", reflect.TypeOf(x))
}
