package core

// GlossDocument represents the entire file. It contains metadata and a slice of GlossBlocks.
type GlossDocument struct {
	Version string       `json:"version"` // To handle future format migrations
	Config  DocConfig    `json:"config"`  // Document-specific settings
	Blocks  []GlossBlock `json:"blocks"`  // The main content of the document
}

// DocConfig holds settings for the document, like the number of lines and their labels.
type DocConfig struct {
	LineCount int      `json:"line_count"`
	LineLabels []string `json:"line_labels"` // e.g., ["Source", "Morphemes", "Gloss", "Translation"]
}

// GlossBlock represents a single "paragraph" of interlinear text.
// It is composed of a series of vertical columns, where each column is a word or morpheme.
type GlossBlock struct {
	Columns []GlossColumn `json:"columns"`
}

// GlossColumn holds the text for the different lines of a single, vertically-aligned unit.
// The length of the `Lines` slice is dictated by the `DocConfig.LineCount`.
type GlossColumn struct {
	Lines []string `json:"lines"` // e.g., Lines[0] is source, Lines[1] is morpheme, etc.
}
