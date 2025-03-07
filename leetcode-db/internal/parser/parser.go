package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"leetcode-db/internal/models"
)

type ProblemParser struct {
	rootDir string
	nextID  int // For generating sequential IDs for problems without numbers
}

func NewParser(rootDir string) *ProblemParser {
	return &ProblemParser{
		rootDir: rootDir,
		nextID:  1000000, // Start high to avoid conflicts with real LeetCode IDs
	}
}

// extractProblemID attempts to extract a problem ID from various formats
func (p *ProblemParser) extractProblemID(name string, dirPath string) int {
	// Try direct number prefix (e.g., "1234-problem-name")
	if parts := strings.SplitN(name, "-", 2); len(parts) == 2 {
		if id, err := strconv.Atoi(parts[0]); err == nil {
			return id
		}
	}

	// Try roman numerals (e.g., "two_sum_ii_-_input_array_is_sorted")
	romanMap := map[string]int{
		"i":    1,
		"ii":   2,
		"iii":  3,
		"iv":   4,
		"v":    5,
		"vi":   6,
		"vii":  7,
		"viii": 8,
		"ix":   9,
		"x":    10,
	}

	// Extract roman numeral if present
	re := regexp.MustCompile(`_([ivx]+)_`)
	matches := re.FindStringSubmatch(name)
	if len(matches) > 1 {
		if num, ok := romanMap[matches[1]]; ok {
			return num
		}
	}

	// Generate a deterministic ID based on the full path
	// This ensures the same problem always gets the same ID
	hash := 0
	fullPath := filepath.Join(dirPath, name)
	for _, c := range fullPath {
		hash = (hash*31 + int(c)) % 900000 // Keep under 1000000 to avoid conflicts
	}
	return p.nextID + hash
}

// cleanTitle formats the title consistently
func cleanTitle(name string) string {
	// Remove any leading numbers and hyphens
	name = regexp.MustCompile(`^\d+\-`).ReplaceAllString(name, "")

	// Remove roman numeral suffix if present
	name = regexp.MustCompile(`_[ivx]+_`).ReplaceAllString(name, " ")

	// Replace special characters with spaces
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, ",", " ")

	// Clean up multiple spaces
	name = regexp.MustCompile(`\s+`).ReplaceAllString(name, " ")
	name = strings.TrimSpace(name)

	// Capitalize first letter of each word
	words := strings.Fields(name)
	for i, word := range words {
		if len(word) <= 3 && (word == "and" || word == "or" || word == "the" || word == "in" || word == "of" || word == "to" || word == "is" || word == "a") {
			continue
		}
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}

	return strings.Join(words, " ")
}

// getGitCommitDate gets the latest commit date for a file
func (p *ProblemParser) getGitCommitDate(filePath string) time.Time {
	// Use --follow to track file history through renames
	cmd := exec.Command("git", "log", "--follow", "-1", "--format=%aI", "--", filePath)
	cmd.Dir = p.rootDir
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error getting commit date for %s: %v\n", filePath, err)
		return time.Now() // Fallback to current time if git command fails
	}
	date := strings.TrimSpace(string(out))
	if date == "" {
		fmt.Printf("Empty commit date for %s, using current time\n", filePath)
		return time.Now()
	}
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		fmt.Printf("Error parsing commit date %q for %s: %v\n", date, filePath, err)
		return time.Now()
	}
	return t
}

// isLeetCodeProblem checks if a directory is likely a LeetCode problem
func isLeetCodeProblem(name string) bool {
	// Skip special directories
	if name == "problems" || name == "leetcode-db" || name == "scripts" || strings.HasPrefix(name, ".") {
		return false
	}

	// LeetCode problem directories should have this format:
	// "1234-some-problem-name" or "1234_some_problem_name"
	// where 1234 is the problem number

	// Split on first hyphen or underscore
	parts := regexp.MustCompile(`[-_]`).Split(name, 2)
	if len(parts) != 2 {
		return false
	}

	// First part should be a number
	problemID := parts[0]
	if _, err := strconv.Atoi(problemID); err != nil {
		return false
	}

	// Second part should be the problem name with only letters, numbers, hyphens, or underscores
	problemName := parts[1]
	validNamePattern := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	if !validNamePattern.MatchString(problemName) {
		return false
	}

	return true
}

