package main

import "fmt"

func main() {
	//ele não retornou true / false		// contador = contador + 1 ou contador += 1
	for contador := 0; contador <= 10; contador++ {
		fmt.Println(contador)
	}

	//while True{}

	var contador int = 0
	for {
		contador++

		fmt.Println(contador)

		if contador == 20 {
			break
		}
	}

	mapas := map[string]string{
		"Cidade": "Salvador",
		"Idade":  "467",
		"Estado": "Bahia",
	}

	for keys, _ := range mapas {
		fmt.Println(keys)
	}

}
