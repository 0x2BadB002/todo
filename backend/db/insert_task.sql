INSERT INTO tasks (name, description, priority, due_at, updated_at, created_at, deleted_at)
VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_DATE, NULL);
