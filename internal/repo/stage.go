package repo


import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	
)

func AddFile(repoPath, filePath string) error {
    gitDir := filepath.Join(repoPath, ".gitclone")
    
    // Check if the repository exists
    if _, err := os.Stat(gitDir); os.IsNotExist(err) {
        return errors.New("not a valid repository")
    }

    // Check if the file exists
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return fmt.Errorf("file %s does not exist", filePath)
    }

    // Create the staging area if it doesn't exist
    stagingArea := filepath.Join(gitDir, "staging")
    if err := os.MkdirAll(stagingArea, 0755); err != nil {
        return fmt.Errorf("failed to create staging area: %w", err)
    }

    // Copy the file to the staging area
    dest := filepath.Join(stagingArea, filepath.Base(filePath))
    if err := copyFile(filePath, dest); err != nil {
        return fmt.Errorf("failed to copy file to staging area: %w", err)
    }

    // Print a success message to inform the user
    fmt.Printf("File '%s' successfully added to staging area.\n", filePath)

    return nil
}

//The copyFile function is a utility function that copies the contents of a file from a source (src) to a destination (dest). Here’s a detailed breakdown of how it works:


func copyFile(src, dest string) error{
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}

	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil{
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}