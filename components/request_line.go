package components

import (
	"strings"

	"github.com/ammardev/zajil/style"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var Methods []string

type RequestLineInput struct {
    textinput.Model
    Width int
    SelectedMethod int
}

func NewInput(width int) RequestLineInput {
    Methods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

    input := RequestLineInput{
        Model: textinput.New(),
    }
    input.Model.Prompt = ""
    input.Resize(width)
    input.SelectedMethod = 0

    return input
}

func (input *RequestLineInput) Insert(key tea.KeyMsg) tea.Cmd {
    var cmd tea.Cmd
    input.Model, cmd = input.Model.Update(key)

    return cmd
}

func (input *RequestLineInput) Resize(width int) {
    input.Width = width
    input.Model.Width = width - 6 - 12
}

func (input *RequestLineInput) SwitchMethod() {
    input.SelectedMethod++

    if input.SelectedMethod == len(Methods) {
        input.SelectedMethod = 0
    }
}

func (input RequestLineInput) Render() string {
    var builder strings.Builder

    if (input.Width < 20) {
        return ""
    }

    m := Methods[input.SelectedMethod] + strings.Repeat(" ", 7 - len(Methods[input.SelectedMethod]))
    style.WrapItemInBorder(&builder, m, 7, 1)
    style.WrapItemInBorder(&builder, input.Model.View(), input.Model.Width+1, 1)

    return builder.String()
}
