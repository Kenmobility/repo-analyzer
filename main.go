package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/kenmobility/repo-analyzer/analyzer"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <repository-url>", os.Args[0])
	}

	repoURL := os.Args[1]
	tempDir, err := os.MkdirTemp("", "repo-*")
	if err != nil {
		log.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	fmt.Println("Cloning repository...")
	_, err = git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatalf("Error cloning repository: %v", err)
	}

	fmt.Println("Analyzing repository...")
	analyzedFolder, totalSize, err := analyzer.AnalyzeFolder(tempDir)
	if err != nil {
		log.Fatalf("Error analyzing repository: %v", err)
	}

	analysis := map[string]any{
		"clone_url": repoURL,
		"size":      analyzer.FormatSize(totalSize),
		"folders":   analyzedFolder,
	}

	fmt.Println("Outputting structure...")
	jsonOutput, err := outputJSON(analysis)
	if err != nil {
		fmt.Printf("Error generating JSON output: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(jsonOutput)
}

func outputJSON(data map[string]interface{}) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to format JSON: %w", err)
	}

	return string(jsonData), nil
}
