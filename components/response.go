package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	//tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ResponseView struct {
    textinput.Model
    Width int
    style lipgloss.Style
}
var a lipgloss.Style

func NewResponseView(width int) ResponseView {
    response := ResponseView{
        Model: textinput.New(),
        style: lipgloss.NewStyle().Width(width).Border(lipgloss.RoundedBorder()),
    }
    response.Model.Prompt = ""

    return response
}

func (view *ResponseView) Resize(width, height int) {
    view.style.Width(width)
    view.style.Height(height)
}

func (view ResponseView) Render() string {
    return view.style.Render(
        "hello",
    )
}
