package core

import (
	"encoding/json"
	"os"
)

// ReadDocumentFromFile reads a .glossa file from the given path and unmarshals it.
func ReadDocumentFromFile(path string) (*GlossDocument, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var doc GlossDocument
	err = json.Unmarshal(data, &doc)
	return &doc, err
}

// WriteDocumentToFile marshals the document to JSON and writes it to the given path.
func WriteDocumentToFile(doc *GlossDocument, path string) error {
	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
