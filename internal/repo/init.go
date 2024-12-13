package repo

import (
	"fmt"
	"os"
	"path/filepath"
)

// InitRepository initializes a new repository in the specified path
func InitRepository(path string) error {
	// Create the repository directory if it doesn't exist
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}

	// Define the path for the .gitclone directory
	gitDir := filepath.Join(path, ".gitclone")

	// Check if the repository already exists
	if _, err := os.Stat(gitDir); err == nil {
		return fmt.Errorf("repository already exists in %s", path)
	}

	// Create subdirectories for staging, commits, and branches
	subDirs := []string{"staging", "commits", "branches"}
	for _, dir := range subDirs {
		if err := os.MkdirAll(filepath.Join(gitDir, dir), 0755); err != nil {
			return fmt.Errorf("failed to create %s: %w", dir, err)
		}
	}


	// Create the default 'main' branch in the branches directory
	mainBranchPath := filepath.Join(gitDir, "branches", "main")
	if err := os.WriteFile(mainBranchPath, []byte{}, 0644); err != nil {
		return fmt.Errorf("failed to create main branch: %w", err)
	}
	
	// Create the HEAD file and point it to the default branch (main)
	headFilePath := filepath.Join(gitDir, "HEAD")
	headContent := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(headFilePath, headContent, 0644); err != nil {
		return fmt.Errorf("failed to create HEAD file: %w", err)
	}

	// Print a success message
	fmt.Printf("Initialized empty repository in %s\n", gitDir)
	return nil
}
