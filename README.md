# LeetCode Solutions Manager

This repository contains my LeetCode solutions along with a management system that:
1. Automatically syncs solutions from LeetCode
2. Organizes problems into structured directories
3. Maintains a PostgreSQL database of problems and solutions
4. Provides a REST API to query solutions (Using Postgrest)

## Project Structure

```
.
├── problems/             # Organized LeetCode solutions
│   ├── 0/                # Problems without LeetCode IDs
│   ├── 1-100/            # Problems 1-100
│   ├── 101-200/          # Problems 101-200
│   └── ...               # More problem groups
├── leetcode-db/          # Database and API management
│   ├── cmd/             
│   │   └── populate-db/  # Problem organization scripts
│   ├── internal/         # Internal packages
│   └── scripts/          # Database initialization
└── .github/workflows/    # GitHub Actions workflows
```

## Features

### 1. Automatic Syncing
- Weekly automatic sync with LeetCode (every Saturday at 8:00 UTC)
- Manual sync available through GitHub Actions
- Uses [leetcode-sync](https://github.com/joshcai/leetcode-sync)

### 2. Problem Organization
- Problems are grouped by hundreds (1-100, 101-200, etc.)
- Special '0' directory for problems without LeetCode IDs
- Consistent naming format: `{id}-{problem_name}`

### 3. Database Features
- Stores problem metadata and solutions
- Tracks multiple solutions per problem
- Maintains solution history with commit dates
- Full-text search for problem titles
- Language statistics

### 4. REST API (via PostgREST)
Available endpoints:
- `/recent_solutions`: Get recently updated solutions
- `/problem_solutions`: Get all solutions with problem details
- `/problem_summary`: Get problem statistics
- `/language_stats`: Get programming language statistics
- `/rpc/search_problems`: Search problems by title (function endpoint)

## Setup

1. **Clone the Repo**
   ```bash
   git clone https://github.com/nathannaveen/leetcode-submissions.git
   cd leetcode-submissions
   ```

2. **Database Setup**
   ```bash
   cd leetcode-db
   docker-compose up -d
   ```

3. **Populate Database**
   ```bash
   go run cmd/main.go
   ```

## GitHub Actions Workflows

### 1. Sync Workflow (`sync_leetcode.yml`)
- Syncs solutions from LeetCode
- Requires LeetCode session tokens (set in repository secrets)
- Runs weekly or manually

### 2. Process Workflow (`process_leetcode.yml`)
Triggered after successful sync:
1. Organizes problems into directories
2. Renames problems in the '0' directory
3. Updates the database

## API Usage Examples

1. **Get Recent Solutions**
   ```
   GET http://localhost:3000/recent_solutions?limit=10
   ```

2. **Search Problems**
   ```
   GET http://localhost:3000/rpc/search_problems?search_query=two%20sum
   ```
   Response:
   ```json
   [
     {
       "id": 1,
       "title": "Two Sum",
       "description": "Given an array of integers...",
       "relevance": 0.8
     }
   ]
   ```

3. **Get Language Statistics**
   ```
   GET http://localhost:3000/language_stats
   ```
   Response:
   ```json
   [
     {
       "language": "go",
       "solution_count": 150,
       "problem_ids": [1, 2, 3, ...]
     }
   ]
   ```

4. **Get Problem Solutions**
   ```
   GET http://localhost:3000/problem_solutions?problem_id=eq.1
   ```
   Response:
   ```json
   [
     {
       "id": 1,
       "title": "Two Sum",
       "language": "java",
       "code": "...",
       "commit_date": "..."
     }
   ]
   ```

5. **Get Problem Summary**
   ```
   GET http://localhost:3000/problem_summary?order=solved_at.desc
   ```
   Response:
   ```json
   [
     {
       "id": 1,
       "title": "Two Sum",
       "solution_count": 2,
       "languages": ["go", "java"]
     }
   ]
   ```

Note: All endpoints support PostgREST's filtering, ordering, and pagination:
- `?limit=10`: Limit results
- `?order=column.desc`: Order results
- `?column=eq.value`: Exact match
- `?column=like.*pattern*`: Pattern match
- `?select=column1,column2`: Select specific columns