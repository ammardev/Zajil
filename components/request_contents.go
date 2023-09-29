package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type RequestContents struct {
    width int
    style lipgloss.Style
    activeTabStyle lipgloss.Style
    tabStyle lipgloss.Style
    tabGapStyle lipgloss.Style
    tabContentStyle lipgloss.Style
}

func NewRequestContents(width int) RequestContents {
    activeTabBorder := lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomRight:  "╰",
		BottomLeft: "╯",
    }

    tabBorder := lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomRight:  "┴",
		BottomLeft: "┴",
    }

    return RequestContents{
        style: lipgloss.NewStyle().Width(width),
        activeTabStyle: lipgloss.NewStyle().Border(activeTabBorder),
        tabStyle: lipgloss.NewStyle().Border(tabBorder),
        tabGapStyle: lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), false, false, true, false),
        tabContentStyle: lipgloss.NewStyle().Width(width).Height(18).Border(lipgloss.RoundedBorder()),
    }
}

func (rc *RequestContents) Resize(width int) {
    rc.style.Width(width)
    rc.tabContentStyle.Width(width-2)
}

func (rc RequestContents) Render() string {
    if rc.style.GetWidth() - 30 < 0 {
        return ""
    }
    return rc.style.Render(
        lipgloss.JoinHorizontal(
            lipgloss.Center,
            rc.activeTabStyle.Render(" Headers (H) "),
            rc.tabStyle.Render(" Body (B) "),
            rc.tabGapStyle.Render(strings.Repeat(" ", rc.style.GetWidth() - 28)),
        ),
        rc.tabContentStyle.Render("Authorization: Bearer 123\nContent-Type: application/json\nAccept: application/json"),
    )
}
