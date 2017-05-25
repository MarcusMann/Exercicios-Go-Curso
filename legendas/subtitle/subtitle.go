package subtitle

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Legendas estrutura responsável pelas legendas
type Legendas struct {
	Nome      string
	Diretorio string
}

// Return only subtitles with .srt extension
func onlySrtExtension(diretorio string) (srtList []string) {
	filename, err := ioutil.ReadDir(diretorio) // recebe o endereço do diretório
	checkError(err)                            // chama a função checkError

	for _, v := range filename {
		if ext := filepath.Ext(v.Name()); ext == ".srt" { // verifica se existe arquivos com extensão .srt
			srtList = append(srtList, v.Name()) // retorna apeans os arquivos de extensão .srt
		}
	}
	return
}

// Filtrar remove todos os arquivos não filtrados
func (l Legendas) Filtrar() {
	legendas := onlySrtExtension(filepath.Dir(l.Diretorio))

	for _, v := range legendas {
		if !strings.Contains(v, l.Nome) {
			err := os.Remove(l.Diretorio + "/" + v)
			checkError(err)
		}
	}
}

// checkError Verifica se existe algum erro ao tentar executar algum evento.
func checkError(err error) {
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}
}
