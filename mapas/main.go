package main

import "fmt"

func main() {
	cidades := make(map[string]string)
	cidades["Nome"] = "Salvador"
	cidades["Idade"] = "467"

	//cidades := map[string]string{"Nome": "Salvador", "Idade": "467"}

	fmt.Println(cidades)
}
