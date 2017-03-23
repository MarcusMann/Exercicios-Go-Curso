package main

import (
	"fmt"
)

func main() {
	//var lista []string
	lista := make([]string, 3)
	//lista := []string{}

	lista = append(lista, "Salvador", "467", "Bahia")

		fmt.Println(cap(lista))
}
