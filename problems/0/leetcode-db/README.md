# LeetCode Solutions Database

A PostgreSQL database with PostgREST API for managing and querying LeetCode solutions.

## Prerequisites

- Docker and Docker Compose
- Go 1.x (for running the parser)
- Git (for cloning the repository)

## Setup

1. Start the database and API services:
```bash
cd leetcode-db
docker-compose up -d
```

This will start:
- PostgreSQL database on port 5432
- PostgREST API on port 3000
- Swagger UI on port 8080

2. Run the parser to populate the database:
```bash
cd leetcode-db
go run cmd/main.go
```

## API Endpoints

### View API Documentation
Open http://localhost:8080 in your browser to access the Swagger UI.

### Common Queries

1. List all problems with their solutions:
```
http://localhost:3000/problem_summary
```

2. Search problems by language (note: use lowercase language names):
```
http://localhost:3000/formatted_solutions?language=eq.go
```

3. Get language statistics:
```
http://localhost:3000/language_stats
```

4. View recent solutions:
```
http://localhost:3000/recent_solutions
```

5. Search problems by text:
```
http://localhost:3000/rpc/search_problems?search_query=two%20sum
```

### Advanced Filtering

PostgREST supports various filtering operators:

- Exact match: `?column=eq.value`
- Greater than: `?column=gt.value`
- Less than: `?column=lt.value`
- Like: `?column=like.*pattern*`
- In list: `?column=in.(value1,value2)`

Example:
```
# Find problems solved after 2023
http://localhost:3000/problems?solved_at=gt.2023-01-01

# Find solutions in multiple languages (note: use lowercase)
http://localhost:3000/problem_solutions?language=in.(go,java)
```

### Ordering Results

Add `order` parameter to sort results:
```
# Sort problems by solved date (descending)
http://localhost:3000/problems?order=solved_at.desc

# Sort language stats by solution count
http://localhost:3000/language_stats?order=solution_count.desc
```

## Database Schema

### Tables

1. `problems`
   - `id`: Problem ID from LeetCode
   - `title`: Problem title
   - `description`: Problem description
   - `folder_path`: Path to the solution folder
   - `solved_at`: When the problem was solved
   - `created_at`: Record creation timestamp
   - `updated_at`: Record update timestamp

2. `solutions`
   - `id`: Auto-incrementing ID
   - `problem_id`: References problems(id)
   - `language`: Programming language
   - `code`: Solution code
   - `file_path`: Path to the solution file
   - `commit_date`: When the solution was committed
   - `created_at`: Record creation timestamp
   - `updated_at`: Record update timestamp

### Views

1. `problem_solutions`: Joins problems with their solutions
2. `problem_summary`: Problem statistics with solution counts
3. `language_stats`: Solution counts by programming language
4. `recent_solutions`: Latest solutions with problem details

## Maintenance

### Reset Database
To reset the database and start fresh:
```bash
docker-compose down -v
rm -rf pgdata/
docker-compose up -d
```

### View Database Logs
```bash
docker-compose logs db
```

### Connect to Database Directly
```bash
docker-compose exec db psql -U leetcode_user -d leetcode
```