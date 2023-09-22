package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(NewApplicationModel())

	if _, err := p.Run(); err != nil {
		log.Fatalf("%+v", err)
	}

}
