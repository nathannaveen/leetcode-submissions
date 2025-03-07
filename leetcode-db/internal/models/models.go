package models

import "time"

// Problem represents a LeetCode problem
type Problem struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	FolderPath  string    `json:"folder_path"`
	SolvedAt    time.Time `json:"solved_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Solution represents a solution to a LeetCode problem
type Solution struct {
	ID         int       `json:"id"`
	ProblemID  int       `json:"problem_id"`
	Language   string    `json:"language"`
	Code       string    `json:"code"`
	FilePath   string    `json:"file_path"`
	CommitDate time.Time `json:"commit_date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
