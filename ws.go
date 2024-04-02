package main

import (
	"flag"
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
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

	var searchURL string
	switch *engine {
	case "google":
		searchURL = "https://www.google.com/search?q=" + url.QueryEscape(*query)
	case "duckduckgo":
		searchURL = "https://duckduckgo.com/?q=" + url.QueryEscape(*query)
	default:
		fmt.Printf("Mecanismo de pesquisa '%s' não suportado.\n", *engine)
		return
	}

	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c"}
		if *browser != "" {
			args = append(args, "start", *browser)
		} else {
			args = append(args, "start")
		}
	case "darwin":
		cmd = "open"
	default: // linux, *bsd, etc.
		cmd = "xdg-open"
	}

	args = append(args, searchURL)
	err := exec.Command(cmd, args...).Run()
	if err != nil {
		fmt.Printf("Erro ao abrir o navegador: %v\n", err)
	}
}
