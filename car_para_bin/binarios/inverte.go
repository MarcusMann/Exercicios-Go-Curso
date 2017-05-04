package binarios

// Inverte números binários
func Inverte(g ...int) []int {
	for i, j := 0, len(g)-1; i < j; i, j = i+1, j-1 {
		g[i], g[j] = g[j], g[i]
	}

	return g
}
