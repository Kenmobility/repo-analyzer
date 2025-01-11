package analyzer

import (
	"fmt"
	"os"
	"path/filepath"
)

// formatSize converts a file size in bytes to a human-readable string format.
func FormatSize(sizeInBytes int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case sizeInBytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(sizeInBytes)/float64(GB))
	case sizeInBytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(sizeInBytes)/float64(MB))
	case sizeInBytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(sizeInBytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", sizeInBytes)
	}
}

// isDirectory checks if a given path is a directory.
// Returns true if the path is a directory, false otherwise.
func isDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

// safeJoin joins path segments and ensures the resulting path does not escape the root directory.
// This is useful for preventing directory traversal attacks.
func safeJoin(root, subPath string) (string, error) {
	if root == "" {
		return "", fmt.Errorf("root path cannot be empty")
	}

	joinedPath := filepath.Join(root, subPath)
	cleanedPath := filepath.Clean(joinedPath)
	if !filepath.HasPrefix(cleanedPath, root) {
		return "", fmt.Errorf("invalid path: %s", subPath)
	}
	return cleanedPath, nil
}
