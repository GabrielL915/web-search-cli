package test

import (
	"testing"

	wsc "github.com/GabrielL915/web-search-cli/internal"
)

func TestNewWebSearch(t *testing.T) {
	cases := []struct {
		engine, query, wantErr string
	}{
		{"google", "golang", ""},
		{"duckduckgo", "rust", ""},
		{"engineNãoSuportado", "java", "engine não suportado: engineNãoSuportado"},
	}

	for _, c := range cases {
		ws, err := wsc.NewWebSearch(c.engine, c.query)
		if c.wantErr != "" {
			if err == nil || err.Error() != c.wantErr {
				t.Errorf("NewWebSearch(%q, %q) error = %v, wantErr %q", c.engine, c.query, err, c.wantErr)
			}
			continue
		}
		if err != nil {
			t.Errorf("NewWebSearch(%q, %q) unexpected error: %v", c.engine, c.query, err)
			continue
		}
		if ws == nil {
			t.Errorf("NewWebSearch(%q, %q) = nil, want not nil", c.engine, c.query)
		}
	}
}
