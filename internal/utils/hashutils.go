package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// GenerateFileHash computes the SHA-256 hash of a file's contents.

func GenerateFileHash(filePath string)(string, error){
	file, err := os.Open(filePath)
	if err != nil{
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()


	hasher := sha256.New()

	if _, err := io.Copy(hasher, file); err != nil{
		return "", fmt.Errorf("failed to hash file contents: %w", err)
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

