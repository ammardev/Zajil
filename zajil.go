package main

import (
	"github.com/ammardev/zajil/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Zajil struct {
	mode       string
	urlInput   components.UrlInput
	windowSize tea.WindowSizeMsg
    methodSelector components.MethodSelector
    responseView components.ResponseView
}

var a lipgloss.Style

func NewApplicationModel() Zajil {
    a = lipgloss.NewStyle()

	return Zajil{
		mode:     "normal",
		urlInput: components.NewInput(10),
        methodSelector: components.NewMethodSelector(),
        responseView: components.NewResponseView(10),
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
        zajil.urlInput.Resize(zajil.windowSize.Width - 13)
        zajil.responseView.Resize(zajil.windowSize.Width - 2, zajil.windowSize.Height - 6)
		return zajil, tea.ClearScreen
	}

	return zajil, nil
}

func (zajil Zajil) View() string {
    return a.Render(
        lipgloss.JoinVertical(
            lipgloss.Left,
            lipgloss.JoinHorizontal(
                lipgloss.Center,
                zajil.methodSelector.Render(),
                zajil.urlInput.Render(),
            ),
            zajil.responseView.Render(),
        ),
    )
}

func (zajil *Zajil) processKeyboardInput(key tea.KeyMsg) tea.Cmd {
	if zajil.mode == "normal" {
		switch key.String() {
		case "q", "esc":
			return tea.Quit
        case "v", "m":
            zajil.methodSelector.NextMethod()
            return nil
        case "V", "M":
            zajil.methodSelector.PreviousMethod()
            return nil
		case "i", "I":
			zajil.mode = "url"
			zajil.urlInput.Focus()
			return nil
		}
	} else if zajil.mode == "url" {
		switch key.Type {
		case tea.KeyEsc, tea.KeyCtrlC, tea.KeyEnter:
			zajil.mode = "normal"
			zajil.urlInput.Blur()
			return nil
		default:
            return zajil.urlInput.Insert(key)
		}
	}

	return nil
}
