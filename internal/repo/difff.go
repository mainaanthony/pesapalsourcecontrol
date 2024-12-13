package repo

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

//DiffFiles Compares two files line-by-line and prints the differences
func DiffFiles(file1Path, file2Path string) error{
	file1, err := os.Open(file1Path)
	if err != nil{
		return fmt.Errorf("failed to open file1: %w", err)
	}
	defer file1.Close()

	file2, err := os.Open(file2Path)
	if err != nil {
		return fmt.Errorf("failed to open file2: %w", err)
	}
	defer file2.Close()

	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	lineNumber := 1
	for scanner1.Scan() || scanner2.Scan(){
		line1 := ""
		line2 := ""

		if scanner1.Scan(){
			line1 = scanner1.Text()
		}

		if scanner2.Scan(){
			line2 = scanner1.Text()
		}
		if line1 != line2{
			fmt.Printf("Line %d:\n- %s\n+ %s\n", lineNumber, line1, line2)
		}
		lineNumber++
	}
	return nil
}


//DiffComits compares two commits by showing file-level differences.
func DiffCommits(repoPath, commit1, commit2 string) error{
	gitDir := filepath.Join(repoPath, ".gitclone", "commits")
	commit1Dir := filepath.Join(gitDir, commit1)
	commit2Dir := filepath.Join(gitDir, commit2)


	if _, err := os.Stat(commit1Dir); os.IsNotExist(err) {
		return fmt.Errorf("commit '%s' does not exist", commit1)
	}


	if _, err := os.Stat(commit2Dir); os.IsNotExist(err){
		return fmt.Errorf("commit '%s' does not exist", commit2)
	}

	files1,err := os.ReadDir(commit1Dir)
	if err != nil {
		return fmt.Errorf("failed to read commit1 directory: %w", err)
	}
	files2,err  := os.ReadDir(commit2Dir)

	if err != nil {
		return fmt.Errorf("failed to read commit2 directory: %w", err)
	}

	fmt.Printf("Differences between commits %s and %s:\n", commit1, commit2)

	for _, file := range files1{
		filePath1 := filepath.Join(commit1Dir, file.Name())
		filePath2 := filepath.Join(commit2Dir, file.Name())
		if _, err := os.Stat(filePath2); os.IsNotExist(err){
			fmt.Printf("- %s\n", file.Name())
		}else{
			DiffFiles(filePath1, filePath2)
		}
	}

	for _, file := range files2 {
		filePath1 := filepath.Join(commit1Dir, file.Name())
		if _, err := os.Stat(filePath1); os.IsNotExist(err) {
			fmt.Printf("+ %s\n", file.Name())
		}
	}

	return nil
		
	
}
