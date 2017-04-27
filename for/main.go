package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }
	// for {
	// 	fmt.Println("Loop Infinito")
	// }
	// numero := 1
	// for numero == 1 {
	// 	fmt.Println("loop infinito")
	// }

	mapa := map[string]string{"Cidade": "Salvador", "Idade": "467"}

	for keys, values := range mapa {
		fmt.Println("Chaves:" + keys + "Valores: " + values)
	}

}
