package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
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

// Use git mv instead of os.Rename to preserve history
func gitMv(oldPath, newPath string) error {
	cmd := exec.Command("git", "mv", oldPath, newPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git mv failed: %v", err)
	}
	return nil
}

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

		// Rename the directory using git mv
		err := gitMv(oldPath, newPath)
		if err != nil {
			fmt.Printf("Error renaming %s to %s: %v\n", oldPath, newPath, err)
			continue
		}

		fmt.Printf("Renamed: %s -> %s\n", oldName, newName)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
