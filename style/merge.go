package style

import (
	"strings"
)

func CombineIntoOneLine(item string, numberOfComponents int, height int) string {
    var builder strings.Builder
    lines := strings.Split(strings.TrimRight(item, "\n"), "\n")

    for line := 0; line < height; line++ {
        for component := 0; component < numberOfComponents; component++ {
            builder.WriteString(lines[line+(height*component)])
        }
        builder.WriteString("\n")
    }

    return builder.String()
}
