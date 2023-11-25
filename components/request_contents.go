package components

import (
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/lipgloss"
)

const requestContentsTabViewOuterHeight = 3

type RequestContents struct {
    width int
    style lipgloss.Style
    activeTabStyle lipgloss.Style
    tabStyle lipgloss.Style
    tabGapStyle lipgloss.Style
    tabContentStyle lipgloss.Style
    HeadersTextInput textarea.Model
    BodyTextInput textarea.Model
    activeTab int
}

func NewRequestContents() RequestContents {
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
        style: lipgloss.NewStyle().Bold(false),
        activeTabStyle: lipgloss.NewStyle().Border(activeTabBorder),
        tabStyle: lipgloss.NewStyle().Border(tabBorder),
        tabGapStyle: lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), false, false, true, false),
        tabContentStyle: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()),
        HeadersTextInput: textarea.New(),
        BodyTextInput: textarea.New(),
        activeTab: 0,
    }
}

func (rc *RequestContents) Resize(width, height int) {
    rc.style.Width(width)
    rc.style.Height(height)

    rc.tabContentStyle.Width(width - borderPadding)
    rc.tabContentStyle.Height(height - borderPadding - requestContentsTabViewOuterHeight)
}

// TODO: fix this ugly code
func (rc *RequestContents) ActivateTab(tabIndex int) {
    rc.activeTab = tabIndex
    if tabIndex == 0 {
        rc.HeadersTextInput.Focus()
    } else {
        rc.BodyTextInput.Focus()
    }
}

func (rc RequestContents) Render() string {
    if rc.style.GetWidth() - 30 < 0 {
        return ""
    }

    headerTab := rc.tabStyle
    bodyTab := rc.tabStyle

    activeView := ""
    if rc.activeTab == 0 {
        activeView = rc.HeadersTextInput.View()
        headerTab = rc.activeTabStyle
    } else {
        activeView = rc.BodyTextInput.View()
        bodyTab = rc.activeTabStyle
    }

    return rc.style.Render(
        lipgloss.JoinHorizontal(
            lipgloss.Center,
            headerTab.Render(" Headers (H) "),
            bodyTab.Render(" Body (B) "),
            rc.tabGapStyle.Render(strings.Repeat(" ", rc.style.GetWidth() - 28)),
        ),
        rc.tabContentStyle.Render(
            activeView,
        ),
    )
}
