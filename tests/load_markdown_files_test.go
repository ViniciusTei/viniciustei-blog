package tests

import (
	"os"
	"testing"

	"github.com/ViniciusTei/viniciustei-blog/internal/handlers"
)

func TestLoadMarkdownFiles(t *testing.T) {
	// Create a temporary directory with test Markdown files
	dir := t.TempDir()
	os.WriteFile(dir+"/test1.md", []byte("# Title\nContent"), 0644)

	// Call function to load files
	files, err := handlers.LoadMarkdownFiles(dir)
	if err != nil {
		t.Fatalf("Failed to load articles: %v", err)
	}

	// Assert files were loaded correctly
	if len(files) != 1 {
		t.Fatalf("Expected 1 file, got %d", len(files))
	}

	if files[0].Title != "Title" {
		t.Errorf("Expected title 'Title', got '%s'", files[0].Title)
	}
}
