package repo
import(
	"fmt"
	"os"
	"path/filepath"
)

//MergeBranches merges the target branch into the current branch
func MergeBranches(repoPath, targetBranch string) error{
   gitDir := filepath.Join(repoPath, ".gitclone")
   branchPath := filepath.Join(gitDir, "branches", targetBranch)

   if _, err := os.Stat(branchPath); os.IsNotExist(err){
	return fmt.Errorf("branch '%s does not exist", targetBranch)
   }

   headPath := filepath.Join(gitDir, "HEAD")
   currentBranch, err := os.ReadFile(headPath)
   if err != nil{
	return fmt.Errorf("failed to read HEAD: %w", err)
   }
   fmt.Printf("Merging branch '%s' .\n", targetBranch,string(currentBranch))

   if targetBranch == string(currentBranch) {
	fmt.Println("Cannot merge a branch into itself.")
	return nil
}

fmt.Printf("Merge successful. (This is a placeholder, conflicts are not resolved).\n")
return nil
}