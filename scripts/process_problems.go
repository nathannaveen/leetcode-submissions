package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// getHundredGroup calculates the hundred group for a problem number
func getHundredGroup(num int) (int, int) {
	hundredStart := ((num-1)/100)*100 + 1
	hundredEnd := hundredStart + 99
	return hundredStart, hundredEnd
}

// moveToZeroDirectory moves a problem to the zero directory
func moveToZeroDirectory(problemPath string, name string) error {
	zeroDir := filepath.Join("problems", "0")
	err := os.MkdirAll(zeroDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating zero directory: %v", err)
	}
	newPath := filepath.Join(zeroDir, name)
	if problemPath == newPath {
		return nil
	}
	err = os.Rename(problemPath, newPath)
	if err != nil {
		return fmt.Errorf("error moving directory %s to %s: %v", problemPath, newPath, err)
	}
	log.Printf("Moved problem %s to zero directory", name)
	return nil
}

// moveToHundredGroup moves a problem directory to its correct hundred group
func moveToHundredGroup(problemPath string, name string) error {
	// First try to get the problem number
	parts := strings.Split(name, "-")
	if len(parts) == 0 {
		return moveToZeroDirectory(problemPath, name)
	}

	// Try to get problem number
	num, err := strconv.Atoi(parts[0])
	if err != nil {
		return moveToZeroDirectory(problemPath, name)
	}

	// Calculate which hundred group this belongs to
	hundredStart, hundredEnd := getHundredGroup(num)
	hundredDir := filepath.Join("problems", fmt.Sprintf("%d-%d", hundredStart, hundredEnd))

	// Create the hundred directory if it doesn't exist
	err = os.MkdirAll(hundredDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating hundred directory %s: %v", hundredDir, err)
	}

	// Move the problem directory to its hundred group
	newPath := filepath.Join(hundredDir, name)

	// If the problem is already in the correct location, skip it
	if problemPath == newPath {
		return nil
	}

	err = os.Rename(problemPath, newPath)
	if err != nil {
		return fmt.Errorf("error moving directory %s to %s: %v", problemPath, newPath, err)
	}

	log.Printf("Moved problem %s to group %d-%d", name, hundredStart, hundredEnd)
	return nil
}

// processNewProblems moves problems from root to their correct hundred groups
func processNewProblems() {
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

		// Move all files to a temporary location in problems directory
		tempDir := filepath.Join("problems", name)

		// Create the temporary directory
		err = os.MkdirAll(tempDir, 0755)
		if err != nil {
			log.Printf("Error creating directory %s: %v", tempDir, err)
			continue
		}

		// Move all files from the old directory to the temporary one
		oldDir := name
		files, err := os.ReadDir(oldDir)
		if err != nil {
			log.Printf("Error reading directory %s: %v", oldDir, err)
			continue
		}

		moveSuccess := true
		for _, file := range files {
			oldPath := filepath.Join(oldDir, file.Name())
			newPath := filepath.Join(tempDir, file.Name())
			err = os.Rename(oldPath, newPath)
			if err != nil {
				log.Printf("Error moving file %s to %s: %v", oldPath, newPath, err)
				moveSuccess = false
				break
			}
		}

		if moveSuccess {
			// Move from temporary location to correct hundred group or zero directory
			err = moveToHundredGroup(tempDir, name)
			if err != nil {
				log.Printf("Error organizing problem %s: %v", name, err)
				continue
			}

			// Remove the old directory
			remaining, err := os.ReadDir(oldDir)
			if err != nil {
				log.Printf("Error checking if directory is empty %s: %v", oldDir, err)
				continue
			}

			if len(remaining) == 0 {
				err = os.RemoveAll(oldDir)
				if err != nil {
					log.Printf("Error removing directory %s: %v", oldDir, err)
				}
			}
		}
	}
}

// reorganizeExistingProblems ensures all problems in the problems directory are in correct groups
func reorganizeExistingProblems() {
	problemsDir := "problems"
	entries, err := os.ReadDir(problemsDir)
	if err != nil {
		log.Fatal("Error reading problems directory:", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()

		// Skip the zero directory
		if name == "0" {
			continue
		}

		// Check if this is a hundred group directory (format: "1-100", "101-200", etc.)
		isHundredGroup := false
		if parts := strings.Split(name, "-"); len(parts) == 2 {
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil && start%100 == 1 && end == start+99 {
				// This is a hundred group directory, process its contents
				isHundredGroup = true
				groupPath := filepath.Join(problemsDir, name)
				problems, err := os.ReadDir(groupPath)
				if err != nil {
					log.Printf("Error reading group directory %s: %v", groupPath, err)
					continue
				}

				for _, problem := range problems {
					if problem.IsDir() {
						problemPath := filepath.Join(groupPath, problem.Name())
						err = moveToHundredGroup(problemPath, problem.Name())
						if err != nil {
							log.Printf("Error reorganizing problem %s: %v", problem.Name(), err)
						}
					}
				}
			}
		}

		// If not a hundred group, treat it as a problem directory
		if !isHundredGroup {
			problemPath := filepath.Join(problemsDir, name)
			err = moveToHundredGroup(problemPath, name)
			if err != nil {
				log.Printf("Error organizing problem %s: %v", name, err)
			}
		}
	}
}

func main() {
	// First process any new problems from the root directory
	processNewProblems()

	// Then ensure all problems in the problems directory are organized correctly
	reorganizeExistingProblems()
}
