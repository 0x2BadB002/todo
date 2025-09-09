CREATE TABLE IF NOT EXISTS tasks (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(140) NOT NULL,
  description VARCHAR(200),
  priority INT,
  due_at TIME,
  updated_at TIMESTAMP,
  created_at DATETIME,
  deleted_at DATETIME
);
