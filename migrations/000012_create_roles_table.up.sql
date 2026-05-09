CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    level INT NOT NULL DEFAULT 0,
    description TEXT
);


INSERT INTO roles (name, level, description)
VALUES 
('user', 1, 'Can create posts, comments, and follow other users'),
('moderator', 2, 'Can manage and moderate other users posts'),
('admin', 3, 'Has full access to manage users and posts');