CREATE ROLE web_anon NOLOGIN;
GRANT ALL ON ALL TABLES IN SCHEMA public TO web_anon;

CREATE TABLE problems (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    folder_path TEXT,
    solved_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE solutions (
    id SERIAL PRIMARY KEY,
    problem_id INTEGER REFERENCES problems(id),
    language TEXT NOT NULL,
    code TEXT NOT NULL,
    file_path TEXT,
    commit_date TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(problem_id, language)
);

CREATE INDEX idx_problems_title ON problems USING gin(to_tsvector('english', title));
CREATE INDEX idx_solutions_language ON solutions(language);

CREATE VIEW formatted_solutions AS
SELECT 
    id,
    problem_id,
    language,
    replace(
        replace(
            replace(code, '\n', E'\n'),
            '\t', E'\t'
        ),
        '\"', '"'
    ) as code,
    file_path,
    commit_date,
    updated_at
FROM solutions;

CREATE VIEW problem_solutions AS
SELECT 
    p.id,
    s.problem_id,
    p.title,
    p.solved_at,
    s.language,
    s.code,
    s.file_path,
    s.commit_date
FROM problems p
LEFT JOIN solutions s ON p.id = s.problem_id;

CREATE VIEW problem_summary AS
SELECT 
    p.id,
    p.title,
    p.solved_at,
    COUNT(DISTINCT s.language) as solution_count,
    array_to_json(array_agg(DISTINCT s.language)) as languages
FROM problems p
LEFT JOIN solutions s ON p.id = s.problem_id
GROUP BY p.id, p.title, p.solved_at;

CREATE VIEW language_stats AS
SELECT 
    language,
    COUNT(*) as solution_count,
    array_agg(problem_id) as problem_ids
FROM solutions
GROUP BY language
ORDER BY solution_count DESC;

CREATE VIEW recent_solutions AS
SELECT 
    s.id,
    s.problem_id,
    p.title as problem_title,
    s.language,
    s.commit_date,
    s.updated_at
FROM solutions s
JOIN problems p ON s.problem_id = p.id
ORDER BY s.commit_date DESC NULLS LAST;

CREATE FUNCTION search_problems(search_query TEXT) 
RETURNS TABLE (
    id INTEGER,
    title TEXT,
    description TEXT,
    relevance REAL
) 
SECURITY DEFINER
AS $$
BEGIN
    RETURN QUERY
    SELECT 
        p.id,
        p.title,
        p.description,
        ts_rank(to_tsvector('english', p.title || ' ' || COALESCE(p.description, '')), 
                plainto_tsquery('english', search_query)) as relevance
    FROM problems p
    WHERE to_tsvector('english', p.title || ' ' || COALESCE(p.description, '')) @@ 
          plainto_tsquery('english', search_query)
    ORDER BY relevance DESC;
END;
$$ LANGUAGE plpgsql;

GRANT SELECT ON problem_solutions TO web_anon;
GRANT SELECT ON problem_summary TO web_anon;
GRANT SELECT ON language_stats TO web_anon;
GRANT SELECT ON recent_solutions TO web_anon;
GRANT SELECT ON formatted_solutions TO web_anon;
GRANT EXECUTE ON FUNCTION search_problems(TEXT) TO web_anon; 