package export

import (
	"glossa/core"
	"os"
	"testing"
)

func TestToPDF(t *testing.T) {
	// Change working directory to project root so assets/fonts is found
	os.Chdir("..")

	doc := &core.GlossDocument{
		Config: core.DocConfig{
			LineCount: 2,
			LineOptions: []core.LineOption{
				{Label: "SOU", Visible: true},
				{Label: "TRA", Visible: true},
			},
			ShowLabels: true,
		},
		Blocks: []core.GlossBlock{
			{
				Columns: []core.GlossColumn{
					{Lines: []string{"Shé:kon", "Hello"}},
					{Lines: []string{"kenòn:we's", "I like it"}},
					{Lines: []string{"ne", "the"}},
					{Lines: []string{"thí:ken", "that"}},
					{Lines: []string{"kà:sere", "car"}},
				},
			},
		},
	}

	outputPath := "test_output.pdf"
	err := ToPDF(doc, outputPath)
	if err != nil {
		t.Fatalf("ToPDF failed: %v", err)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatalf("Output file was not created")
	}

	// Clean up
	os.Remove(outputPath)
}
