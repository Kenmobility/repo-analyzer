package analyzer

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileInfo represents the details of a file in the repository.
type FileInfo struct {
	Name string `json:"name"`
	Size string `json:"size"`
}

// FolderInfo represents the details of a folder, including its subfolders and files.
type FolderInfo struct {
	Name    string       `json:"name"`
	Files   []FileInfo   `json:"files"`
	Folders []FolderInfo `json:"folders"`
}

// AnalyzeFolder recursively analyzes the given folder and returns its structure
// along with the total size of files in the folder.
//
// Parameters:
// - rootPath: The root directory to analyze.
//
// Returns:
// - FolderInfo: A structured representation of the folder and its contents.
// - error: An error object if any issues occur during traversal.
func AnalyzeFolder(rootPath string) (FolderInfo, int64, error) {
	folderInfo := FolderInfo{
		Name:    filepath.Base(rootPath),
		Files:   []FileInfo{},   // Initialize as empty slice
		Folders: []FolderInfo{}, // Initialize as empty slice
	}

	entries, err := os.ReadDir(rootPath)
	if err != nil {
		return folderInfo, 0, fmt.Errorf("failed to read cloned repo directory %s: %w", rootPath, err)
	}

	var totalSize int64

	for _, entry := range entries {

		// Use safeJoin to ensure the entry path is within the rootPath.
		entryPath, err := safeJoin(rootPath, entry.Name())
		if err != nil {
			return folderInfo, 0, err
		}

		isDir, err := isDirectory(entryPath)
		if err != nil {
			return folderInfo, 0, err
		}

		if isDir {
			// Recursively analyze subfolder
			subFolder, subFolderSize, err := AnalyzeFolder(entryPath)
			if err != nil {
				return folderInfo, 0, err
			}
			folderInfo.Folders = append(folderInfo.Folders, subFolder)
			folderInfo.Name = entry.Name()
			totalSize += int64(subFolderSize)
		} else {
			// Analyze file
			info, err := entry.Info()
			if err != nil {
				return folderInfo, 0, err
			}
			fileSize := info.Size()
			totalSize += fileSize
			folderInfo.Files = append(folderInfo.Files, FileInfo{
				Name: entry.Name(),
				Size: FormatSize(fileSize),
			})
		}
	}

	return folderInfo, totalSize, nil
}
