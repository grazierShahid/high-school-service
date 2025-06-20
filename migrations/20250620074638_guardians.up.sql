CREATE TABLE IF NOT EXISTS guardians (
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    addr TEXT,
    profession TEXT,
    student_id INT
);

--CREATE INDEX idx_guardians_id ON guardians(id);
CREATE INDEX idx_guardians_student_id ON guardians(student_id);