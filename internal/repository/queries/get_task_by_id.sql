SELECT
    id,
    title,
    is_done,
    created_at
FROM task
WHERE id = $1
LIMIT 1;