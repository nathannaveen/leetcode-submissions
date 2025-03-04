package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Path to the "0" problems directory
	zeroDir := "problems/0"

	// Read all entries in the directory
	entries, err := os.ReadDir(zeroDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		oldName := entry.Name()

		// Skip if it already starts with a number
		if strings.HasPrefix(oldName, "0-") {
			continue
		}

		// Create new name with "0-" prefix
		newName := "0-" + oldName

		// Create full paths
		oldPath := filepath.Join(zeroDir, oldName)
		newPath := filepath.Join(zeroDir, newName)

		// Rename the directory
		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("Error renaming %s to %s: %v\n", oldPath, newPath, err)
			continue
		}

		fmt.Printf("Renamed: %s -> %s\n", oldName, newName)
	}
}
