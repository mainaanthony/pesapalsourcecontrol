# The following is a step by step guide to run the project:

### In order to run the project, the user need:
1. Go installed on their machine.
2. Cloning the repository from your Git hosting service.
3. Running go mod tidy to fetch dependencies (if necessary).
4. Building the project using go build or running directly with go run.
5. Executing commands via the CLI to initialize a repo, add files, commit, view logs, etc.

### To run the project and use the distributed source control system you’ve built, a user will need the following:

#### 1. Go Installed
Since this is a Go project, the user needs to have Go installed on their machine to compile and run the project.
Installation Instructions:

They can install Go from the official [Go website](https://go.dev/dl/): .

After installing, they should verify the installation by running:
`go version`

This should output the Go version, confirming the installation was successful.
#### 2. Clone the Repository
The user needs to clone the repository to their local machine using Git: for instance
```
git clone https://your-repo-url.git
cd your-repo-directory
```

#### 3. Run go mod tidy (Optional, if dependencies are missing)
If you're using Go modules (go.mod), the user may need to run go mod tidy to download all dependencies. This is useful if there are any dependencies not included in the vendor folder or the repository.
Run this inside the repository directory:
`go mod tidy`

#### 4. Build the Project
To build the project, the user needs to run:
`go build -o gitclone cmd/main.go`
This will create an executable binary named gitclone (or whatever they specify in the -o flag).

#### 5. Run the Project
Once the build is successful, the user can run the project by executing the built binary:
`./gitclone`

Alternatively, they can run the project directly using:
`go run cmd/main.go`

Then you can now use the various commands (like init, add, commit, log, etc.) as outlined in the main.go file.


#### Now to the fun part:
# Testing the Source Control System (via CLI)
#### 1. Initialize a Repository
`go run cmd/main.go init myrepo`

#### 2. Add Files to the Repository

Create a file and add it to the repository:(please run the following lines of codes independently that is line by line)

```
cd myrepo
echo "Hello, SCM!" > file.txt
go run ../cmd/main.go add -file file.txt
```

This will add file.txt to the staging area.


#### 3. Commit Changes
`go run ../cmd/main.go commit -message "Initial commit"`

This will commit the staged changes with the message "Initial commit".


#### 4. View Commit History
`go run ../cmd/main.go log`


This will display the commit history.

#### 5. Create and Switch Branches((please run the following lines of codes independently that is line by line))
```
go run ../cmd/main.go branch -name feature-branch
go run ../cmd/main.go switch -name feature-branch
```
This creates and switches to a new branch called feature-branch.

#### 6. Merge Branches

`go run ../cmd/main.go merge -branch feature-branch`

This will merge the feature-branch into the current branch.

#### 7. Diff Between Commits
`go run ../cmd/main.go diff -commit1 <commit1ID> -commit2 <commit2ID>`

Replace <commit1ID> and <commit2ID> with actual commit IDs from the log to compare the differences.

### Additional Considerations:
File Permissions: Ensure that the user has the proper file system permissions to read/write to the repository directory and its files.
Go Version Compatibility: Make sure the Go version used is compatible with the project (Go 1.16 or later is recommended).