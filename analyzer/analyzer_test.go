package analyzer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAnalyzeFolder_EmptyFolder(t *testing.T) {
	tempDir := t.TempDir()
	folder, _, err := AnalyzeFolder(tempDir)
	if err != nil {
		t.Fatalf("AnalyzeFolder failed: %v", err)
	}

	if len(folder.Files) != 0 || len(folder.Folders) != 0 {
		t.Errorf("Expected empty files and folders, got %+v", folder)
	}
}

func TestAnalyzeFolder_WithFilesAndFolders(t *testing.T) {
	tempDir := t.TempDir()

	// Create a subfolder and files
	subDir := filepath.Join(tempDir, "subdir")
	err := os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(subDir, "file2.txt")

	err = os.WriteFile(file1, []byte("Hello"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file1: %v", err)
	}

	err = os.WriteFile(file2, []byte("World"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file2: %v", err)
	}

	// Analyze folder
	folder, _, err := AnalyzeFolder(tempDir)
	if err != nil {
		t.Fatalf("AnalyzeFolder failed: %v", err)
	}

	if len(folder.Files) != 1 || folder.Files[0].Name != "file1.txt" {
		t.Errorf("Expected file 'file1.txt', got %+v", folder.Files)
	}

	if len(folder.Folders) != 1 || folder.Folders[0].Name != "subdir" {
		t.Errorf("Expected folder 'subdir', got %+v", folder.Folders)
	}

	if folder.Folders[0].Files[0].Name != "file2.txt" {
		t.Errorf("Expected file 'file2.txt' in subdir, got %+v", folder.Folders[0].Files)
	}
}
