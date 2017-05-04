package main

import (
	app "exercicios/06/binarios"
	"fmt"
)

func main() {
	x := app.ParaBinarios("d")
	// resultado := app.ParaBytes(x...)

	// fmt.Println(resultado)
	// fmt.Println(string(resultado))
	f := app.Inverte(x...)
	fmt.Println(f)
}
