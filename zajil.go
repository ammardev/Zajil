package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Zajil struct {
	mode       string
	urlInput   textinput.Model
	windowSize tea.WindowSizeMsg
}

func NewApplicationModel() Zajil {
	urlInput := textinput.New()
	urlInput.Prompt = ""

	return Zajil{
		mode:     "normal",
		urlInput: urlInput,
		windowSize: tea.WindowSizeMsg{
			Width:  4,
			Height: 4,
		},
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
		zajil.urlInput.Width = zajil.windowSize.Width - 7
		return zajil, nil
	}

	return zajil, nil
}

func (zajil Zajil) View() string {
	view := ""
	view += " ╭" + strings.Repeat("─", zajil.windowSize.Width-4) + "╮ \n"
	view += " │ " + zajil.urlInput.View() + " │ \n"
	view += " ╰" + strings.Repeat("─", zajil.windowSize.Width-4) + "╯ \n"

	return view
}

func (zajil *Zajil) processKeyboardInput(key tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd

	if zajil.mode == "normal" {
		switch key.String() {
		case "q", "esc":
			return tea.Quit
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
			zajil.urlInput, cmd = zajil.urlInput.Update(key)
            return cmd
		}
	}

	return nil
}
