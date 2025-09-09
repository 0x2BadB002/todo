CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(140) NOT NULL,
  description VARCHAR(200),
  priority INT,
  due_at DATETIME,
  updated_at TIMESTAMP,
  created_at DATETIME,
  deleted_at DATETIME
);
