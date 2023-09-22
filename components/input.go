package components

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input struct {
    textinput.Model
    Width int
}

func NewInput(width int) Input {
    input := Input{
        Model: textinput.New(),
    }
    input.Model.Prompt = ""
    input.Resize(width)

    return input
}

func (input *Input) Insert(key tea.KeyMsg) tea.Cmd {
    var cmd tea.Cmd
    input.Model, cmd = input.Model.Update(key)

    return cmd
}

func (input *Input) Resize(width int) {
    input.Width = width
    input.Model.Width = width - 7
}

func (input Input) Render() string {
    view := ""
	view += " ╭" + strings.Repeat("─", input.Width-4) + "╮ \n"
	view += " │ " + input.Model.View() + " │ \n"
	view += " ╰" + strings.Repeat("─", input.Width-4) + "╯ \n"

    return view
}
