package database

import (
	"database/sql"
	"fmt"

	"leetcode-db/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) StoreProblem(problem *models.Problem) error {
	_, err := r.db.Exec(`
        INSERT INTO problems (id, title, description, folder_path, solved_at)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (id) DO UPDATE SET
            title = EXCLUDED.title,
            description = EXCLUDED.description,
            folder_path = EXCLUDED.folder_path,
            solved_at = EXCLUDED.solved_at,
            updated_at = CURRENT_TIMESTAMP
    `, problem.ID, problem.Title, problem.Description, problem.FolderPath, problem.SolvedAt)

	if err != nil {
		return fmt.Errorf("failed to store problem: %v", err)
	}
	return nil
}

func (r *Repository) StoreSolution(solution *models.Solution) error {
	_, err := r.db.Exec(`
        INSERT INTO solutions (problem_id, language, code, file_path, commit_date)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (problem_id, language) DO UPDATE SET
            code = EXCLUDED.code,
            file_path = EXCLUDED.file_path,
            commit_date = EXCLUDED.commit_date,
            updated_at = CURRENT_TIMESTAMP
    `, solution.ProblemID, solution.Language, solution.Code, solution.FilePath, solution.CommitDate)

	if err != nil {
		return fmt.Errorf("failed to store solution: %v", err)
	}
	return nil
}
