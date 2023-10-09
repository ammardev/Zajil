package components

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
    borderPadding = 2
    inputPromptPadding = 1
)

type UrlInput struct {
    textinput.Model
    Width int
    style lipgloss.Style
}

func NewInput(width int) UrlInput {
    input := UrlInput{
        Model: textinput.New(),
        Width: width,
        style: lipgloss.NewStyle().Width(width - borderPadding).Height(1).Border(lipgloss.RoundedBorder()),
    }
    input.Model.Prompt = ""
    input.Resize(width)

    return input
}

func (input UrlInput) GetUrl() string {
    url := input.Model.Value()

    if !strings.HasPrefix(url, "http") {
        url = "http://" + url
    }

    return url
}

func (input *UrlInput) Insert(key tea.KeyMsg) tea.Cmd {
    var cmd tea.Cmd
    input.Model, cmd = input.Model.Update(key)

    return cmd
}

func (input *UrlInput) Resize(width int) {
    input.Width = width
    input.Model.Width = width - borderPadding - inputPromptPadding
    input.style.Width(width - borderPadding)
}

func (input UrlInput) Render() string {
    return input.style.Render(
        input.Model.View(),
    )
}
