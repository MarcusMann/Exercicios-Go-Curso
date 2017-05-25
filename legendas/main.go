package main

import (
	"flag"
	"legendas/subtitle"
)

func main() {
	var diretorio = flag.String("diretorio", "", "Nome do diretório")
	var nome = flag.String("nome", "", "Exemplo: WEB-DL.HEVC.x265-RMTeam")
	var filtrar = flag.Bool("filtrar", false, "Remove as legendas não filtradas")

	flag.Parse()

	legenda := subtitle.Legendas{Nome: *nome, Diretorio: *diretorio}

	if *diretorio != "" && *nome != "" && *filtrar {
		legenda.Filtrar() // filtra as legendas e removes as legendas não filtradas.
	}
}
