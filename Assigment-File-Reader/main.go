package main

import (
	"fmt"
	"os"
	"io"
)

func main()  {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the path to a text file as a command-line argument.")
		return
	}

	// Get the file path from the command-line arguments
	filePath := os.Args[1]

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Use io.Copy to copy the file contents to os.Stdout (terminal)
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
}