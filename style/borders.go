package style

import "strings"

const (
    topLeftCorner = "╭"
    topRightCorner = "╮"
    bottomLeftCorner = "╰"
    bottomRightCorner = "╯"
    verticalLine = "│"
    horizontalLine = "─"
)

func WrapItemInBorder(builder *strings.Builder, item string, innerWidth, innerHeight int) {
    horizontalBorder := strings.Repeat(horizontalLine, innerWidth)
    verticalBorder := strings.Repeat(verticalLine, innerHeight)

    builder.WriteString(topLeftCorner)
    builder.WriteString(horizontalBorder)
    builder.WriteString(topRightCorner)

    builder.WriteString("\n")

    builder.WriteString(verticalBorder)
    builder.WriteString(item)
    builder.WriteString(verticalBorder)

    builder.WriteString("\n")

    builder.WriteString(bottomLeftCorner)
    builder.WriteString(horizontalBorder)
    builder.WriteString(bottomRightCorner)

    builder.WriteString("\n")
}
