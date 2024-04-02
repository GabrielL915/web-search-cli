package internal

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
)

type WebSearch struct {
	engine    string
	searchURL string
}

func NewWebSearch(engine, query string) (*WebSearch, error) {
	var searchURL string

	switch engine {
	case "google":
		searchURL = "https://www.google.com/search?q=" + url.QueryEscape(query)
	case "duckduckgo":
		searchURL = "https://duckduckgo.com/?q=" + url.QueryEscape(query)
	default:
		return nil, fmt.Errorf("Mecanismo de pesquisa '%s' não suportado", engine)
	}

	return &WebSearch{engine: engine, searchURL: searchURL}, nil
}

func (ws *WebSearch) OpenBrowser(browser string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c"}
		if browser != "" {
			args = append(args, "start", browser)
		} else {
			args = append(args, "start")
		}
	case "darwin":
		cmd = "open"
	default: // linux, *bsd, etc.
		cmd = "xdg-open"
	}

	args = append(args, ws.searchURL)
	return exec.Command(cmd, args...).Run()
}
