package handlers

import (
	"filereader/services"
	"fmt"
)

func Start() {
	var filename string

	fmt.Println("Enter the file name to read: ")
	fmt.Scanln(&filename)

	content, err := services.ReadFileContent(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nFile content:")
	fmt.Println(content)
}