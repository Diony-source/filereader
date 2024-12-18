// Manages user interaction for file reading
package handlers

import (
	"filereader/services"
	"fmt"
	"runtime"
)

func Start() {
	var filename string

	// Get file name from the user
	fmt.Print("Enter the file name to search: ")
	fmt.Scanln(&filename)

	// Determine the root directory to start search
	rootDir := "/"
	if runtime.GOOS == "windows" {
		rootDir = "C:\\" // Root directory for Windows
	}

	fmt.Println("Searching for the file...")

	// Search for the file
	foundPath, err := services.FindFileInSystem(rootDir, filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File found at:", foundPath)

	// Read file content
	content, err := services.ReadFileContent(foundPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Display the file content
	fmt.Println("\nFile Content:")
	fmt.Println(content)
}
