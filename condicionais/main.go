package main

import "fmt"

func main() {

	x := 5

	if x == 4 { // true (verdadeiro)
		fmt.Println("X é igual a 4")
	} else if x == 5 {
		fmt.Println("X é igual a 5")
	} else {
		fmt.Println("X não é igual a 4 e nem igual a 5")
	}

}
