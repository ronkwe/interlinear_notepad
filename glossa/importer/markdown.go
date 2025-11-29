package importer

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"glossa/core"
	"strings"
)

func FromMarkdown(content string) (*core.GlossDocument, error) {
	doc := &core.GlossDocument{
		Version: "1.0",
		Config: core.DocConfig{
			LineCount: 4, // Default, will be updated based on rows found
			LineOptions: []core.LineOption{
				{Label: "Source", Visible: true},
				{Label: "Morphemes", Visible: true},
				{Label: "Gloss", Visible: true},
				{Label: "Translation", Visible: true},
			},
			ShowLabels: true,
		},
		Blocks: []core.GlossBlock{},
	}

	scanner := bufio.NewScanner(strings.NewReader(content))
	var currentBlock *core.GlossBlock
	var tableRows [][]string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			// End of block or empty line
			if currentBlock != nil && len(tableRows) > 0 {
				processTable(doc, currentBlock, tableRows)
				doc.Blocks = append(doc.Blocks, *currentBlock)
				currentBlock = nil
				tableRows = nil
			}
			continue
		}

		if strings.HasPrefix(line, "|") {
			// It's a table row
			if currentBlock == nil {
				currentBlock = &core.GlossBlock{
					Columns: []core.GlossColumn{},
				}
			}

			// Parse row
			// Remove leading/trailing pipes
			line = strings.Trim(line, "|")
			cells := strings.Split(line, "|")
			
			// Check if it's a separator row
			isSeparator := true
			for _, cell := range cells {
				if !strings.Contains(strings.TrimSpace(cell), "---") {
					isSeparator = false
					break
				}
			}

			if !isSeparator {
				// Clean up cells
				var rowData []string
				for _, cell := range cells {
					rowData = append(rowData, strings.TrimSpace(cell))
				}
				tableRows = append(tableRows, rowData)
			}
		}
	}

	// Process last block if exists
	if currentBlock != nil && len(tableRows) > 0 {
		processTable(doc, currentBlock, tableRows)
		doc.Blocks = append(doc.Blocks, *currentBlock)
	}

	// If no blocks found, return empty doc
	if len(doc.Blocks) == 0 {
		doc.Blocks = append(doc.Blocks, core.GlossBlock{
			Columns: []core.GlossColumn{{Lines: make([]string, doc.Config.LineCount)}},
		})
	}

	return doc, nil
}

func processTable(doc *core.GlossDocument, block *core.GlossBlock, rows [][]string) {
	if len(rows) == 0 {
		return
	}

	// Number of columns in the gloss block corresponds to number of cells in the table rows
	// Number of lines in the gloss block corresponds to number of table rows
	
	numLines := len(rows)
	// Update config line count if this block has more lines than default
	if numLines > doc.Config.LineCount {
		doc.Config.LineCount = numLines
		// Add default labels for extra lines
		for i := len(doc.Config.LineOptions); i < numLines; i++ {
			doc.Config.LineOptions = append(doc.Config.LineOptions, core.LineOption{
				Label: "Line " + string(rune('0'+i+1)), 
				Visible: true,
			})
		}
	}

	// Determine max columns in this table
	maxCols := 0
	for _, row := range rows {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	// Initialize columns
	block.Columns = make([]core.GlossColumn, maxCols)
	for i := range block.Columns {
		block.Columns[i].ID = generateID()
		block.Columns[i].Lines = make([]string, doc.Config.LineCount)
	}

	// Fill data
	// rows[i] is the i-th line (Source, Morph, etc.)
	// rows[i][j] is the content for the j-th column
	for lineIdx, row := range rows {
		for colIdx, cellContent := range row {
			if colIdx < len(block.Columns) {
				// Handle escaped pipes
				cellContent = strings.ReplaceAll(cellContent, "\\|", "|")
				if lineIdx < len(block.Columns[colIdx].Lines) {
					block.Columns[colIdx].Lines[lineIdx] = cellContent
				}
			}
		}
	}
}

func generateID() string {
	// Simple random ID generation
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "error-id"
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
