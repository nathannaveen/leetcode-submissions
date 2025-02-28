package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Process all directories in the root directory
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal("Error reading directory:", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()

		// Skip the problems directory and hidden directories
		if name == "problems" || strings.HasPrefix(name, ".") {
			continue
		}

		// Check if it's a problem directory (starts with a number)
		if len(name) > 0 && (name[0] >= '0' && name[0] <= '9') {
			// Create the new directory path
			newDir := filepath.Join("problems", name)

			// Create the new directory
			err = os.MkdirAll(newDir, 0755)
			if err != nil {
				log.Printf("Error creating directory %s: %v", newDir, err)
				continue
			}

			// Move all files from the old directory to the new one
			oldDir := name
			files, err := os.ReadDir(oldDir)
			if err != nil {
				log.Printf("Error reading directory %s: %v", oldDir, err)
				continue
			}

			moveSuccess := true
			for _, file := range files {
				oldPath := filepath.Join(oldDir, file.Name())
				newPath := filepath.Join(newDir, file.Name())
				err = os.Rename(oldPath, newPath)
				if err != nil {
					log.Printf("Error moving file %s to %s: %v", oldPath, newPath, err)
					moveSuccess = false
					break
				}
			}

			// Only try to remove the directory if all files were moved successfully
			if moveSuccess {
				// Check if directory is empty
				remaining, err := os.ReadDir(oldDir)
				if err != nil {
					log.Printf("Error checking if directory is empty %s: %v", oldDir, err)
					continue
				}

				if len(remaining) == 0 {
					err = os.RemoveAll(oldDir)
					if err != nil {
						log.Printf("Error removing directory %s: %v", oldDir, err)
						continue
					}
					log.Printf("Successfully moved and removed directory: %s", name)
				} else {
					log.Printf("Directory not empty after move, could not remove: %s", oldDir)
				}
			}
		}
	}
}
