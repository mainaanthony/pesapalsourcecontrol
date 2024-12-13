package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mainaanthony/ppsc/internal/repo"
)

func main() {
	// Define commands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
	branchCmd := flag.NewFlagSet("branch", flag.ExitOnError)
	switchBranchCmd := flag.NewFlagSet("switch", flag.ExitOnError)
	diffCmd := flag.NewFlagSet("diff", flag.ExitOnError)
	mergeCmd := flag.NewFlagSet("merge", flag.ExitOnError)

	// Command arguments
	addFile := addCmd.String("file", "", "File to add to staging")
	commitMessage := commitCmd.String("message", "", "Commit message")
	createBranch := branchCmd.String("name", "", "Branch name")
	switchBranch := switchBranchCmd.String("name", "", "Branch name to switch to")
	diffCommit1 := diffCmd.String("commit1", "", "First commit ID for diff")
	diffCommit2 := diffCmd.String("commit2", "", "Second commit ID for diff")
	mergeBranch := mergeCmd.String("branch", "", "Branch to merge")

	// Validate commands
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  init [directory] - Initialize a new repository")
		fmt.Println("  add -file [file] - Stage a file for commit")
		fmt.Println("  commit -message [message] - Commit staged files")
		fmt.Println("  log - View commit history")
		fmt.Println("  branch -name [name] - Create a new branch")
		fmt.Println("  switch -name [name] - Switch to an existing branch")
		fmt.Println("  diff -commit1 [id] -commit2 [id] - View differences between commits")
		fmt.Println("  merge -branch [name] - Merge a branch into the current branch")
		os.Exit(1)
	}

	repoPath := "." // Assume current directory for simplicity

	switch os.Args[1] {
	case "init":
		if len(os.Args) < 3 {
			fmt.Println("Usage: init [directory]")
			os.Exit(1)
		}
		dir := os.Args[2]
		if err := repo.InitRepository(dir); err != nil {
			log.Fatalf("Error initializing repository: %v", err)
		}

	case "add":
		addCmd.Parse(os.Args[2:])
		if *addFile == "" {
			log.Fatalf("Error: -file argument is required")
		}
		if err := repo.AddFile(repoPath, *addFile); err != nil {
			log.Fatalf("Error adding file: %v", err)
		}

	case "commit":
		commitCmd.Parse(os.Args[2:])
		if *commitMessage == "" {
			log.Fatalf("Error: -message argument is required")
		}
		if err := repo.CommitChanges(repoPath, *commitMessage); err != nil {
			log.Fatalf("Error committing changes: %v", err)
		}

	case "log":
		if err := repo.ViewLog(repoPath); err != nil {
			log.Fatalf("Error viewing log: %v", err)
		}

	case "branch":
		branchCmd.Parse(os.Args[2:])
		if *createBranch == "" {
			log.Fatalf("Error: -name argument is required")
		}
		if err := repo.CreateBranch(repoPath, *createBranch); err != nil {
			log.Fatalf("Error creating branch: %v", err)
		}

	case "switch":
		switchBranchCmd.Parse(os.Args[2:])
		if *switchBranch == "" {
			log.Fatalf("Error: -name argument is required")
		}
		if err := repo.SwitchBranch(repoPath, *switchBranch); err != nil {
			log.Fatalf("Error switching branch: %v", err)
		}

	case "diff":
		diffCmd.Parse(os.Args[2:])
		if *diffCommit1 == "" || *diffCommit2 == "" {
			log.Fatalf("Error: -commit1 and -commit2 arguments are required")
		}
		if err := repo.DiffCommits(repoPath, *diffCommit1, *diffCommit2); err != nil {
			log.Fatalf("Error generating diff: %v", err)
		}

	case "merge":
		mergeCmd.Parse(os.Args[2:])
		if *mergeBranch == "" {
			log.Fatalf("Error: -branch argument is required")
		}
		if err := repo.MergeBranches(repoPath, *mergeBranch); err != nil {
			log.Fatalf("Error merging branches: %v", err)
		}

	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}
