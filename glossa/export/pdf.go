package export

import (
	"glossa/core"
	"github.com/go-pdf/fpdf"
)

func ToPDF(doc *core.GlossDocument) (*fpdf.Fpdf, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	lineHeight := 6.0
	pageWidth, _ := pdf.GetPageSize()
	leftMargin, _, rightMargin, _ := pdf.GetMargins()
	printableWidth := pageWidth - leftMargin - rightMargin

	for _, block := range doc.Blocks {
		// Calculate column widths
		var colWidths []float64
		for _, col := range block.Columns {
			maxWidth := 0.0
			for _, line := range col.Lines {
				w := pdf.GetStringWidth(line)
				if w > maxWidth {
					maxWidth = w
				}
			}
			colWidths = append(colWidths, maxWidth+2) // +2 for padding
		}

		// Check if block fits on current line (simplified: just wrap whole blocks for now if they are too wide, or wrap columns)
		// Actually, standard interlinear glossing wraps *sets* of columns.
		// Let's implement a simple column wrapping.

		currentX := leftMargin
		// Check if we need a new page for this block (rough estimate)
		if pdf.GetY() > 250 {
			pdf.AddPage()
			currentX = leftMargin
		}

		// We need to group columns that fit on a line.
		startIndex := 0
		for startIndex < len(block.Columns) {
			// Find how many columns fit
			endIndex := startIndex
			currentLineWidth := 0.0
			for endIndex < len(block.Columns) {
				if currentLineWidth+colWidths[endIndex] > printableWidth {
					break
				}
				currentLineWidth += colWidths[endIndex]
				endIndex++
			}

			// If even one column is too wide, force it (or handle it better, but for now force)
			if endIndex == startIndex {
				endIndex++
				currentLineWidth = colWidths[startIndex]
			}

			// Render these columns
			// We need to render line by line.
			// doc.Config.LineCount tells us how many lines per column.
			
			// Save Y position (not used for now)
			// startY := pdf.GetY()
			
			for lineIdx := 0; lineIdx < doc.Config.LineCount; lineIdx++ {
				pdf.SetX(currentX)
				for i := startIndex; i < endIndex; i++ {
					col := block.Columns[i]
					text := ""
					if lineIdx < len(col.Lines) {
						text = col.Lines[lineIdx]
					}
					
					// Set style based on line index (optional)
					if lineIdx == 0 {
						pdf.SetFont("Arial", "B", 12)
					} else if lineIdx == doc.Config.LineCount - 1 {
						pdf.SetFont("Arial", "I", 12)
					} else {
						pdf.SetFont("Arial", "", 12)
					}

					pdf.CellFormat(colWidths[i], lineHeight, text, "", 0, "L", false, 0, "")
				}
				pdf.Ln(lineHeight)
			}
			
			// Add some spacing after the block part
			pdf.Ln(lineHeight / 2)
			
			startIndex = endIndex
			
			// Check for page break
			if pdf.GetY() > 270 {
				pdf.AddPage()
			}
		}
		
		pdf.Ln(lineHeight)
	}

	return pdf, nil
}
