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

type browser func(ws *WebSearch, browser string) error

var engines = map[string]string{
	"google":     "https://www.google.com/search?q=%s",
	"duckduckgo": "https://duckduckgo.com/?q=%s",
}
var so = map[string]browser{
	"windows": openBrowserWindows,
	"darwin":  openBrowserMacOS,
	"linux":   openBrowserLinux,
}

func NewWebSearch(engine, query string) (*WebSearch, error) {
	searchURL, ok := engines[engine]
	if !ok {
		return nil, fmt.Errorf("engine não suportado: %s", engine)
	}
	searchURL = fmt.Sprintf(searchURL, url.QueryEscape(query))

	return &WebSearch{engine: engine, searchURL: searchURL}, nil
}

func openBrowserWindows(ws *WebSearch, browser string) error {
	cmd := "cmd"
	args := []string{"/c"}
	if browser != "" {
		args = append(args, "start", browser)
	} else {
		args = append(args, "start")
	}
	args = append(args, ws.searchURL)
	return exec.Command(cmd, args...).Run()
}

func openBrowserMacOS(ws *WebSearch, browser string) error {
	cmd := "open"
	args := []string{ws.searchURL}
	return exec.Command(cmd, args...).Run()
}

func openBrowserLinux(ws *WebSearch, browser string) error {
	cmd := "xdg-open"
	args := []string{ws.searchURL}
	return exec.Command(cmd, args...).Run()

}

func (ws *WebSearch) OpenBrowser(browser string) error {
	open, ok := so[runtime.GOOS]
	if !ok {
		return fmt.Errorf("sistema operacional não suportado: %s", runtime.GOOS)
	}
	return open(ws, browser)
}
