package main

import (
	"log"
	"os"

	_ "github.com/besser/goinaction/chapter2/matcher"
	"github.com/besser/goinaction/chapter2/search"
)

func init() {
	// Muda o dispositivo de logging para stdout
	log.SetOutput(os.Stdout)i

	// Faz o log incluir o nome do arquivo e o número da linha que o erro ocorreu
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// Pesquisa o termo especificado
	search.Run("president")
}
