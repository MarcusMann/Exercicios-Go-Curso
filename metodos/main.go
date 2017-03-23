package main

import (
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) info() string {
	return fmt.Sprintf("Nome: %s \nIdade: %d", p.Nome, p.Idade)
}

func main() {
	marcus := Pessoa{Nome: "Marcus Mann", Idade: 25}
	rosana := Pessoa{Nome: "Rosana Santos", Idade: 31}

	fmt.Println(marcus.info())
	fmt.Println(rosana.info())
}
