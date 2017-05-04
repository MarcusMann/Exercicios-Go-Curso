package binarios

// ParaBinarios transforma strings em binarios
func ParaBinarios(caracter string) (lista []int) {
	b := []byte(caracter)

	for b[0] > 0 {
		lista = append(lista, int(b[0])%2)
		b[0] /= 2 // b[0] = b[0] / 2
	}

	return
}
