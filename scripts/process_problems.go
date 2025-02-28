package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// moveProblemsToRoot moves all problem directories to the problems directory
func moveProblemsToRoot() {
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

// organizeProblemsByHundreds organizes problems into groups of 100
func organizeProblemsByHundreds() {
	problemsDir := "problems"
	entries, err := os.ReadDir(problemsDir)
	if err != nil {
		log.Fatal("Error reading problems directory:", err)
	}

	// First, create all the hundred directories we need
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		parts := strings.Split(name, "-")
		if len(parts) == 0 {
			continue
		}

		// Get problem number
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Printf("Error parsing problem number from %s: %v", name, err)
			continue
		}

		// Calculate which hundred group this belongs to
		hundredStart := ((num-1)/100)*100 + 1
		hundredEnd := hundredStart + 99
		hundredDir := filepath.Join(problemsDir, fmt.Sprintf("%d-%d", hundredStart, hundredEnd))

		err = os.MkdirAll(hundredDir, 0755)
		if err != nil {
			log.Printf("Error creating hundred directory %s: %v", hundredDir, err)
			continue
		}
	}

	// Now move all problem directories into their respective hundred directories
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		parts := strings.Split(name, "-")
		if len(parts) == 0 {
			continue
		}

		// Get problem number
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Printf("Error parsing problem number from %s: %v", name, err)
			continue
		}

		// Calculate which hundred group this belongs to
		hundredStart := ((num-1)/100)*100 + 1
		hundredEnd := hundredStart + 99
		hundredDir := filepath.Join(problemsDir, fmt.Sprintf("%d-%d", hundredStart, hundredEnd))

		// Move the problem directory to its hundred group
		oldPath := filepath.Join(problemsDir, name)
		newPath := filepath.Join(hundredDir, name)

		err = os.Rename(oldPath, newPath)
		if err != nil {
			log.Printf("Error moving directory %s to %s: %v", oldPath, newPath, err)
			continue
		}

		log.Printf("Moved problem %s to group %d-%d", name, hundredStart, hundredEnd)
	}
}

func main() {
	// First move all problems to the problems directory
	moveProblemsToRoot()

	// Then organize them into hundreds
	organizeProblemsByHundreds()
}
