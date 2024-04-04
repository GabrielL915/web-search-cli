package test

import (
	"testing"

	wsc "github.com/GabrielL915/web-search-cli/internal"

	"github.com/stretchr/testify/mock"
)

type MockExecute struct {
	mock.Mock
}

func (m *MockExecute) Execute(command string, args ...string) error {
	arguments := m.Called(command, args)
	return arguments.Error(0)
}

func TestOpenBrowserInWindows(t *testing.T) {
	executor := new(MockExecute)
	executor.On("Execute", mock.Anything, mock.Anything).Return(nil)

	open := &wsc.OpenInWindows{}
	ws, _ := wsc.NewWebSearch("google", "golang")
	err := open.OpenBrowser(executor, ws, "chrome")
	if err != nil {
		t.Errorf("OpenBrowser() error = %v, want nil", err)
	}
	executor.AssertExpectations(t)
}

func TestOpenBrowserInMacOs(t *testing.T) {
	executor := new(MockExecute)
	executor.On("Execute", mock.Anything, mock.Anything).Return(nil)

	open := &wsc.OpenInMacOs{}
	ws, _ := wsc.NewWebSearch("google", "golang")
	err := open.OpenBrowser(executor, ws, "firefox")
	if err != nil {
		t.Errorf("OpenBrowser() error = %v, want nil", err)
	}
	executor.AssertExpectations(t)
}

func TestOpenBrowserInLinux(t *testing.T) {
	executor := new(MockExecute)
	executor.On("Execute", mock.Anything, mock.Anything).Return(nil)

	open := &wsc.OpenInLinux{}
	ws, _ := wsc.NewWebSearch("google", "golang")
	err := open.OpenBrowser(executor, ws, "chrome")
	if err != nil {
		t.Errorf("OpenBrowser() error = %v, want nil", err)
	}
	executor.AssertExpectations(t)
}
