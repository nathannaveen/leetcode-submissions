package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"leetcode-db/internal/database"
	"leetcode-db/internal/parser"
)

func main() {
	// Get database configuration
	config := database.DefaultConfig()

	// Override with environment variables if present
	if host := os.Getenv("DB_HOST"); host != "" {
		config.Host = host
	}
	if user := os.Getenv("DB_USER"); user != "" {
		config.User = user
	}
	if pass := os.Getenv("DB_PASSWORD"); pass != "" {
		config.Password = pass
	}
	if name := os.Getenv("DB_NAME"); name != "" {
		config.DBName = name
	}

	// Create database connection
	db, err := database.NewConnection(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create repository
	repo := database.NewRepository(db)

	// Get the root directory (parent of leetcode-db)
	rootDir, err := filepath.Abs("..")
	if err != nil {
		log.Fatalf("Failed to get root directory: %v", err)
	}

	// Create parser
	parser := parser.NewParser(rootDir)

	// Find all problem directories
	dirs, err := parser.FindProblemDirs()
	if err != nil {
		log.Fatalf("Failed to find problem directories: %v", err)
	}

	fmt.Printf("Found %d problem directories\n", len(dirs))

	// Process each directory
	var problemCount, solutionCount int
	for _, dir := range dirs {
		problem, solutions, err := parser.ParseProblem(dir)
		if err != nil {
			log.Printf("Failed to parse problem in %s: %v", dir, err)
			continue
		}

		// Store problem
		if err := repo.StoreProblem(problem); err != nil {
			log.Printf("Failed to store problem %d: %v", problem.ID, err)
			continue
		}
		problemCount++

		// Store solutions
		for _, solution := range solutions {
			if err := repo.StoreSolution(solution); err != nil {
				log.Printf("Failed to store solution for problem %d: %v", problem.ID, err)
				continue
			}
			solutionCount++
		}
	}

	fmt.Printf("Successfully imported:\n")
	fmt.Printf("- %d problems\n", problemCount)
	fmt.Printf("- %d solutions\n", solutionCount)
}
