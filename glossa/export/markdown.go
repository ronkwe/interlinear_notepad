package export

import (
	"glossa/core"
	"strings"
)

func ToMarkdown(doc *core.GlossDocument) (string, error) {
	var sb strings.Builder

	for _, block := range doc.Blocks {
		if len(block.Columns) == 0 {
			continue
		}

		lineCount := doc.Config.LineCount
		lines := make([]strings.Builder, lineCount)

		for _, col := range block.Columns {
			for i := 0; i < lineCount; i++ {
				if i < len(col.Lines) {
					// Add padding or just space? Just space for now.
					lines[i].WriteString(col.Lines[i] + " ")
				}
			}
		}

		for i := 0; i < lineCount; i++ {
			sb.WriteString(strings.TrimSpace(lines[i].String()) + "\n")
		}
		sb.WriteString("\n")
	}

	return sb.String(), nil
}
