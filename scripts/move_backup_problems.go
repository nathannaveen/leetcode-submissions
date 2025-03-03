package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Create the problems/0 directory
	zeroDir := filepath.Join("problems", "0")
	err := os.MkdirAll(zeroDir, 0755)
	if err != nil {
		log.Fatal("Error creating zero directory:", err)
	}

	// Read all directories from problems_backup
	entries, err := os.ReadDir("problems_backup")
	if err != nil {
		log.Fatal("Error reading problems_backup directory:", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		oldPath := filepath.Join("problems_backup", name)
		newPath := filepath.Join(zeroDir, name)

		// Move the directory to problems/0
		err = os.Rename(oldPath, newPath)
		if err != nil {
			log.Printf("Error moving directory %s to %s: %v", oldPath, newPath, err)
			continue
		}

		log.Printf("Successfully moved %s to problems/0", name)
	}
}
