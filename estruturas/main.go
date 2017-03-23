package main

import "fmt"

type Pessoas struct {
	Nome     string
	Idade    int
	Trabalho string
}

func main() {
	marcus := Pessoas{Nome: "Marcus Mann", Idade: 25, Trabalho: "Desenvolvedor"}
	//morgana := Pessoas{Nome: "Morgana Santos", Idade: 31, Trabalho: "Advogada"}

	fmt.Println(marcus.Trabalho)
	//fmt.Println(morgana)
}
