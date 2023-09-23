package main

import (
	"github.com/ammardev/zajil/components"
	tea "github.com/charmbracelet/bubbletea"
)

type Zajil struct {
	mode       string
	requestLineInput   components.RequestLineInput
	windowSize tea.WindowSizeMsg
}

func NewApplicationModel() Zajil {
	return Zajil{
		mode:     "normal",
		requestLineInput: components.NewInput(10),
	}

}

func (zajil Zajil) Init() tea.Cmd {
	return tea.ClearScreen
}

func (zajil Zajil) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return zajil, zajil.processKeyboardInput(msg.(tea.KeyMsg))
	case tea.WindowSizeMsg:
		zajil.windowSize = msg.(tea.WindowSizeMsg)
        zajil.requestLineInput.Resize(zajil.windowSize.Width)
		return zajil, nil
	}

	return zajil, nil
}

func (zajil Zajil) View() string {
    return zajil.requestLineInput.Render()
}

func (zajil *Zajil) processKeyboardInput(key tea.KeyMsg) tea.Cmd {
	if zajil.mode == "normal" {
		switch key.String() {
		case "q", "esc":
			return tea.Quit
        case "v", "V", "M", "m":
            zajil.requestLineInput.SwitchMethod()
            return nil
		case "i", "I":
			zajil.mode = "url"
			zajil.requestLineInput.Focus()
			return nil
		}
	} else if zajil.mode == "url" {
		switch key.Type {
		case tea.KeyEsc, tea.KeyCtrlC, tea.KeyEnter:
			zajil.mode = "normal"
			zajil.requestLineInput.Blur()
			return nil
		default:
            return zajil.requestLineInput.Insert(key)
		}
	}

	return nil
}
