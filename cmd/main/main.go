package main

import (
	"os"

	"github.com/Yahar4/internal/model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	filePath := "quotes.json"

	m := model.InitialModel(filePath)
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
