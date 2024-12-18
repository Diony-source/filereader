// Contains the logic for reading files and searching across directories
package services

import (
	"fmt"
	"os"
	"path/filepath"
)

// FindFileInSystem searches for the file across the entire system
func FindFileInSystem(rootDir, targetFile string) (string, error) {
	var foundPath string

	// Walk through all directories and files starting from rootDir
	err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // Skip directories we cannot access
		}

		// Check if the current file matches the target file
		if !d.IsDir() && d.Name() == targetFile {
			foundPath = path
			return filepath.SkipAll // Stop walking once file is found
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if foundPath == "" {
		return "", fmt.Errorf("file '%s' not found in the system", targetFile)
	}
	return foundPath, nil
}

// ReadFileContent reads the content of the given file path
func ReadFileContent(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("could not read file: %w", err)
	}
	return string(content), nil
}
