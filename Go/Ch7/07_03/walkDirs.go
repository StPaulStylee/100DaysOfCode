package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root, _ := filepath.Abs(".") // current directory
	fmt.Println("Processing path:", root)

	err := filepath.Walk(root, processPath)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func processPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if path != "." { // if path is not equal to the root
		if info.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)
		}
	}

	return nil // I have to return something
}
