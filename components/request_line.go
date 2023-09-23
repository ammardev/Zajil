package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
    renderStyle lipgloss.Style
)

type RequestLineInput struct {
    textinput.Model
    Width int
}

func NewInput(width int) RequestLineInput {
    renderStyle = lipgloss.NewStyle().Width(width).Border(lipgloss.RoundedBorder())

    input := RequestLineInput{
        Model: textinput.New(),
    }
    input.Model.Prompt = ""
    input.Resize(width)

    return input
}

func (input *RequestLineInput) Insert(key tea.KeyMsg) tea.Cmd {
    var cmd tea.Cmd
    input.Model, cmd = input.Model.Update(key)

    return cmd
}

func (input *RequestLineInput) Resize(width int) {
    input.Width = width
    input.Model.Width = width
    renderStyle.Width(width)
}

func (input RequestLineInput) Render() string {
    return renderStyle.Render(
        input.Model.View(),
    )
}
