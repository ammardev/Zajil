package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(NewApplicationModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatalf("%+v", err)
	}

}
