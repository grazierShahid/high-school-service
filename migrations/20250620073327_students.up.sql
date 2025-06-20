CREATE TABLE IF NOT EXISTS students (
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    nickname TEXT NOT NULL,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    addr TEXT,
    student_background_id INT,
    class_id INT,
    guardian_id INT,
    payment_id INT,
    result_id INT,
    admitted_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_students_username ON students(username);
--CREATE INDEX idx_students_id ON students(id);
CREATE INDEX idx_students_email ON students(email);