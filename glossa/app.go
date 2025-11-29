package main

import (
	"context"
	"glossa/core"
	"glossa/export"
	"glossa/importer"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

// App struct holds the application's state and context.
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct.
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. We save the context
// so we can interact with the Wails runtime.
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// OpenFile prompts the user to select a file and returns the parsed document.
func (a *App) OpenFile() (*core.GlossDocument, error) {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Open Glossa File",
		Filters: []runtime.FileFilter{{DisplayName: "Glossa Files (*.glossa)", Pattern: "*.glossa"}},
	})
	if err != nil {
		return nil, err
	}
	if path == "" {
		return nil, nil // User cancelled
	}
	// We'll need to implement the file reading logic in the core package.
	// For now, let's assume it exists.
	return core.ReadDocumentFromFile(path)
}

// SaveFile saves the given document to a path chosen by the user.
func (a *App) SaveFile(doc *core.GlossDocument) (string, error) {
	path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Save Glossa File",
		DefaultFilename: "Untitled.glossa",
		Filters: []runtime.FileFilter{{DisplayName: "Glossa Files (*.glossa)", Pattern: "*.glossa"}},
	})
	if err != nil {
		return "", err
	}
	if path == "" {
		return "", nil // User cancelled
	}
	// We'll need to implement the file writing logic in the core package.
	// For now, let's assume it exists.
	return path, core.WriteDocumentToFile(doc, path)
}

// ExportToMarkdown converts the document to a Markdown string. The frontend
// will then save this string to a .md file.
func (a *App) ExportToMarkdown(doc *core.GlossDocument) (string, error) {
	return export.ToMarkdown(doc)
}

// ExportToPDF triggers the system's print dialog. The frontend is responsible
// for rendering the document in a print-friendly HTML/CSS format.
func (a *App) ExportToPDF() error {
	runtime.WindowPrint(a.ctx)
	return nil
}

// SaveMarkdownFile saves the document as a Markdown file.
func (a *App) SaveMarkdownFile(doc *core.GlossDocument) (string, error) {
	path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Save Markdown File",
		DefaultFilename: "Untitled.md",
		Filters: []runtime.FileFilter{{DisplayName: "Markdown Files (*.md)", Pattern: "*.md"}},
	})
	if err != nil {
		return "", err
	}
	if path == "" {
		return "", nil // User cancelled
	}

	content, err := export.ToMarkdown(doc)
	if err != nil {
		return "", err
	}

	return path, os.WriteFile(path, []byte(content), 0644)
}

// SavePDFFile saves the document as a PDF file.
func (a *App) SavePDFFile(doc *core.GlossDocument) (string, error) {
	path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Export PDF",
		Filters: []runtime.FileFilter{
			{DisplayName: "PDF Files (*.pdf)", Pattern: "*.pdf"},
		},
	})

	if err != nil || path == "" {
		return "", err
	}

	err = export.ToPDF(doc, path)
	if err != nil {
		return "", err
	}

	return path, nil
}

func (a *App) OpenMarkdownFile() (*core.GlossDocument, error) {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Open Markdown File",
		Filters: []runtime.FileFilter{
			{DisplayName: "Markdown Files (*.md)", Pattern: "*.md"},
		},
	})

	if err != nil || path == "" {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return importer.FromMarkdown(string(content))
}