// FindProblemDirs finds all problem directories in both root and problems/
func (p *ProblemParser) FindProblemDirs() ([]string, error) {
	var dirs []string

	// Check root directory for problem folders (e.g., "1234-problem-name")
	rootEntries, err := ioutil.ReadDir(p.rootDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range rootEntries {
		if !entry.IsDir() {
			continue
		}
		if !isLeetCodeProblem(entry.Name()) {
			continue
		}
		dirs = append(dirs, filepath.Join(p.rootDir, entry.Name()))
	}

	// Check problems/ directory and its subdirectories
	problemsDir := filepath.Join(p.rootDir, "problems")
	if _, err := os.Stat(problemsDir); err == nil {
		// First, get all hundred groups and the zero directory
		hundredGroups, err := ioutil.ReadDir(problemsDir)
		if err != nil {
			return nil, err
		}

		// Process each hundred group directory
		for _, group := range hundredGroups {
			if !group.IsDir() {
				continue
			}

			// Process problems in this hundred group
			groupPath := filepath.Join(problemsDir, group.Name())
			problems, err := ioutil.ReadDir(groupPath)
			if err != nil {
				fmt.Printf("Error reading directory %s: %v\n", groupPath, err)
				continue
			}

			for _, problem := range problems {
				if !problem.IsDir() {
					continue
				}
				if !isLeetCodeProblem(problem.Name()) {
					continue
				}
				dirs = append(dirs, filepath.Join(groupPath, problem.Name()))
			}
		}
	}

	return dirs, nil
}

// ParseProblem extracts problem information from a directory
func (p *ProblemParser) ParseProblem(dirPath string) (*models.Problem, []*models.Solution, error) {
	base := filepath.Base(dirPath)

	// Skip non-LeetCode problems
	if !isLeetCodeProblem(base) {
		return nil, nil, fmt.Errorf("skipping non-LeetCode directory: %s", base)
	}

	// Extract ID and clean up title
	id := p.extractProblemID(base, dirPath)
	title := cleanTitle(base)

	// Read README.md if it exists
	var description string
	readmePath := filepath.Join(dirPath, "README.md")
	if content, err := ioutil.ReadFile(readmePath); err == nil {
		description = string(content)
	}

	// Get the latest commit date for any file in the directory
	var latestCommitDate time.Time
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.Join(dirPath, file.Name())
		commitDate := p.getGitCommitDate(filePath)
		if commitDate.After(latestCommitDate) {
			latestCommitDate = commitDate
		}
	}

	if latestCommitDate.IsZero() {
		latestCommitDate = time.Now()
	}

	// Create problem
	problem := &models.Problem{
		ID:          id,
		Title:       title,
		Description: description,
		FolderPath:  dirPath,
		SolvedAt:    latestCommitDate,
	}

	// Find all solution files
	var solutions []*models.Solution
	for _, file := range files {
		if file.IsDir() || file.Name() == "README.md" {
			continue
		}

		// Get file extension (language)
		ext := strings.TrimPrefix(filepath.Ext(file.Name()), ".")
		if ext == "" {
			continue
		}

		// Only process common programming language files
		validExts := map[string]bool{
			"go": true, "java": true, "py": true, "cpp": true,
			"c": true, "js": true, "ts": true, "rb": true,
			"rs": true, "php": true, "cs": true, "scala": true,
			"kt": true, "swift": true,
		}
		if !validExts[strings.ToLower(ext)] {
			continue
		}

		// Read solution code
		filePath := filepath.Join(dirPath, file.Name())
		code, err := ioutil.ReadFile(filePath)
		if err != nil {
			continue
		}

		// Handle UTF-8 encoding issues
		codeStr := string(code)
		if !utf8.ValidString(codeStr) {
			// Try to clean the string by removing invalid UTF-8 sequences
			codeStr = strings.ToValidUTF8(codeStr, "")
		}

		solution := &models.Solution{
			ProblemID:  problem.ID,
			Language:   strings.ToLower(ext),
			Code:       codeStr,
			FilePath:   filePath,
			CommitDate: p.getGitCommitDate(filePath),
		}
		solutions = append(solutions, solution)
	}

	return problem, solutions, nil
}
