package services

import (
	"fmt"
	"os"
)

func ReadFileContent(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("could not read file: %w", err)
	}
	return string(content), nil
}