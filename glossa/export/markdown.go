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
		// We need to transpose the data: columns are vertical in memory, but horizontal in Markdown table
		// Actually, standard Markdown tables are row-based.
		// Header | Header
		// --- | ---
		// Cell | Cell
		//
		// In our case, a "Gloss Column" (Word + Morph + Gloss) is a vertical slice.
		// But in a table, we probably want:
		// | Word 1 | Word 2 | ...
		// | Morph 1 | Morph 2 | ...
		// | Gloss 1 | Gloss 2 | ...
		//
		// So each row in the table corresponds to a "Line" type (Source, Morphemes, etc.)
		// And each column in the table corresponds to a "Gloss Column".

		// 1. Determine number of columns
		numCols := len(block.Columns)

		// 2. Iterate through lines (rows)
		for lineIdx := 0; lineIdx < lineCount; lineIdx++ {
			// Write row content
			sb.WriteString("|")
			for colIdx := 0; colIdx < numCols; colIdx++ {
				content := ""
				if lineIdx < len(block.Columns[colIdx].Lines) {
					content = block.Columns[colIdx].Lines[lineIdx]
				}
				// Escape pipes in content
				content = strings.ReplaceAll(content, "|", "\\|")
				sb.WriteString(" " + content + " |")
			}
			sb.WriteString("\n")

			// Write separator row after the first row (header)
			// Actually, Markdown tables require a header row and a separator row.
			// We can treat the first line (Source) as the header.
			if lineIdx == 0 {
				sb.WriteString("|")
				for colIdx := 0; colIdx < numCols; colIdx++ {
					sb.WriteString(" --- |")
				}
				sb.WriteString("\n")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String(), nil
}
