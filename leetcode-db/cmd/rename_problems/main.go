package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Problem struct {
	StatStatusPairs []struct {
		Stat struct {
			QuestionID        int    `json:"question_id"`
			QuestionTitle     string `json:"question__title"`
			QuestionTitleSlug string `json:"question__title_slug"`
		} `json:"stat"`
	} `json:"stat_status_pairs"`
}

type GraphQLResponse struct {
	Data struct {
		Question struct {
			QuestionID string `json:"questionId"`
			Title      string `json:"title"`
			TitleSlug  string `json:"titleSlug"`
		} `json:"question"`
	} `json:"data"`
}

// Known problem mappings that might not be in the API response
var knownProblemMappings = map[string]string{
	"bulb-switcher-iii": "1375",
}

func fetchProblemBySlug(slug string) (string, error) {
	// Check known mappings first
	if id, ok := knownProblemMappings[slug]; ok {
		return id, nil
	}

	query := fmt.Sprintf(`{
		"query": "query questionData($titleSlug: String!) { question(titleSlug: $titleSlug) { questionId title titleSlug } }",
		"variables": {"titleSlug": "%s"}
	}`, slug)

	req, err := http.NewRequest("POST", "https://leetcode.com/graphql", bytes.NewBuffer([]byte(query)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result GraphQLResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if result.Data.Question.QuestionID == "" {
		return "", fmt.Errorf("problem not found")
	}

	return result.Data.Question.QuestionID, nil
}

func normalizeTitle(title string) string {
	// Remove leading "0-" prefix if it exists
	if strings.HasPrefix(title, "0-") {
		title = title[2:]
	}

	// Convert to lowercase
	title = strings.ToLower(title)

	// Remove special characters
	title = strings.ReplaceAll(title, "'", "")
	title = strings.ReplaceAll(title, ".", "")
	title = strings.ReplaceAll(title, ",", "")
	title = strings.ReplaceAll(title, "(", "")
	title = strings.ReplaceAll(title, ")", "")
	title = strings.ReplaceAll(title, "?", "")
	title = strings.ReplaceAll(title, "!", "")
	title = strings.ReplaceAll(title, ":", "")
	title = strings.ReplaceAll(title, ";", "")
	title = strings.ReplaceAll(title, "+", "plus")
	title = strings.ReplaceAll(title, "=", "equals")
	title = strings.ReplaceAll(title, "&", "and")

	// Replace spaces and underscores with hyphens
	title = strings.ReplaceAll(title, " ", "-")
	title = strings.ReplaceAll(title, "_", "-")

	// Handle consecutive hyphens
	for strings.Contains(title, "--") {
		title = strings.ReplaceAll(title, "--", "-")
	}

	// Remove leading/trailing hyphens
	title = strings.Trim(title, "-")

	return title
}

func generateVariations(title string) []string {
	variations := []string{title}

	// Add variations with different number formats
	re := regexp.MustCompile(`-([0-9]+|[ivx]+)$`)
	if matches := re.FindStringSubmatch(title); len(matches) > 1 {
		num := matches[1]
		base := title[:len(title)-len(num)-1]

		// Convert roman numerals to numbers
		romanToNum := map[string]string{
			"ii":   "2",
			"iii":  "3",
			"iv":   "4",
			"v":    "5",
			"vi":   "6",
			"vii":  "7",
			"viii": "8",
			"ix":   "9",
		}
		if arabic, ok := romanToNum[num]; ok {
			variations = append(variations, base+"-"+arabic)
		}

		// Convert numbers to roman numerals
		numToRoman := map[string]string{
			"2": "ii",
			"3": "iii",
			"4": "iv",
			"5": "v",
			"6": "vi",
			"7": "vii",
			"8": "viii",
			"9": "ix",
		}
		if roman, ok := numToRoman[num]; ok {
			variations = append(variations, base+"-"+roman)
		}
	}

	// Add variations without hyphens
	variations = append(variations, strings.ReplaceAll(title, "-", ""))
	variations = append(variations, strings.ReplaceAll(title, "-", " "))

	return variations
}

func fetchProblems() (map[string]string, map[string]string, error) {
	url := "https://leetcode.com/api/problems/all/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var problemList Problem
	if err := json.Unmarshal(body, &problemList); err != nil {
		return nil, nil, err
	}

	slugToID := make(map[string]string)
	titleToID := make(map[string]string)
	for _, problem := range problemList.StatStatusPairs {
		slug := problem.Stat.QuestionTitleSlug
		id := fmt.Sprintf("%d", problem.Stat.QuestionID)
		slugToID[slug] = id
		slugToID[normalizeTitle(slug)] = id
		titleToID[normalizeTitle(problem.Stat.QuestionTitle)] = id

		// Print out any problems with "bulb" in the title
		if strings.Contains(strings.ToLower(slug), "bulb") {
			fmt.Printf("Found bulb problem: %s (ID: %s)\n", slug, id)
		}
	}

	// Add known mappings
	for slug, id := range knownProblemMappings {
		slugToID[slug] = id
		slugToID[normalizeTitle(slug)] = id
	}

	fmt.Printf("Found %d problems from LeetCode\n", len(problemList.StatStatusPairs))
	return slugToID, titleToID, nil
}

func isLeetCodeProblemDir(path string) bool {
	// Skip if it's a system directory
	base := filepath.Base(path)
	if strings.HasPrefix(base, ".") || strings.HasPrefix(base, "_") {
		return false
	}

	// Skip if it's a known non-problem directory
	knownNonProblems := map[string]bool{
		"cmd":             true,
		"leetcode-db":     true,
		"problems_backup": true,
	}
	if knownNonProblems[base] {
		return false
	}

	// Check if it contains a Go file
	files, err := os.ReadDir(path)
	if err != nil {
		return false
	}

	hasGoFile := false
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			hasGoFile = true
			break
		}
	}

	return hasGoFile
}

