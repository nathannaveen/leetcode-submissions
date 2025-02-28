package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Problem struct {
	Title      string `json:"title"`
	TitleSlug  string `json:"titleSlug"`
	Difficulty string `json:"difficulty"`
	Number     int    `json:"questionId"`
}

type ProblemListResponse struct {
	Data struct {
		Problems []Problem `json:"problemsetQuestionList"`
	} `json:"data"`
}

func fetchProblems() (map[string]int, error) {
	query := `{
		problemsetQuestionList(limit: 5000) {
			questionId
			title
			titleSlug
			difficulty
		}
	}`

	jsonData := map[string]interface{}{
		"query": query,
	}

	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("POST", "https://leetcode.com/graphql", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	req.Header.Set("Origin", "https://leetcode.com")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("x-csrftoken", os.Getenv("LEETCODE_CSRF_TOKEN"))
	req.Header.Set("Cookie", fmt.Sprintf("csrftoken=%s; LEETCODE_SESSION=%s", os.Getenv("LEETCODE_CSRF_TOKEN"), os.Getenv("LEETCODE_SESSION")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ProblemListResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	slugToNumber := make(map[string]int)
	for _, p := range result.Data.Problems {
		slugToNumber[p.TitleSlug] = p.Number
	}

	return slugToNumber, nil
}

func normalizeTitle(title string) string {
	// Convert to lowercase
	title = strings.ToLower(title)
	// Replace spaces and special characters with underscore
	title = strings.ReplaceAll(title, " ", "_")
	title = strings.ReplaceAll(title, "-", "_")
	title = strings.ReplaceAll(title, "'", "")
	title = strings.ReplaceAll(title, ".", "")
	title = strings.ReplaceAll(title, ",", "")
	title = strings.ReplaceAll(title, "(", "")
	title = strings.ReplaceAll(title, ")", "")
	return title
}

func main() {
	if os.Getenv("LEETCODE_CSRF_TOKEN") == "" || os.Getenv("LEETCODE_SESSION") == "" {
		fmt.Println("Please set LEETCODE_CSRF_TOKEN and LEETCODE_SESSION environment variables")
		fmt.Println("You can get these values from your browser after logging into LeetCode:")
		fmt.Println("1. Open Chrome DevTools (F12)")
		fmt.Println("2. Go to Application > Cookies > https://leetcode.com")
		fmt.Println("3. Find the csrftoken and LEETCODE_SESSION cookies")
		return
	}

	// Get the absolute path of the workspace root
	workspaceRoot, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}
	workspaceRoot = filepath.Dir(workspaceRoot) // Go up one directory to reach the root

	fmt.Printf("Processing directories in: %s\n", workspaceRoot)

	// Fetch problem numbers from LeetCode API
	slugToNumber, err := fetchProblems()
	if err != nil {
		fmt.Printf("Error fetching problems: %v\n", err)
		return
	}

	fmt.Printf("Found %d problems from LeetCode\n", len(slugToNumber))

	// Read workspace root directory
	entries, err := os.ReadDir(workspaceRoot)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	// System directories to skip
	systemDirs := map[string]bool{
		"cmd":             true,
		"leetcode-db":     true,
		"problems_backup": true,
		".git":            true,
		".github":         true,
	}

	// Process each directory
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		oldName := entry.Name()

		// Skip system directories and hidden directories
		if systemDirs[oldName] || strings.HasPrefix(oldName, ".") {
			continue
		}

		// Skip if already has a number prefix
		if _, err := strconv.Atoi(strings.Split(oldName, "-")[0]); err == nil {
			continue
		}

		// Skip if it's not a LeetCode problem directory (should contain at least one .go file)
		files, err := os.ReadDir(filepath.Join(workspaceRoot, oldName))
		if err != nil {
			continue
		}
		hasGoFile := false
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
				hasGoFile = true
				break
			}
		}
		if !hasGoFile {
			continue
		}

		normalizedName := normalizeTitle(oldName)

		// Try exact match first
		number, ok := slugToNumber[normalizedName]
		if !ok {
			// Try without special characters
			for slug, num := range slugToNumber {
				if normalizeTitle(slug) == normalizedName {
					number = num
					ok = true
					break
				}
			}
		}

		if ok && number > 0 {
			// Create new directory name with problem number
			newName := fmt.Sprintf("%d-%s", number, strings.ReplaceAll(oldName, "_", "-"))
			oldPath := filepath.Join(workspaceRoot, oldName)
			newPath := filepath.Join(workspaceRoot, newName)

			// Move directory
			if err := os.Rename(oldPath, newPath); err != nil {
				fmt.Printf("Error renaming %s to %s: %v\n", oldPath, newPath, err)
			} else {
				fmt.Printf("Renamed %s to %s\n", oldName, newName)
			}
		} else {
			fmt.Printf("Could not find problem number for: %s (normalized: %s)\n", oldName, normalizedName)
		}
	}
}
