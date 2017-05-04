package binarios

import (
	"math"
)

// ParaBytes Transforma números binários em números bytes
func ParaBytes(g ...int) (soma byte) {
	var resultado float64

	for i := 0; i < len(g); i++ {
		resultado = float64(g[i]) * (math.Pow(2.0, float64(i)))
		soma = soma + byte(resultado)
	}

	return
}