func main() {
	// Get current working directory
	workspaceRoot, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	// Fetch problems from LeetCode API
	slugToID, titleToID, err := fetchProblems()
	if err != nil {
		fmt.Printf("Error fetching problems: %v\n", err)
		return
	}

	// Read all directories
	entries, err := os.ReadDir(workspaceRoot)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	// Process each directory
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		dirName := entry.Name()
		dirPath := filepath.Join(workspaceRoot, dirName)

		// Skip if not a LeetCode problem directory
		if !isLeetCodeProblemDir(dirPath) {
			continue
		}

		// Skip if already in correct format (starts with a number but not "0-")
		if matched, _ := regexp.MatchString(`^[1-9]\d*-`, dirName); matched {
			continue
		}

		// Get the problem name without the "0-" prefix
		problemName := dirName
		if strings.HasPrefix(dirName, "0-") {
			problemName = dirName[2:]
		}

		// Try different variations of the directory name
		normalizedName := normalizeTitle(problemName)
		variations := generateVariations(normalizedName)

		var id string
		var found bool
		for _, variation := range variations {
			if id, found = slugToID[variation]; found {
				break
			}
			if id, found = titleToID[variation]; found {
				break
			}
		}

		// If not found in the list, try fetching directly from GraphQL API
		if !found {
			id, err = fetchProblemBySlug(normalizedName)
			if err == nil {
				found = true
			}
		}

		if found {
			newName := fmt.Sprintf("%s-%s", id, normalizedName)
			newPath := filepath.Join(workspaceRoot, newName)

			if err := os.Rename(dirPath, newPath); err != nil {
				fmt.Printf("Error renaming %s to %s: %v\n", dirPath, newPath, err)
			} else {
				fmt.Printf("Renamed %s to %s\n", dirName, newName)
			}
		} else {
			fmt.Printf("Could not find problem number for: %s (normalized: %s)\n", dirName, normalizedName)
			fmt.Printf("  Tried variations: %v\n", variations)
			// Print a few similar slugs for debugging
			fmt.Printf("  Similar slugs in LeetCode:\n")
			count := 0
			for slug := range slugToID {
				if strings.Contains(slug, normalizedName[:min(5, len(normalizedName))]) {
					fmt.Printf("    - %s\n", slug)
					count++
					if count >= 5 {
						break
					}
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
