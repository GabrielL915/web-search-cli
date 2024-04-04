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
type openBrowser interface {
	OpenBrowser(ws *WebSearch, browser string) error
}

var engines = map[string]string{
	"google":     "https://www.google.com/search?q=%s",
	"duckduckgo": "https://duckduckgo.com/?q=%s",
}

var browserOpeners = map[string]openBrowser{
	"windows": &openInWindows{},
	"darwin":  &openInMacOs{},
	"linux":   &openInLinux{},
}

type openInWindows struct{}
type openInMacOs struct{}
type openInLinux struct{}

func Execute(command string, args ...string) error {
	return exec.Command(command, args...).Run()
}

func NewWebSearch(engine, query string) (*WebSearch, error) {
	searchURL, ok := engines[engine]
	if !ok {
		return nil, fmt.Errorf("engine não suportado: %s", engine)
	}
	searchURL = fmt.Sprintf(searchURL, url.QueryEscape(query))

	return &WebSearch{engine: engine, searchURL: searchURL}, nil
}

func (o *openInWindows) OpenBrowser(ws *WebSearch, browser string) error {
	if browser == "" {
		return fmt.Errorf("navegador não especificado: %s", browser)
	}
	cmd := "cmd"
	args := []string{"/c", "start"}
	if browser != "" {
		args = append(args, browser)
	}
	args = append(args, ws.searchURL)
	return Execute(cmd, args...)
}

func (o *openInMacOs) OpenBrowser(ws *WebSearch, browser string) error {
	return Execute("open", ws.searchURL)
}

func (o *openInLinux) OpenBrowser(ws *WebSearch, browser string) error {
	return Execute("xdg-open", ws.searchURL)
}

func (ws *WebSearch) OpenBrowser(browser string) error {
	opener, ok := browserOpeners[runtime.GOOS]
	if !ok {
		return fmt.Errorf("sistema operacional não suportado: %s", runtime.GOOS)
	}
	return opener.OpenBrowser(ws, browser)
}
