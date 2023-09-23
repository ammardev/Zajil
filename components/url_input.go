package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
    renderStyle lipgloss.Style
)

type UrlInput struct {
    textinput.Model
    Width int
}

func NewInput(width int) UrlInput {
    renderStyle = lipgloss.NewStyle().Width(width).Border(lipgloss.RoundedBorder())

    input := UrlInput{
        Model: textinput.New(),
    }
    input.Model.Prompt = ""
    input.Resize(width)

    return input
}

func (input *UrlInput) Insert(key tea.KeyMsg) tea.Cmd {
    var cmd tea.Cmd
    input.Model, cmd = input.Model.Update(key)

    return cmd
}

func (input *UrlInput) Resize(width int) {
    input.Width = width
    input.Model.Width = width
    renderStyle.Width(width)
}

func (input UrlInput) Render() string {
    return renderStyle.Render(
        input.Model.View(),
    )
}
