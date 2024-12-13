package repo

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// createBranch createsa new branch by cop[ying the current HEAD commit
func CreateBranch(repoPath, branchName string) error {
	gitDir := filepath.Join(repoPath, ".gitclone")
	branchesDir := filepath.Join(gitDir, "branches")

	if _, err := os.Stat(branchesDir); os.IsNotExist(err) {
		return errors.New("branches directory does not exist")
	}

	branchPath := filepath.Join(branchesDir, branchName)
	if _, err := os.Stat(branchPath); err == nil {
		return fmt.Errorf("branch '%s' already exists", branchName)
	}

	headPath := filepath.Join(gitDir, "HEAD")
	head, err := os.ReadFile(headPath)
	if err != nil {
		return fmt.Errorf("failed to read HEAD: %w", err)
	}

	if err := os.WriteFile(branchPath, head, 0644); err != nil {
		return fmt.Errorf("failed to create branch '%s':%w", branchName, err)
	}

	fmt.Printf("Branch '%s' created.\n", branchName)
	return nil
}

// SwitchBranch switches to an existing branch by updating HEAD
func SwitchBranch(repoPath, branchName string) error {
	gitDir := filepath.Join(repoPath, ".gitclone")
	branchPath := filepath.Join(gitDir, "branches", branchName)

	if _, err := os.Stat(branchPath); os.IsNotExist(err) {
		return fmt.Errorf("branch  '%s' does not exist", branchName)
	}

	headPath := filepath.Join(gitDir, "HEAD")
	branchContext, err := os.ReadFile(branchPath)
	if err != nil {
		return fmt.Errorf("failed to read branch '%s':%w", branchName, err)
	}

	if err := os.WriteFile(headPath, branchContext, 0644); err != nil {
		return fmt.Errorf("failed to switch to branch '%s': %w", branchName, err)
	}
	fmt.Printf("Switched to branch '%s'.\n", branchName)
	return nil
}
