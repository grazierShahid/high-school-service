CREATE TABLE IF NOT EXISTS teachers (
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    addr TEXT,
    teacher_background_id INT,
    major_sub TEXT,
    minor_sub TEXT,
    joined_date TIMESTAMP,
    salary_id INT,
    birth_year INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_teacher_username ON teachers(username);
--CREATE INDEX idx_teacher_id ON teachers(id);
CREATE INDEX idx_teacher_id ON teachers(email);