package main

import (
	"flag"
	"fmt"

	"github.com/GabrielL915/web-search-cli/internal"
)

func main() {
	query := flag.String("query", "", "A consulta de pesquisa")
	engine := flag.String("engine", "google", "O mecanismo de pesquisa a ser usado (google ou duckduckgo)")
	browser := flag.String("browser", "", "O navegador a ser usado (chrome, firefox, safari, etc.)")
	flag.Parse()

	if *query == "" {
		fmt.Println("Erro: A consulta de pesquisa é obrigatória.")
		flag.PrintDefaults()
		return
	}

	ws, err := internal.NewWebSearch(*engine, *query)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	err = ws.OpenBrowser(*browser)
	if err != nil {
		fmt.Printf("Erro ao abrir o navegador: %v\n", err)
	}
}
