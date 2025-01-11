package analyzer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHumanReadableSize(t *testing.T) {
	tests := []struct {
		size     int64
		expected string
	}{
		{size: 512, expected: "512 B"},
		{size: 2048, expected: "2.00 KB"},
		{size: 5242880, expected: "5.00 MB"},
		{size: 1073741824, expected: "1.00 GB"},
	}

	for _, test := range tests {
		result := FormatSize(test.size)
		if result != test.expected {
			t.Errorf("Expected %s for size %d, got %s", test.expected, test.size, result)
		}
	}
}

func TestIsDirectory(t *testing.T) {
	tempDir := t.TempDir()

	// Test with a directory
	isDir, err := isDirectory(tempDir)
	if err != nil {
		t.Fatalf("Error checking if directory: %v", err)
	}
	if !isDir {
		t.Errorf("Expected %s to be a directory", tempDir)
	}

	// Test with a file
	tempFile := filepath.Join(tempDir, "testfile.txt")
	err = os.WriteFile(tempFile, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	isDir, err = isDirectory(tempFile)
	if err != nil {
		t.Fatalf("Error checking if file is directory: %v", err)
	}
	if isDir {
		t.Errorf("Expected %s to be a file, not a directory", tempFile)
	}

	// Test with a non-existent path
	nonExistentPath := filepath.Join(tempDir, "nonexistent")
	isDir, err = isDirectory(nonExistentPath)
	if err == nil {
		t.Fatalf("Expected error for non-existent path, got none")
	}
	if isDir {
		t.Errorf("Expected %s to not be a directory", nonExistentPath)
	}
}

func TestSafeJoin(t *testing.T) {
	root := "/home/user/projects"
	subPath := "myproject"

	// Test normal case
	result, err := safeJoin(root, subPath)
	if err != nil {
		t.Fatalf("SafeJoin failed: %v", err)
	}
	expected := filepath.Join(root, subPath)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test directory traversal attempt
	subPath = "../etc/passwd"
	result, err = safeJoin(root, subPath)
	if err == nil {
		t.Fatalf("Expected error for directory traversal attempt, got none")
	}

	// Test with root as empty
	emptyRoot := ""
	result, err = safeJoin(emptyRoot, subPath)
	if err == nil {
		t.Fatalf("Expected error for empty root path, got none")
	}

	// Test with subPath as empty
	result, err = safeJoin(root, "")
	if err != nil {
		t.Fatalf("SafeJoin failed with empty subPath: %v", err)
	}
	if result != root {
		t.Errorf("Expected %s, got %s", root, result)
	}
}
