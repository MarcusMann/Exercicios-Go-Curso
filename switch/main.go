package main

import "fmt"

func main() {

	var numero int = 1

	switch numero {
	case 2:
		fmt.Println("Numero é 2")
	case 3:
		fmt.Println("Numero é 3")
	default:
		fmt.Println("Número não é 2 e nem 3")
	}

}
