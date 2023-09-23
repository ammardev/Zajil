package components

import "github.com/charmbracelet/lipgloss"

var (
    methods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
    style lipgloss.Style
)


type MethodSelector struct {
    selectedMethod int
}

func NewMethodSelector() MethodSelector {
    style = lipgloss.NewStyle().
        Width(9).
        AlignHorizontal(lipgloss.Center).
        Border(lipgloss.RoundedBorder())

    return MethodSelector{
        selectedMethod: 0,
    }
}

func (selector *MethodSelector) NextMethod() {
    selector.selectedMethod++

    if selector.selectedMethod == len(methods) {
        selector.selectedMethod = 0
    }
}

func (selector *MethodSelector) PreviousMethod() {
    selector.selectedMethod--

    if selector.selectedMethod < 0 {
        selector.selectedMethod = len(methods) - 1
    }
}

func (selector MethodSelector) Render() string {
    return style.Render(methods[selector.selectedMethod])
}
