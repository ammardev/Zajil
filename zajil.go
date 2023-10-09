package main

import (
	"github.com/ammardev/zajil/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Zajil struct {
	mode           string
	urlInput       components.UrlInput
	windowSize     tea.WindowSizeMsg
	methodSelector components.MethodSelector
	rc             components.RequestContents
	responseView   components.ResponseView
	style          lipgloss.Style
}

func NewApplicationModel() Zajil {
	zajil := Zajil{
		mode:           "normal",
		methodSelector: components.NewMethodSelector(),
		urlInput:       components.NewInput(),
		rc:             components.NewRequestContents(),
		responseView:   components.NewResponseView(),
		style:          lipgloss.NewStyle(),
	}

	return zajil
}

func (zajil Zajil) Init() tea.Cmd {
	return nil
}

func (zajil Zajil) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	switch msg.(type) {
	case tea.KeyMsg:
		cmds = append(cmds, zajil.processKeyboardInput(msg.(tea.KeyMsg)))
	case tea.WindowSizeMsg:
		zajil.windowSize = msg.(tea.WindowSizeMsg)
		zajil.urlInput.Resize(zajil.windowSize.Width - zajil.methodSelector.Width)
		zajil.rc.Resize(zajil.windowSize.Width, (zajil.windowSize.Height-zajil.urlInput.Height)/2)
		zajil.responseView.Resize(zajil.windowSize.Width, (zajil.windowSize.Height-zajil.urlInput.Height)/2)

		cmds = append(cmds, tea.ClearScreen)
	}

	cmds = append(cmds, zajil.responseView.HandleEvents(msg))

	return zajil, tea.Batch(cmds...)
}

func (zajil Zajil) View() string {
	return zajil.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				zajil.methodSelector.Render(),
				zajil.urlInput.Render(),
			),
			zajil.rc.Render(),
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
		case "enter":
			sendHttpRequest(zajil)
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
