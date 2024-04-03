package cmd

import (
	"fmt"
	"os"

	"github.com/GabrielL915/web-search-cli/internal"
	"github.com/spf13/cobra"
)

var query, engine, browser string

var rootCmd = &cobra.Command{
	Use:   "ws-cli",
	Short: "Uma aplicação de linha de comando para pesquisar na web",
	Long: `web-search-cli é uma aplicação de linha de comando que permite pesquisar na web, 
	usando flags para especificar a consulta de pesquisa, o mecanismo de pesquisa e o navegador.`,
	Run: func(cmd *cobra.Command, args []string) {
		if query == "" {
			fmt.Println("Erro: A consulta de pesquisa é obrigatória.")
			cmd.Help()
			return
		}

		ws, err := internal.NewWebSearch(engine, query)
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
			return
		}

		err = ws.OpenBrowser(browser)
		if err != nil {
			fmt.Printf("Erro ao abrir o navegador: %v\n", err)
		}
	},
}

func init() {
	cobra.OnInitialize()
	rootCmd.Flags().StringVarP(&query, "query", "q", "", "A consulta de pesquisa (obrigatório)")
	rootCmd.Flags().StringVarP(&engine, "engine", "e", "google", "O mecanismo de pesquisa a ser usado (google ou duckduckgo)")
	rootCmd.Flags().StringVarP(&browser, "browser", "b", "", "O navegador a ser usado (chrome, firefox, safari, etc.)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
