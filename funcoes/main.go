package main

import "fmt"

func soma(numero1, numero2 int) int {
	return numero1 + numero2
}

func main() {
	s := soma(1, 2)

	fmt.Println(s)
}
