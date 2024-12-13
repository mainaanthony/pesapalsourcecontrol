package repo


import(
	"fmt"
	"os"
	"path/filepath"
)

func ViewLog(repoPath string) error{
	commitDir := filepath.Join(repoPath, ".gitclone", "commits")
	commitDirs, err := os.ReadDir(commitDir)
	if err != nil{
		return fmt.Errorf("failed to read commit directory: %w", err)
	}

	for _, dir := range commitDirs{
		metaFile := filepath.Join(commitDir, dir.Name(), "metadata.txt")
		content, err := os.ReadFile(metaFile)
		if err != nil{
			return fmt.Errorf("failed to read commit metadata: %w", err)
		}
		fmt.Println(string(content))
	}
	return nil
}