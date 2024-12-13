package repo

import(
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)


func CommitChanges(repoPath, message string) error{
	gitDir := filepath.Join(repoPath, ".gitclone")
	if _, err := os.Stat(gitDir); os.IsNotExist(err){
		return errors.New("not a valid repository")
	}

	stagingArea := filepath.Join(gitDir, "staging")
	files, err := os.ReadDir(stagingArea)
	if err != nil || len(files) == 0 {
		return errors.New("no files staged for commit")
	} 

   //generate a random id 
	commitID := fmt.Sprintf("%x", time.Now().UnixNano())
	commitDir := filepath.Join(gitDir, "commits", commitID)
	if err := os.MkdirAll(commitDir, 0755); err !=  nil{
		return fmt.Errorf("failed to create commit directory: %w", err)
	}

	for _, file := range files{
		src := filepath.Join(stagingArea, file.Name())
		dst := filepath.Join(commitDir, file.Name())
		if err := os.Rename(src, dst); err != nil{
			return fmt.Errorf("failed to move file %s to commit: %w", file.Name(), err)
		}
	}

	meta := fmt.Sprintf("Commit : %s\nTimestamp: %s\nMessage: %s\n", commitID,time.Now(), message)
	if err := os.WriteFile(filepath.Join(commitDir, "metadata.txt"), []byte(meta), 0644); err != nil{
		return fmt.Errorf("failed to write commit metadata: %w", err)
	}

	fmt.Printf("Commit created: %s\n", message)
	return nil
}

